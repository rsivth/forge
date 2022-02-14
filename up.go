// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"strconv"
)

// InferUnknownPersons returns a slice of unknown person profiles inferred from
// sample s using the four parameters.
func (s Sample) InferUnknownPersons(heteroImbalance, majorCompRatio,
	minHomozygous, weakSignal float64) []Sample {

	var UPs []Sample
	source := s.ID
	upNr := 1
	for {

		// 1. Find a major component
		mc := s.MajorComponent(heteroImbalance, majorCompRatio, minHomozygous, weakSignal)
		source += "::MC"

		// 2. collect the loci of the major component
		var autosomalLoci []Locus
		for _, l := range mc.Loci {

			// UPs are only relevant for autosomal loci
			if l.Linkage() == AUTOSOMAL {
				autosomalLoci = append(autosomalLoci, l)
			}
		}

		// 3. Do we have autosomal loci?
		if len(autosomalLoci) == 0 {
			break
		}

		// 4. build new UP object and append to the up list
		up := newUP(s.ID+"::UP"+strconv.Itoa(upNr), source, autosomalLoci)
		UPs = append(UPs, up)
		upNr++

		// 5. remove the obtained up from the stain and start over
		s = s.RemovePersonFromSample(up)
		source += "::REM"
	}

	return UPs
}

// newUP builds a new sample for an unknown person.
func newUP(id, source string, loci []Locus) Sample {

	up := NewSample(id, source)

	for _, l := range loci {
		nl := NewLocus(l.ID)
		for _, a := range l.Alleles {
			nl.AddAllele(NewAllele(a.ID))
		}
		up.AddLocus(nl)
	}

	up.UnknownKit()
	return up
}

// RemovePersonFromSample removes the alleles of profile p from sample s and
// returns a sample containing the remaining alleles.
func (s Sample) RemovePersonFromSample(p Sample) Sample {

	ns := NewSample(s.ID, s.Source+"::REM:"+p.ID)

	for _, l := range s.Loci {
		nl := l.removePersonFromLocus(p.Locus(l.ID))
		if len(nl.Alleles) > 0 {
			ns.AddLocus(nl)
		}
	}

	return ns
}

// removePersonFromLocus removes alleles from locus l that are present in the
// person's profile p.
func (l Locus) removePersonFromLocus(p Locus) Locus {

	if len(p.Alleles) == 0 {
		return l
	}

	if len(p.Alleles) > 2 {
		return Locus{}
	}

	nl := NewLocus(l.ID)
	for _, a := range l.Alleles {
		// we keep all alleles that are not present in the persons locus
		if !p.HasAllele(a.ID) {
			nl.AddAllele(a)
		}
	}

	return nl
}

// Sex attempts to infer the gender of a person with profile s.
func (s Sample) Sex() string {

	if s.MinContributor() > 1 {
		return "na"
	}

	for _, l := range s.Loci {
		if l.ID == "AMEL" {
			switch len(l.Alleles) {
			case 0:
				return "na"
			case 1:
				if l.Alleles[0].ID == -2 {
					return "female"
				}
				return "na" // single allele is Y, something is weird
			default: // i.e. 2, cannot be more because of s.MinContributor() above
				return "male"
			}
		}
	}

	return "na"
}

// mismatch
type mismatch int

const (
	match   mismatch = iota // indicates a full match, e.g. L1 15,16; L2 15,16
	nomatch                 // indicates a hard mismatch, e.g. L1 15,16; L2 21,22 and L1 15,16; L2 15,22
	fuzzy                   // indicates a soft match, e.g. L1 15,16; L2 15
)

