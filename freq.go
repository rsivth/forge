// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

// Freqs holds the frequency information for a specific population
// (e.g. Europe).
type Freqs struct {
	// Source of the frequency information, e.g. the STRider XML file name.
	Source string
	// Name of the population from which the frequencies have been estimated.
	Pop string
	// Minimum allele frequency to use.
	Fmin float64
	// Slice of STR markers (e.g. VWA) with the respective frequency data.
	Floci []Flocus
}

// Flocus holds the frequency information for a locus.
type Flocus struct {
	ID       string    // e.g. VWA
	Falleles []Fallele // Slice of alleles with frequency information
}

// Fallele holds the frequency information for an allele.
type Fallele struct {
	ID   float64 // e.g. 9.3
	Freq float64 // frequency
}

// Flocus returns a Flocus object with locus name id from f.
func (f Freqs) Flocus(id string) Flocus {
	for _, l := range f.Floci {
		if l.ID == id {
			return l
		}
	}

	return Flocus{}
}

// HasFlocus tests whether the frequency data f contains a locus with
// name id.
func (f Freqs) HasFlocus(id string) bool {
	return len(f.Flocus(id).Falleles) > 0
}

// Fallele returns a Fallele object with allele name id from l.
func (l Flocus) Fallele(id float64) Fallele {
	for _, a := range l.Falleles {
		if a.ID == id {
			return a
		}
	}

	return Fallele{}
}

// HasFallele tests whether the locus frequency data l contains an
// allele with name id.
func (l Flocus) HasFallele(id float64) bool {
	return l.Fallele(id).ID != 0
}
