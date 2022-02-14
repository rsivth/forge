// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func Test_NewSample(t *testing.T) {

	type test struct {
		inID string
		want Sample
	}

	tests := []test{
		{"stain 1", Sample{ID: "stain 1", Info: make(map[string]string)}},
		{"", Sample{ID: "", Info: make(map[string]string)}},
	}

	for i, tc := range tests {
		s := NewSample(tc.inID, "")
		if !reflect.DeepEqual(s, tc.want) {
			t.Fatalf("test %d (sample): expected: %v, got: %v", i+1, tc.want, s)
		}
	}
}

// =============================================================================
func TestSample_AddLocus(t *testing.T) {

	type test struct {
		inSample Sample
		inLocus  Locus
		want     Sample
	}

	tests := []test{
		{
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
			}},
			Locus{ID: "VWA", Alleles: []Allele{{ID: 21}, {ID: 23}}},
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "VWA", Alleles: []Allele{{ID: 21}, {ID: 23}}},
			}},
		},
		{
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
			}},
			Locus{},
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
			}},
		},
		{
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
			}},
			// the empty locus must still be sampled to get the kit right
			Locus{ID: "VWA"},
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "VWA"},
			}},
		},
		{
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18, Height: 100}, {ID: 29, Height: 100}}},
			}},
			// the empty locus must still be sampled to get the kit right
			Locus{ID: "SE33", Alleles: []Allele{{ID: 18, Height: 430}, {ID: 29, Height: 430}}},
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18, Height: 100}, {ID: 29, Height: 100}}},
			}},
		},
	}

	for i, tc := range tests {
		tc.inSample.AddLocus(tc.inLocus)
		if !reflect.DeepEqual(tc.inSample, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, tc.inSample)
		}
	}
}

// =============================================================================
func TestSample_Locus(t *testing.T) {

	type test struct {
		inSample Sample
		inID     string
		want     Locus
	}

	tests := []test{
		{ // 1
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "VWA", Alleles: []Allele{{ID: 21}, {ID: 23}}},
			}},
			"VWA",
			Locus{ID: "VWA", Alleles: []Allele{{ID: 21}, {ID: 23}}},
		},
		{ // 2
			Sample{Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "VWA", Alleles: []Allele{{ID: 21}, {ID: 23}}},
			}},
			"FGA",
			Locus{},
		},
	}

	for i, tc := range tests {
		locus := tc.inSample.Locus(tc.inID)
		if !reflect.DeepEqual(locus, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, locus)
		}
	}
}

// =============================================================================
func TestSample_HasLocus(t *testing.T) {

	type test struct {
		inSample Sample
		inLocus  string
		want     bool
	}

	tests := []test{
		{
			Sample{
				Loci: []Locus{
					{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
					{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				}},
			"SE33",
			true,
		},
		{
			Sample{
				Loci: []Locus{
					{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
					{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				}},
			"FGA",
			false,
		},
	}

	for i, tc := range tests {
		res := tc.inSample.HasLocus(tc.inLocus)
		if res != tc.want {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestSample_MaxAlleles(t *testing.T) {

	type test struct {
		inSample Sample
		want     int
	}

	tests := []test{
		{
			Sample{Loci: []Locus{
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "SE33", Alleles: []Allele{{ID: 17}, {ID: 21.2}, {ID: 21.2}}},
			}},
			3,
		},
		{
			Sample{Loci: []Locus{
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "SE33", Alleles: []Allele{{ID: 21.2}, {ID: 21.2}}},
			}},
			2,
		},
		{
			Sample{Loci: []Locus{}},
			0,
		},
	}

	for i, tc := range tests {
		gotInt := tc.inSample.MaxAlleles()
		if tc.want != gotInt {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, gotInt)
		}
	}
}

// =============================================================================
func TestSample_MinContributor(t *testing.T) {

	type test struct {
		inSample Sample
		want     int
	}

	tests := []test{
		{
			Sample{Loci: []Locus{
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "SE33", Alleles: []Allele{{ID: 17}, {ID: 21.2}, {ID: 21.2}}},
			}},
			2,
		},
		{
			Sample{Loci: []Locus{
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "SE33", Alleles: []Allele{{ID: 17}, {ID: 21.2}, {ID: 21.2}}},
				{ID: "TH01", Alleles: []Allele{{ID: 6}, {ID: 6.3}, {ID: 7}, {ID: 9.3}}},
				{ID: "FGA", Alleles: []Allele{{ID: 17}, {ID: 21.2}, {ID: 21.2}}},
			}},
			2,
		},
		{
			Sample{Loci: []Locus{
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "SE33", Alleles: []Allele{{ID: 17}, {ID: 21.2}, {ID: 21.2}}},
				{ID: "TH01", Alleles: []Allele{{ID: 6}, {ID: 6.3}, {ID: 7}, {ID: 9}, {ID: 9.3}}},
				{ID: "FGA", Alleles: []Allele{{ID: 17}, {ID: 21.2}, {ID: 21.2}}},
			}},
			3,
		},
		{
			Sample{Loci: []Locus{
				{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
				{ID: "SE33", Alleles: []Allele{{ID: 21.2}, {ID: 21.2}}},
			}},
			1,
		},
		{
			Sample{Loci: []Locus{}},
			0,
		},
	}

	for i, tc := range tests {
		gotInt := tc.inSample.MinContributor()
		if tc.want != gotInt {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, gotInt)
		}
	}
}
