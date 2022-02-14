// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import "fmt"

// TrueAlleleHeights returns a sample s with allele heights of all loci
// corrected for contributions from alleles ins plus/minus stutter positions.
func (s Sample) TrueAlleleHeights(ms, ps float64) Sample {

	r := NewSample(s.ID, "correctedAlleleHeight::"+s.Source)
	for _, l := range s.Loci {
		r.AddLocus(l.trueAlleleHeights(ms, ps))
	}

	r.AssignKit(s.Kit)

	return r
}

// trueAlleleHeights returns a Locus with allele heights corrected for
// contributions from alleles ins plus/minus stutter positions.
func (l Locus) trueAlleleHeights(ms, ps float64) Locus {

	if len(l.Alleles) < 2 {
		return l
	}

	r := NewLocus(l.ID)
	for i, a := range l.Alleles {

		// Are there more than i alleles at this locus? Only then we can
		// investigate if there is an allele in the plus stutter position.
		if len(l.Alleles)-1 > i && a.ID+1 == l.Alleles[i+1].ID {

			// true height of the allele in minus stutter position; is zero if
			// there is no allele there.
			var Aminus1_true float64

			// There is an allele in minus stutter position; we need the true
			// height of that allele (here, 'i' must not be the first allele
			// in the loop).
			if i > 0 && a.ID-1 == l.Alleles[i-1].ID {
				Aminus1_true = r.Allele(l.Alleles[i-1].ID).Height
			}

			r.AddAllele(Allele{
				ID:     a.ID,
				Height: truePeak(a.Height, l.Alleles[i+1].Height, Aminus1_true, ms, ps),
			})

			continue
		}

		// No allele in  plus stutter position but maybe in minus stutter position?
		// If so, this is the last allele in the current set.
		if i > 0 && a.ID-1 == l.Alleles[i-1].ID {
			r.AddAllele(Allele{
				ID:     a.ID,
				Height: a.Height - r.Allele(l.Alleles[i-1].ID).Height*ps,
			})
			continue
		}

		// No alleles in plus- or minus stutter position; lonely allele.
		r.AddAllele(a)
	}

	return r
}

// truePeak returns the true height of peak A, A_true, given its observed
// height, A_obs, the observed height of the successor peak, B_obs, as well as
// the minus- and plus stutter ratios, ms and ps. Aminus1_true is the true height
// (determined in the previous iteration of this function) of a putative
// predeccessor peak to the observed A (in -1 stutter position to A). Its
// contribution to A_obs is calculated using ps. If A is the first peak in a
// set of peaks, Aminus1_true = 0.
//
// Termininology: _obs:  observed peak height
//                _true: true peack height
//                _fwd:  forward stutter
//                _back: back stutter
//
// Solution for A_true = A_obs - B_fwd:
// -------------------------------------
//
// B_fwd = B_true * ms                           [ B_true = B_obs - A_back ]
//       = (B_obs - A_back) * ms                 [ A_back = ps * A_true ]
//       = (B_obs - ps * A_true) * ms            [ A_true = A_obs - B_fwd ]
//       = (B_obs - ps * (A_obs - B_fwd)) * ms
//       = (B_obs - ps*A_obs + ps*B_fwd) * ms
//       = ms*B_obs - ms*ps*A_obs + ms*ps*B_fwd
//
// B_fwd - ms*ps*B_fwd = ms*B_obs - ms*ps*A_obs
//    B_fwd(1 - ms*ps) = ms*B_obs - ms*ps*A_obs
//               B_fwd = (ms*B_obs - ms*ps*A_obs) / (1 - ms*ps)
//               B_fwd = ms(B_obs - ps*A_obs) / (1 - ms*ps)
//
// A_obs corrected for back stutter of previous peak = A_obs - Aminus1_true * ps.
//
// Hence, the function returns the evaluation of the term A_obs - B_fwd.
func truePeak(A_obs, B_obs, Aminus1_true, ms, ps float64) float64 {

	fmt.Println("(A_obs - Aminus1_true*ps):", A_obs-Aminus1_true*ps)
	fmt.Println("(ps*A_obs):", ps*A_obs)
	fmt.Println("(B_obs-ps*A_obs):", B_obs-ps*A_obs)
	fmt.Println("ms*(B_obs-ps*A_obs):", ms*(B_obs-ps*A_obs))
	fmt.Println("(1-ms*ps):", 1-ms*ps)

	return (A_obs - Aminus1_true*ps) - ms*(B_obs-ps*A_obs)/(1-ms*ps)
}
