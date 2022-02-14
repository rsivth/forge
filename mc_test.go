// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func Test_mcLocus(t *testing.T) {

	type test struct {
		inLocus Locus
		want    Locus
	}

	tests := []test{
		{ // 1
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 18, Height: 998},
				{ID: 23, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 23, Height: 7623}}},
		},
		{ // 2
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 23, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 23, Height: 7623}}},
		},
		{ // 3
			Locus{ID: "IQCS", Alleles: []Allele{
				{ID: 9.2, Height: 820}}},
			Locus{ID: "IQCS"},
		},
		{ // 4
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 8020},
				{ID: 18, Height: 998},
				{ID: 23, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 8020},
				{ID: 23, Height: 7623}}},
		},
		{ // 5
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 8020},
				{ID: 18, Height: 9098},
				{ID: 23, Height: 7623}}},
			Locus{ID: "SE33"},
		},
	}

	for i, tc := range tests {
		res := tc.inLocus.MajorComponent(0.67, 2.9, 1500, 500)
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
