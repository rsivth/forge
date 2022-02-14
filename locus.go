// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"sort"
	"strings"
)

// Locus defines the Locus struct.
type Locus struct {
	ID      string   // name of the locus (e.g. SE33)
	Alleles []Allele // slice of all alleles at this locus
}

// NewLocus generates a Locus object with ID id. All IDs will be converted to
// upper case to avoid 'vWA'/'VWA' issues.
func NewLocus(id string) Locus {
	if id == "" {
		return Locus{}
	}
	return Locus{
		ID: strings.ToUpper(id),
	}
}

// AddAllele adds an Allele a to Locus l and sorts the alleles by ID. It will
// only add the allele if it has a non-zero ID.
func (l *Locus) AddAllele(a Allele) {
	// Do not test for a.ID > 0 because X and Y are decoded as -2 und -1.
	// Alleles must also be unique, except if they are something weird (-999).
	if a.ID != 0 && (!l.HasAllele(a.ID) || a.ID == -999) {
		l.Alleles = append(l.Alleles, a)
		l.SortByID()
	}
}

// RemoveAllele removes allele a from locus l. This does not chnage the order
// of alleles, hence no sorting is required
func (l *Locus) RemoveAllele(a Allele) {

	if !l.HasAllele(a.ID) {
		return
	}

	var newAlleles []Allele
	for _, la := range l.Alleles {
		if la.ID != a.ID {
			newAlleles = append(newAlleles, la)
		}
	}

	l.Alleles = newAlleles
}

// Allele returns the allele of name id. If no such allele is found it returns
// an empty struct.
func (l Locus) Allele(id float64) Allele {
	for _, a := range l.Alleles {
		if id == a.ID {
			return a
		}
	}
	return Allele{}
}

// HasAllele returns true if Locus l contains an Allele id. Otherwise it
// returns false.
func (l Locus) HasAllele(id float64) bool {
	return l.Allele(id).ID != 0
}

// SortByID sorts locus l by allele id.
func (l Locus) SortByID() {
	sort.Slice(l.Alleles, func(i, j int) bool {
		return l.Alleles[i].ID < l.Alleles[j].ID
	})
}

// SortByHeight sorts locus l by allele height, highest allele to lowest.
// Sorting by height requires reversing of the sort, hence the implementation of
// the sort interface below.
func (l Locus) SortByHeight() Locus {
	ordL := NewLocus(l.ID)

	var ordA []Allele
	for i := range l.Alleles {
		ordA = append(ordA, l.Alleles[i])
	}
	ordL.Alleles = ordA

	sort.Sort(sort.Reverse(ordL))
	return ordL
}

// Len is needed satisfying the sort.Sort()interface. We can now sort the
// alleles by height (rfu) for locus l.
func (l Locus) Len() int {
	return len(l.Alleles)
}

// Less is needed satisfying the sort.Sort()
func (l Locus) Less(i, j int) bool {
	return l.Alleles[i].Height < l.Alleles[j].Height
}

// Swap is needed satisfying the sort.Sort()
func (l Locus) Swap(i, j int) {
	l.Alleles[i], l.Alleles[j] = l.Alleles[j], l.Alleles[i]
}
