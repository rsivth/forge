// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import "strings"

// MissingFrom infers loci and alleles of person p that are missing from sample
// s and returns them as sample. It only considers loci from sample s.
func (p Sample) MissingFrom(s Sample) Sample {

	m := Sample{
		ID: strings.Join([]string{"Missing", p.ID, "from", s.ID}, "_"),
		// TODO: the kit solution below should not be in forge!
		Kit: s.Kit, // we need this info for the missing-rows in match report
	}

	// walk through the sample's loci
	for _, sl := range s.Loci {

		if p.Sex() == "female" && sl.Linkage() == YLINKED {
			continue
		}

		// if the person does not have the locus, move on
		if !p.HasLocus(sl.ID) {
			continue
		}

		// walk through the person's alleles and check whether the stain has
		// them.
		missLoc := NewLocus(sl.ID)
		for _, pa := range p.Locus(sl.ID).Alleles {
			if !sl.HasAllele(pa.ID) {
				missLoc.AddAllele(pa)
			}
		}

		if len(missLoc.Alleles) > 0 {
			m.AddLocus(missLoc)
		}
	}

	return m
}

// MissingFromInt returns the number of alleles of the person p that are missing
// from sample s.
func (p Sample) MissingFromInt(s Sample) int {

	m := p.MissingFrom(s)

	var i int
	for _, l := range m.Loci {
		i += len(l.Alleles)
	}

	return i
}

// ShareAnyLoci returns whether p has a any loci which are also present in s.
func (p Sample) ShareAnyLoci(s Sample) bool {

	for _, l := range s.Loci {
		if p.HasLocus(l.ID) && len(p.Locus(l.ID).Alleles) > 0 {
			return true
		}

	}

	return false
}
