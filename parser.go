// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadGM parses a Genemapper CSV file of name f into a slice of its individual
// sample objects based on the information of the column named sID. The CSV
// file must contains the following columns: the samples ID column sID,
// "Marker", and "Allele 1" as well as every column provided in info.
func ReadGM(f, sID, kitDir string, info []string) ([]Sample, error) {

	csvHeader, csvBody, err := readFile(f)
	if err != nil {
		return nil, fmt.Errorf(`reading %v fails: %v`, f, err)
	}

	idx, err := indexHeader(sID, csvHeader, info)
	if err != nil {
		return nil, fmt.Errorf(`indexing %v fails: %v`, f, err)
	}

	samples, err := parseSamples(sID, f, kitDir, csvBody, idx)
	if err != nil {
		return nil, fmt.Errorf(`processing %v fails: %v`, f, err)
	}

	return samples, nil
}

// readFile reads the csv file f with field separator sep. It returns the first
// line (header) separated from the data part of the csv file (body).
func readFile(f string) (csvHeader []string, csvBody [][]string, err error) {
	csvF, err := os.Open(f)
	if err != nil {
		return nil, nil, err
	}
	defer func(csvF *os.File) {
		err = csvF.Close()
		if err != nil {
			// TODO: handle error
		}
	}(csvF)

	fReader := csv.NewReader(csvF)
	fReader.Comma = '\t' // convert string to rune set
	// field set to the number of fields in the first record (i. e. line)
	fReader.FieldsPerRecord = 0
	// quote may NOT appear in an unquoted field
	fReader.LazyQuotes = false

	csvData, err := fReader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	return csvData[0], csvData[1:], nil
}

// index holds the column number (value) of columns of interest (key).
type index map[string]int

// newIndex returns an object index which stores the column number of the sample
// ID, the marker and the first allele. Additionally, it stores the offsets of
// Height, Size, and Area information as well as the number of alleles.
func newIndex(sID string) index {
	return index{
		sID:           -1,
		"Marker":      -1,
		"Allele 1":    -1,
		"NoOfAlleles": -1,

		// Difference between an allele column and its respective area, height,
		// and size columns.
		"AreaOffset":   -1,
		"HeightOffset": -1,
		"SizeOffset":   -1,
	}
}

// indexHeader returns the column numbers of all columns of interest (e.g. the
// sample's ID sID and further information in info such as run number) from
// the file's first row header.
func indexHeader(sID string, header, info []string) (index, error) {
	idx := newIndex(sID)
	for _, i := range info { // add all info columns to index
		idx[i] = -1
	}

	var noAlleles int
	for i, h := range header {
		if _, ok := idx[h]; ok {
			idx[h] = i
		}

		switch h {
		case "Area 1":
			idx["AreaOffset"] = i - idx["Allele 1"]
		case "Height 1":
			idx["HeightOffset"] = i - idx["Allele 1"]
		case "Size 1":
			idx["SizeOffset"] = i - idx["Allele 1"]
		}

		if strings.HasPrefix(h, "Allele ") {
			noAlleles++
		}
	}

	idx["NoOfAlleles"] = noAlleles

	// sanity checks
	if idx[sID] == -1 || idx["Marker"] == -1 || idx["Allele 1"] == -1 {
		return nil, fmt.Errorf(`cannot find %v | Marker | Allele 1`, sID)
	}

	for _, i := range info { // add all additional column names of interest
		if idx[i] == -1 {
			return nil, fmt.Errorf(`cannot find %v`, i)
		}
	}

	return idx, nil
}

// parseSamples parses the body of the csv data (csvBody) given the sample ID
// column name (sID) and index idx. It returns a slice of Sample objects.
func parseSamples(sID, f, kitDir string, csvBody [][]string, idx index) ([]Sample, error) {

	sampleMap := make(map[string]Sample)
	for _, line := range csvBody {
		id := line[idx[sID]]

		// If the current line l contains a new sample.
		if _, ok := sampleMap[id]; !ok {
			s := NewSample(id, f)
			for k, v := range idx {
				// adding only additional data requested by caller to the sample
				// (e.g. 'Dye', which is not in the newIndex()). The loci fields
				// (which are present in newIndex()) will be filled later.
				if _, ok = newIndex(sID)[k]; !ok {
					s.Info[k] = line[v]
				}
			}
			sampleMap[id] = s
		}

		locus, err := parseLocus(line, idx)
		if err != nil {
			return nil, fmt.Errorf(`processing sample %v fails at locus %v: %v`,
				id, line[idx["Marker"]], err)
		}

		s := sampleMap[id]
		s.AddLocus(locus)
		sampleMap[id] = s
	}

	// convert Sample map to slice
	var samples []Sample
	for _, s := range sampleMap {
		if err := s.InferKit(kitDir); err != nil {
			return nil, fmt.Errorf(`inferring kit for sample %v fails: %v`, s.ID, err)
		}
		samples = append(samples, s)
	}

	return samples, nil
}

