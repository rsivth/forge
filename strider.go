// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const striderURL = "https://strider.online/frequencies_xml/download"

// STRiderFreqs holds the the frequency information from
// https://strider.online/frequencies
type STRiderFreqs struct {
	// timestamp in the format "2020-10-19T08:50:05+02:00"
	Created string `xml:"created_timestamp,attr"`
	// timestamp in the format "2020-10-19T08:50:05+02:00"
	Validity string `xml:"validity,attr"`
	// slice of all markers in the XML file
	Markers []Marker `xml:"marker"`
}

// Marker holds the allele frequency information for a single marker
// for all populations and alleles.
type Marker struct {
	// name of the marker, e.g. VWA
	Name string `xml:"name"`
	// string of comma seperated allele names (e.g. "6, 7,
	// 9.3...") for this marker but for all populations
	Alleles string `xml:"alleles"`
	// slice of all populations (i.e. Origins) for this marker
	Origins []Origin `xml:"origin"`
}

// Origin holds the frequency information for this specific population.
type Origin struct {
	// name of the population
	Name string `xml:"name,attr"`
	// sample size of the dataset
	Num int `xml:"n,attr"`
	// slice of frequencies for each allele
	Frequencies []Frequency `xml:"frequency"`
}

// Frequency holds the frequency information for an allele.
type Frequency struct {
	// name of the allele
	Allele string `xml:"allele,attr"`
	// frequency of the allele
	Freq string `xml:",innerxml"`
}

// ReadSTRiderXML reads the allele frequency XML file (filename xml).
// Source of frequencies: https://strider.online/frequencies
func ReadSTRiderXML(f string) (STRiderFreqs, error) {

	xmlFile, err := os.Open(f)
	if err != nil {
		return STRiderFreqs{}, err
	}
	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			//TODO: handle error
		}
	}(xmlFile)

	// read opened xmlFile as a byte array.
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return STRiderFreqs{}, err
	}
	var sf STRiderFreqs
	err2 := xml.Unmarshal(byteValue, &sf)
	if err2 != nil {
		return STRiderFreqs{}, err2
	}

	return sf, nil
}

// BuildPop builds a Freqs object for population p and minimal frequency fmin
// from a STRider frequency dataset.
func (sf STRiderFreqs) BuildPop(pop string, fmin float64) (Freqs, error) {

	var fLoci []Flocus
	for _, m := range sf.Markers {

		var fAlleles []Fallele
		for _, o := range m.Origins {

			if o.Name != pop {
				continue
			}

			for _, f := range o.Frequencies {

				freq, err := strconv.ParseFloat(f.Freq, 64)
				if err != nil {
					return Freqs{}, err
				}

				fAlleles = append(fAlleles, Fallele{
					ID:   A2Float(f.Allele),
					Freq: freq,
				})
			}

			break
		}

		if len(fAlleles) == 0 {
			continue
		}

		fLoci = append(fLoci, Flocus{
			ID:       m.Name,
			Falleles: fAlleles,
		})
	}

	if len(fLoci) == 0 {
		return Freqs{}, fmt.Errorf("no frequencies found for population %v", pop)
	}

	return Freqs{
		Source: "STRider_" + sf.Validity,
		Pop:    pop,
		Fmin:   fmin,
		Floci:  fLoci,
	}, nil
}

// DownloadStriderXML downloads the frequencies from
// https://strider.online/frequencies in XML format and saves it to a file f.
func DownloadStriderXML(f string) error {

	// Create the file
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			// TODO: handle error
		}
	}(out)

	// get the data
	resp, err := http.Get(striderURL)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
