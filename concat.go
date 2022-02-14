// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"sort"
	"strings"
)

type Concat int

const (
	COMPOSITE Concat = iota
	CONSENSUS
)

// String returns the concat mode as string.
func (c Concat) String() string {
	switch c {
	case CONSENSUS:
		return "consensus"
	default:
		return "composite"
	}
}

// Composite returns the composite profile of samples with the name id.
func Composite(samples []Sample, link LocusLinkage) Sample {
	return concatSamples(samples, link, COMPOSITE)
}

// Consensus returns the consensus profile of samples with the name id.
func Consensus(samples []Sample, link LocusLinkage) Sample {
	return concatSamples(samples, link, CONSENSUS)
}

// concatSamples concatenates samples according to mode and link into a sample.
// Empty concat loci will not be added to the concat sample. The ID of the
// return sample is the mode of the concat; the source is a string. If link is
// ALLLINKAGE all loci will be considered.
func concatSamples(samples []Sample, link LocusLinkage, mode Concat) Sample {

	if len(samples) == 0 {
		return Sample{}
	}

	source := []string{mode.String(), link.String()}
	// get a set of unique loci names with the correct linkage
	lIDs := make(map[string]int)
	for _, s := range samples {
		source = append(source, s.ID)
		for _, l := range s.Loci {
			if link == l.Linkage() || link == ALLLINKAGE {
				lIDs[l.ID]++
			}
		}
	}

	// get the IDs sorted to guarantee locus order in concat sample
	var sortedLocusIDs []string
	for l := range lIDs {
		sortedLocusIDs = append(sortedLocusIDs, l)
	}
	sort.Slice(sortedLocusIDs, func(i, j int) bool {
		return sortedLocusIDs[i] < sortedLocusIDs[j]
	})

	cs := NewSample(concatID(samples), strings.Join(source, "::"))
	for _, locusID := range sortedLocusIDs {
		var loci []Locus
		for _, s := range samples {
			loci = append(loci, s.Locus(locusID))
		}

		cs.AddLocus(concatLoci(loci, mode))
	}

	return cs
}

// concatLocus concatenates loci. If m = "consensus" it builds a
// consensus locus, if it is "composite" it builds a composite locus.
func concatLoci(loci []Locus, mode Concat) Locus {

	var lID string
	aIDs := make(map[float64]int)

	for _, l := range loci {
		if l.ID == "" {
			continue
		}
		if lID == "" {
			lID = l.ID
		}

		for _, a := range l.Alleles {
			aIDs[a.ID]++
		}
	}

	cl := NewLocus(lID)
	for aID, count := range aIDs {
		if (mode == CONSENSUS && count == len(loci)) || mode == COMPOSITE {
			cl.AddAllele(NewAllele(aID))
		}
	}

	return cl
}

// concatID returns an id string for the concat of samples with link of mode.
func concatID(samples []Sample) string {

	IDs := make(map[string]int)
	for _, s := range samples {
		IDs[s.ID]++
	}

	var r []string
	for id := range IDs {
		r = append(r, id)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return strings.Join(r, "::")
}
