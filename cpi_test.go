// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"testing"
)

// =============================================================================
func TestLocus_PI_PE(t *testing.T) {

	type test struct {
		inLocus Locus
		inFreqs Freqs
		inTheta float64
		wantPI  float64
		wantPE  float64
	}

	tests := []test{
		{
			Locus{
				ID:      "VWA",
				Alleles: []Allele{{ID: 17}, {ID: 21}},
			},
			Freqs{
				Fmin: 0.001,
				Floci: []Flocus{
					{ID: "VWA", Falleles: []Fallele{{ID: 17, Freq: 0.1}, {ID: 21, Freq: 0.03}}},
				},
			},
			0,
			0.016900000000000002,
			0.983099999999999998,
		},
		{
			Locus{
				ID:      "VWA",
				Alleles: []Allele{{ID: 17}, {ID: 21}, {ID: 27.1}},
			},
			Freqs{
				Fmin: 0.001,
				Floci: []Flocus{
					{ID: "VWA", Falleles: []Fallele{{ID: 17, Freq: 0.1}, {ID: 21, Freq: 0.03}}},
				},
			},
			0,
			0.017161000000000003,
			0.982838999999999997,
		},
		{
			Locus{
				ID:      "VWA",
				Alleles: []Allele{{ID: 17}, {ID: 21}},
			},
			Freqs{
				Fmin:  0.001,
				Floci: []Flocus{},
			},
			0,
			0,
			0,
		},
		{
			Locus{
				ID:      "SE33",
				Alleles: []Allele{{ID: 17}, {ID: 21}, {ID: 21.3}},
			},
			Freqs{
				Fmin: 0.001,
				Floci: []Flocus{
					{ID: "SE33", Falleles: []Fallele{{ID: 17, Freq: 0.1}, {ID: 21, Freq: 0.2}, {ID: 21.3, Freq: 0.3}}},
				},
			},
			0.01,
			0.3624000000000001,
			0.6376,
		},
		{
			Locus{
				ID:      "SE33",
				Alleles: []Allele{{ID: 17}, {ID: 21}, {ID: 21.3}},
			},
			Freqs{
				Fmin: 0.01,
				Floci: []Flocus{
					{ID: "SE33", Falleles: []Fallele{{ID: 17, Freq: 0.1}, {ID: 21, Freq: 0.2}, {ID: 21.3, Freq: 0.3}}},
				},
			},
			0.03,
			0.3672000000000001,
			0.6327999999999999,
		},
	}

	for i, tc := range tests {
		resPI := tc.inLocus.PI(tc.inFreqs, tc.inTheta)
		if tc.wantPI != resPI {
			t.Fatalf("test %d (PI): expected: %v, got: %v", i+1, tc.wantPI,
				resPI)
		}

		resPE := tc.inLocus.PE(tc.inFreqs, tc.inTheta)
		if tc.wantPE != resPE {
			t.Fatalf("test %d (PE): expected: %v, got: %v", i+1, tc.wantPE,
				resPE)
		}
	}
}

// =============================================================================
func TestLocus_CPI_CPE_RMNE(t *testing.T) {

	type test struct {
		inSample Sample
		inFreqs  Freqs
		inTheta  float64
		wantCPI  float64
		wantCPE  float64
		wantRMNE float64
	}

	tests := []test{
		{
			Sample{
				Loci: []Locus{
					{ID: "VWA", Alleles: []Allele{{ID: 17}, {ID: 21}}},
					{ID: "FGA", Alleles: []Allele{{ID: 20.1}, {ID: 31}}},
					{ID: "DYS391", Alleles: []Allele{{ID: 12}}}},
			},
			Freqs{
				Fmin: 0.001,
				Floci: []Flocus{
					{ID: "VWA", Falleles: []Fallele{{ID: 17, Freq: 0.1}, {ID: 21, Freq: 0.03}}},
					{ID: "FGA", Falleles: []Fallele{{ID: 20.1, Freq: 0.21}, {ID: 31, Freq: 0.17}}},
				},
			},
			0,
			0.0024403600000000004,
			0.9975596399999999996,
			409.77560687767374,
		},
		{
			Sample{
				Loci: []Locus{
					{ID: "DYS391", Alleles: []Allele{{ID: 12}}}},
			},
			Freqs{
				Fmin: 0.001,
				Floci: []Flocus{
					{ID: "VWA", Falleles: []Fallele{{ID: 17, Freq: 0.1}, {ID: 21, Freq: 0.03}}},
					{ID: "FGA", Falleles: []Fallele{{ID: 20.1, Freq: 0.21}, {ID: 31, Freq: 0.17}}},
				},
			},
			0,
			0,
			1,
			0,
		},
	}

	for i, tc := range tests {
		resCPI := tc.inSample.CPI(tc.inFreqs, tc.inTheta)
		if tc.wantCPI != resCPI {
			t.Fatalf("test %d (CPI): expected: %v, got: %v", i+1, tc.wantCPI,
				resCPI)
		}

		resCPE := tc.inSample.CPE(tc.inFreqs, tc.inTheta)
		if tc.wantCPE != resCPE {
			t.Fatalf("test %d (CPE): expected: %v, got: %v", i+1, tc.wantCPE,
				resCPE)
		}

		resRMNE := tc.inSample.RMNE(tc.inFreqs, tc.inTheta)
		if tc.wantRMNE != resRMNE {
			t.Fatalf("test %d (RMNE): expected: %v, got: %v", i+1, tc.wantRMNE,
				resRMNE)
		}
	}
}
