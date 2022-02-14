// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"math"
)

// PI estimates the combined probability of inclusion for locus l, given the
// allele frequencies f of population f.Pop.
func (l Locus) PI(f Freqs, theta float64) float64 {

	// no freq info for this locus, CPI() tests for this but if PI() is called
	// individually we must take this case into account.
	if !f.HasFlocus(l.ID) {
		return 0
	}

	floc := f.Flocus(l.ID)

	var fSum float64
	for _, a := range l.Alleles {

		if !floc.HasFallele(a.ID) { // no freq info for this allele
			fSum = fSum + f.Fmin
			continue
		}

		fSum = fSum + floc.Fallele(a.ID).Freq
	}

	return math.Pow(fSum, 2) + theta*fSum*(1-fSum)
}

// PE estimates the combined probability of exclusion for locus l, given the
// allele frequencies f of population f.Pop.
func (l Locus) PE(f Freqs, theta float64) float64 {

	// if l.PI(f) == 0 then there is no frequency info for this allele in f,
	// hence we cannot calculate PE.
	if l.PI(f, 0) == 0 {
		return 0
	}

	return 1 - l.PI(f, theta)
}

// CPI estimates the combined probability of inclusion for stain s, given the
// allele frequencies f of population f.Pop.
func (s Sample) CPI(f Freqs, theta float64) float64 {

	var CPI float64
	for _, l := range s.Loci {

		// If the marker is unknown, we exclude it entirely. If the locus has no
		// alleles we exclude it too.
		if !f.HasFlocus(l.ID) || len(l.Alleles) == 0 {
			continue
		}

		if CPI == 0 { // first loop
			CPI = l.PI(f, theta)
			continue
		}

		CPI = CPI * l.PI(f, theta)
	}

	return CPI
}

// CPE estimates the combined probability of exclusion for stain s, given the
// allele frequencies f for population f.Pop.
func (s Sample) CPE(f Freqs, theta float64) float64 {
	return 1 - s.CPI(f, theta)
}

// RMNE estimates the 'random man not excluded' probability of stain s, given
// the allele frequencies f for population f.Pop.
func (s Sample) RMNE(f Freqs, theta float64) float64 {
	if s.CPI(f, theta) == 0 {
		return 0
	}

	return 1 / s.CPI(f, theta)
}
