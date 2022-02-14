// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"math/rand"
	"time"
)

// DrawSamples generates n simulations of samples with p people each based on
// the allele frequencies freqs. Uing time.Now().UnixNano() for a
// constantly-changing number, but a constant for tests.
// https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator
func (freqs Freqs) DrawSamples(n, p int) ([]Sample, error) {

	rand.Seed(time.Now().UnixNano())

	var r []Sample
	for i := 0; i < n; i++ {
		r = append(r, freqs.drawSample(p))
	}

	return r, nil
}

// drawPersons generates a random sample of p persons based on the allele
// frequencies freqs.
func (freqs Freqs) drawSample(p int) Sample {

	var s []Sample
	for i := 0; i < p; i++ {
		s = append(s, freqs.drawPerson())
	}

	return Composite(s, ALLLINKAGE)
}

// drawPerson generates a sample of one person based on the allele
// distribution freqs.
func (freqs Freqs) drawPerson() Sample {

	var loci []Locus
	for _, l := range freqs.Floci {
		loci = append(loci, Locus{
			ID: l.ID,
			Alleles: []Allele{
				l.weightedAlleleDraw(),
				l.weightedAlleleDraw(),
			},
		})
	}

	return Sample{
		ID:   "sim_" + freqs.Pop,
		Loci: loci,
	}
}

// weightedAlleleDraw draws an allele based on the allele frequencies freqs.
func (l Flocus) weightedAlleleDraw() Allele {
	if len(l.Falleles) == 0 {
		return Allele{}
	}

	var cdf float64
	for _, fa := range l.Falleles {
		cdf += fa.Freq
	}

	r := rand.Float64() * cdf

	for _, fa := range l.Falleles {
		r -= fa.Freq
		if r < 0 {
			return Allele{ID: fa.ID}
		}
	}

	return Allele{ID: l.Falleles[len(l.Falleles)-1].ID}
}