// processLocus reads csv line l given index idx and returns its data as Locus
// object. All marker names will be converted to upper to avoid "vWA" vs "VWA"
// confusions.
func parseLocus(l []string, idx index) (Locus, error) {

	if l[idx["Marker"]] == "" {
		return Locus{}, fmt.Errorf("marker name missing")
	}

	loc := NewLocus(strings.ToUpper(l[idx["Marker"]]))

	for n := 1; n <= idx["NoOfAlleles"]; n++ {
		// No allele information at the position of the n_th allele at this
		// locus
		if l[allelePos(n, idx)] == "" {
			continue
		}

		allele, err := parseAllele(n, l, idx)
		if err != nil {
			return Locus{}, err
		}

		loc.AddAllele(allele)
	}

	return loc, nil
}

// offsets is declared only to improve readability in the following functions.
var offsets = []string{"AreaOffset", "HeightOffset", "SizeOffset"}

// processAllele returns the Allele object of the n_th allele in csv line l
// given index idx. Any information field that is not present in the file (i.e.
// idx[field] = -1) will return 0.
func parseAllele(n int, l []string, idx index) (Allele, error) {

	// aPos is the column number of the n_th allele in line l.
	aPos := allelePos(n, idx)

	// This is only for area, height, and size. Formatting the allele ID is
	// done further below.
	var area, height, size float64
	for _, offset := range offsets {
		if idx[offset] != -1 {

			info, err := strconv.ParseFloat(l[aPos+idx[offset]], 64)
			if err != nil {
				return Allele{}, err
			}

			switch offset {
			case "AreaOffset":
				area = info
			case "HeightOffset":
				height = info
			case "SizeOffset":
				size = info
			}
		}
	}

	return Allele{
		ID:     A2Float(l[aPos]),
		Area:   area,
		Height: height,
		Size:   size,
	}, nil
}

// isCollated returns true if the index idx of a Genemapper CSV header
// is collated (e.g. A1, S1, H1, A2, S2, H2) and false if uncollated
// (e.g. A1, A2, A3, H1, H2, H3).
func isCollated(idx index) bool {

	for _, offset := range offsets {
		if idx[offset] == 1 {
			return true
		}
	}

	return false
}

// allelePos returns the (zero-based) column no of the n-th allele given index
// idx.
func allelePos(n int, idx index) int {

	// n is the current allele number. to get the correct column number we have
	// to revert it to a zero based count
	n0 := n - 1

	// NOT collated: A1, A2, A3, H1, H2, H3
	if !isCollated(idx) {
		return idx["Allele 1"] + n0
	}

	// IS collated A1, S1, H1, A2, S2, H2
	var fields int
	for _, offset := range offsets {
		if idx[offset] != -1 {
			fields++
		}
	}
	return idx["Allele 1"] + (n0 * (fields + 1))
}

// ReadGMRefs reads a file f of reference profiles as exported from Genemapper
// and returns the profiles as samples.
func ReadGMRefs(f string) ([]Sample, error) {

	// We don't need the header for it is known, Khaleesi.
	_, csvBody, err := readFile(f)
	if err != nil {
		return nil, fmt.Errorf(`reading %v fails: %v`, f, err)
	}

	sMap := make(map[string]Sample)
	// walk through individual lines
	for _, l := range csvBody {

		if len(l) != 4 {
			return nil, fmt.Errorf(`reading %v fails: expect 4 columns, got %v`,
				f, len(l))
		}

		id := l[0]
		alleles := strings.Split(l[3], ", ")

		s := Sample{
			ID:     id,
			Source: f,
		}
		// sample ID already in map, retrieve the sample
		if _, ok := sMap[id]; ok {
			s = sMap[id]
		}

		loc := NewLocus(l[2])
		for _, a := range alleles {
			loc.AddAllele(Allele{ID: A2Float(a)})
		}

		// add locus info and data to map
		s.AddLocus(loc)

		sMap[id] = s
	}

	// convert Sample map to list
	var samples []Sample
	for _, s := range sMap {
		samples = append(samples, s)
	}

	return samples, nil
}
