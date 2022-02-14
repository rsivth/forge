// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func TestLocus_LocusCoordinates(t *testing.T) {

	type test struct {
		inLocus Locus
		want    GenomicCoordinates
	}

	tests := []test{
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			GenomicCoordinates{Chr: 4, Start: 155866000},
		},
		{
			Locus{ID: "Noname", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			GenomicCoordinates{},
		},
		{
			Locus{ID: "DYS390", Alleles: []Allele{{ID: 12}}},
			GenomicCoordinates{Chr: -1},
		},
	}

	for i, tc := range tests {
		res := LocusCoordinates(tc.inLocus.ID)
		if !reflect.DeepEqual(res, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