// SamePerson evaluates whether the person profiles p1 and p2 differ in not more
// alleles than accepted.
func SamePerson(p1, p2 Sample) bool {

	if p1.MaxAlleles() > 2 || p2.MaxAlleles() > 2 {
		return false
	}

	// We accept a maximum of a quarter of the number of alleles from the
	// shortest profile as tolerance and still call it the same profile.
	// i.e. is the profile is 16 STRs long, we accept 4 mismatches; if it is
	// 8 STRs long we accept only 2 mismatches.
	tolerance := len(p1.Loci) / 4
	if len(p2.Loci) < len(p1.Loci) {
		tolerance = len(p2.Loci) / 4
	}

	count := make(map[mismatch]int)
	for _, l1 := range p1.Loci {

		if !p2.HasLocus(l1.ID) {
			continue
		}

		// count the number of matches, nomatches, and fuzzymatches between the
		// loci of p1 and p2.
		count[matchLoci(l1, p2.Locus(l1.ID))]++
	}

	// add the fuzzy matches to the nomatch batch but give them only half the
	// weight.
	count[nomatch] += count[fuzzy] / 2
	if count[nomatch] > tolerance {
		return false
	}

	return true
}

// matchLoci
// The caller (SamePerson) guarantees that l1 and l2 do not have more than two
// alleles at loci l1 and l2 and that neither locus is empty.
func matchLoci(l1, l2 Locus) mismatch {

	switch {

	case len(l1.Alleles) == 1 && len(l2.Alleles) == 2:
		// one allele is already not matching; result can only be fuzzy or
		// nomatch
		if l2.HasAllele(l1.Alleles[0].ID) {
			// soft match, e.g. L1 15,16; L2 15
			return fuzzy
		}
		return nomatch

	case len(l1.Alleles) == 2 && len(l2.Alleles) == 1:
		// one allele is already not matching; result can only be fuzzy or
		// nomatch
		if l1.HasAllele(l2.Alleles[0].ID) {
			// soft match, e.g. L1 15,16; L2 15
			return fuzzy
		}
		return nomatch

	case len(l1.Alleles) == 2 && len(l2.Alleles) == 2:
		if l2.HasAllele(l1.Alleles[0].ID) {
			if l2.HasAllele(l1.Alleles[1].ID) {
				// both alleles match, e.g. L1 15,16; L2 15,16
				return match
			}
		}
		// one or both alleles do not match
		// e.g. L1 15,16; L2 21,22 and L1 15,16; L2 15,22
		return nomatch

	default:
		// only one scenario left: both loci have exactly one allele.
		if l2.HasAllele(l1.Alleles[0].ID) {
			return match
		}
		return nomatch
	}
}

// UniteUPs unites the samples in ups into a single sample with the name id.
func UniteUPs(ups []Sample, id string) Sample {

	if len(ups) == 0 {
		return Sample{}
	}

	r := NewSample(id, "")

	for _, lc := range Composite(ups, AUTOSOMAL).Loci {

		switch len(lc.Alleles) {
		case 1:
			// If the composite locus has only 1 allele, we can add it since no
			// conflict is possible.
			r.AddLocus(lc)
		case 2:
			// Conflict can only arise from UPs with 1 allele at this locus,
			// if there are more than one UP with 1 allele and if these have
			// not the same allele ID.
			alleleIDs := make(map[float64]int)
			for _, up := range ups {
				if len(up.Locus(lc.ID).Alleles) == 1 {
					alleleIDs[up.Locus(lc.ID).Alleles[0].ID]++
				}
			}
			// If all (or none) loci with one allele at this locus have the
			// same allele id, there is no conflict, and we can add the locus.
			if len(alleleIDs) < 2 {
				r.AddLocus(lc)
			}
		default:
			// len(l.Alleles) must be > 2; ignore this locus
			continue
		}
	}

	return newUP(id, concatID(ups), r.Loci)
}

// HasNewAlleles returns whether UP1 has any alleles that UP2 does not have.
func (up1 Sample) HasNewAlleles(up2 Sample) bool {

	for _, l1 := range up1.Loci {
		if !up2.HasLocus(l1.ID) {
			return true
		}

		for _, a1 := range l1.Alleles {
			if !up2.Locus(l1.ID).HasAllele(a1.ID) {
				return true
			}
		}
	}

	return false
}
