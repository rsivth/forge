// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"math"
	"testing"
)

// =============================================================================
func Test_weightedAlleleDraw(t *testing.T) {

	type test struct {
		inFlocus     Flocus
		wantFalleles []Fallele
	}

	tests := []test{
		{
			Flocus{ID: "TH01", Falleles: []Fallele{{ID: 6, Freq: 0.2}, {ID: 9.3, Freq: 0.8}}},
			[]Fallele{{ID: 6, Freq: 0.2}, {ID: 9.3, Freq: 0.8}},
		},
		{
			Flocus{ID: "VWA", Falleles: []Fallele{{ID: 17, Freq: 0.2}, {ID: 17.2, Freq: 0.75}, {ID: 29, Freq: 0.05}}},
			[]Fallele{{ID: 17, Freq: 0.2}, {ID: 17.2, Freq: 0.75}, {ID: 29, Freq: 0.05}},
		},
		{
			Flocus{ID: "VWA"},
			[]Fallele{},
		},
	}

	numOfSim := 10000.0 // float64 makes calc easier

	for i, tc := range tests {

		simFreq := make(map[float64]float64)
		for x := 0; x < int(numOfSim); x++ {
			simFreq[tc.inFlocus.weightedAlleleDraw().ID]++
		}

		var resFalleles []Fallele
		for id, freq := range simFreq {
			resFalleles = append(resFalleles, Fallele{id, freq / numOfSim})
		}

		// sum up the simulations; must sum up to 1
		var sumFreq float64
		for _, r := range resFalleles {
			sumFreq += r.Freq
		}

		// sumFreq is sometimes 0.9999999999999999 (e.g. second test case) which appears to be
		// a floating point issue. Thus we accept this small deviation from 1.
		if (sumFreq < 0.999999999999999 && sumFreq <= 1) && len(resFalleles) != 0 {
			t.Fatalf("test %d (sum frequencies): expected: %v, got: %v", i+1, 1, sumFreq)
		}

		percentError := 0.01 // the accepted margin of error between expected and simulated
		for _, wf := range tc.wantFalleles {
			var found bool
			for _, rf := range resFalleles {
				if wf.ID == rf.ID {
					found = true
					if math.Abs(wf.Freq-rf.Freq) > percentError {
						t.Fatalf("test %d (indiv. freqs %v): expected: %v, got: %v", i+1, wf.ID, wf.Freq, rf.Freq)
					}
				}
			}
			if !found {
				t.Fatalf("test %d: cannot find %v in results", i+1, wf.ID)
			}
		}
	}
}
