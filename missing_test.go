// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func TestSample_MissingFrom_MissingFromInt(t *testing.T) {

	type test struct {
		inSamplePerson Sample
		inSampleStain  Sample
		wantSample     Sample
		wantInt        int
	}

	tests := []test{
		{
			Sample{ID: "person", Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}}},
				{ID: "VWA", Alleles: []Allele{{ID: 9}, {ID: 11}}},
				{ID: "TH01", Alleles: []Allele{{ID: 17}, {ID: 17.2}}},
			}},
			Sample{ID: "stain", Loci: []Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 29}, {ID: 29.3}, {ID: 31}}},
				{ID: "VWA", Alleles: []Allele{{ID: 7}, {ID: 11}, {ID: 17}}},
				{ID: "Penta", Alleles: []Allele{{ID: 4}, {ID: 7.1}}},
			}},
			Sample{ID: "Missing_person_from_stain", Loci: []Locus{
				{ID: "VWA", Alleles: []Allele{{ID: 9}}},
			}},
			1,
		},
	}

	for i, tc := range tests {
		missing := tc.inSamplePerson.MissingFrom(tc.inSampleStain)
		if !reflect.DeepEqual(missing, tc.wantSample) {
			t.Fatalf("test %d (Sample): expected: %v, got: %v", i+1, tc.wantSample, missing)
		}

		missingInt := tc.inSamplePerson.MissingFromInt(tc.inSampleStain)
		if missingInt != tc.wantInt {
			t.Fatalf("test %d (Int): expected: %v, got: %v", i+1, tc.wantInt, missingInt)
		}
	}
}
