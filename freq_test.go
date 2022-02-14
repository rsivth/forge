// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func TestFreqs_Flocus(t *testing.T) {

	type test struct {
		inFreqs Freqs
		inID    string
		want    Flocus
	}

	tests := []test{
		{
			Freqs{
				Floci: []Flocus{
					{ID: "SE33", Falleles: []Fallele{{ID: 9.3}}},
					{ID: "VWA", Falleles: []Fallele{{ID: 17}}},
					{ID: "FGA", Falleles: []Fallele{{ID: 22.3}}},
				},
			},
			"SE33",
			Flocus{ID: "SE33", Falleles: []Fallele{{ID: 9.3}}},
		},
	}

	for i, tc := range tests {
		res := tc.inFreqs.Flocus(tc.inID)
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestFreqs_HasFlocus(t *testing.T) {

	type test struct {
		inFreqs Freqs
		inID    string
		want    bool
	}

	tests := []test{
		{
			Freqs{
				Floci: []Flocus{
					// HasFlocus returns false when no alleles present
					{ID: "SE33", Falleles: []Fallele{{ID: 9.3}}},
					{ID: "VWA", Falleles: []Fallele{{ID: 17}}},
					{ID: "FGA", Falleles: []Fallele{{ID: 22.3}}},
				},
			},
			"SE33",
			true,
		},
		{
			Freqs{
				Floci: []Flocus{
					// HasFlocus returns false when no alleles present
					{ID: "SE33", Falleles: []Fallele{{ID: 9.3}}},
					{ID: "VWA", Falleles: []Fallele{{ID: 17}}},
					{ID: "FGA", Falleles: []Fallele{{ID: 22.3}}},
				},
			},
			"AMEL",
			false,
		},
	}

	for i, tc := range tests {
		res := tc.inFreqs.HasFlocus(tc.inID)
		if tc.want != res {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestFlocus_Fallele(t *testing.T) {

	type test struct {
		inFlocus Flocus
		inID     float64
		want     Fallele
	}

	tests := []test{
		{
			Flocus{ID: "SE33", Falleles: []Fallele{{ID: 9}, {ID: 9.3}}},
			9.3,
			Fallele{ID: 9.3},
		},
		{
			Flocus{ID: "SE33", Falleles: []Fallele{{ID: 9}, {ID: 9.3}}},
			6,
			Fallele{},
		}}

	for i, tc := range tests {
		res := tc.inFlocus.Fallele(tc.inID)
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestFlocus_HasFallele(t *testing.T) {

	type test struct {
		inFlocus Flocus
		inID     float64
		want     bool
	}

	tests := []test{
		{
			Flocus{ID: "SE33", Falleles: []Fallele{{ID: 9}, {ID: 9.3}}},
			9.3,
			true,
		},
		{
			Flocus{ID: "SE33", Falleles: []Fallele{{ID: 9}, {ID: 9.3}}},
			6,
			false,
		}}

	for i, tc := range tests {
		res := tc.inFlocus.HasFallele(tc.inID)
		if tc.want != res {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
