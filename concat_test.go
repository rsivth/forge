// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"sort"
	"testing"
)

// =============================================================================
func Test_concatLoci(t *testing.T) {

	type test struct {
		inLoci   []Locus
		wantComp Locus
		wantCons Locus
	}

	tests := []test{
		{ // 1
			[]Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 9.2}, {ID: 18}, {ID: 23}}},
				{ID: "SE33", Alleles: []Allele{{ID: 9.2}, {ID: 18}, {ID: 23}}},
				{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 23}, {ID: 25.2}}},
			},
			Locus{ID: "SE33", Alleles: []Allele{{ID: 9.2}, {ID: 18}, {ID: 23}, {ID: 25.2}}},
			Locus{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 23}}},
		},
		{ // 2
			[]Locus{
				{ID: "SE33", Alleles: []Allele{{ID: 9.2}, {ID: 18}}},
				{ID: "SE33", Alleles: []Allele{{ID: 11.1}, {ID: 23}}},
				{ID: "SE33", Alleles: []Allele{{ID: 25.2}}},
			},
			Locus{ID: "SE33", Alleles: []Allele{{ID: 9.2}, {ID: 11.1}, {ID: 18}, {ID: 23}, {ID: 25.2}}},
			Locus{ID: "SE33"},
		},
		{ // 3
			[]Locus{
				{ID: "SE33"},
			},
			Locus{ID: "SE33"},
			Locus{ID: "SE33"},
		},
	}

	for i, tc := range tests {
		gotComp := concatLoci(tc.inLoci, COMPOSITE)
		if !reflect.DeepEqual(tc.wantComp, gotComp) {
			t.Fatalf("test %d (composite): expected: %v, got: %v", i+1, tc.wantComp, gotComp)
		}

		gotCons := concatLoci(tc.inLoci, CONSENSUS)
		if !reflect.DeepEqual(tc.wantCons, gotCons) {
			t.Fatalf("test %d (consensus): expected: %v, got: %v", i+1, tc.wantCons, gotCons)
		}

	}
}

// =============================================================================
func Test_concatSamples(t *testing.T) {

	type test struct {
		inSamples []Sample
		wantComp  Sample
		wantCons  Sample
	}

	tests := []test{
		{ // test 1 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 17.2}, {ID: 18}, {ID: 23}}},
						{ID: "vWA", Alleles: []Allele{{ID: 6}, {ID: 9}, {ID: 9.3}}},
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}}},
					},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 11}, {ID: 18}, {ID: 23}}},
						{ID: "vWA", Alleles: []Allele{{ID: 7}}},
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
					},
				},
				{
					ID:   "test3",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 11}, {ID: 18}, {ID: 23}, {ID: 28}}},
						{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 31}}},
					},
				},
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
					{ID: "SE33", Alleles: []Allele{{ID: 11}, {ID: 17.2}, {ID: 18}, {ID: 23}, {ID: 28}}},
					{ID: "VWA", Alleles: []Allele{{ID: 6}, {ID: 7}, {ID: 9}, {ID: 9.3}}},
				},
				Source: "composite::all_linkage::test1::test2::test3",
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "FGA", Alleles: []Allele{{ID: 31}}},
					{ID: "SE33", Alleles: []Allele{{ID: 18}, {ID: 23}}},
					{ID: "VWA"},
				},
				Source: "consensus::all_linkage::test1::test2::test3",
			},
		},
		{ // test 2 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{},
				},
			},
			Sample{
				ID:     "test1::test2",
				Info:   make(map[string]string),
				Source: "composite::all_linkage::test1::test2",
			},
			Sample{
				ID:     "test1::test2",
				Info:   make(map[string]string),
				Source: "consensus::all_linkage::test1::test2",
			},
		},
		{ // test 3 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}}},
					},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
					},
				},
				{
					ID:   "test3",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 31}}},
					},
				},
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
				},
				Source: "composite::all_linkage::test1::test2::test3",
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "FGA", Alleles: []Allele{{ID: 31}}},
				},
				Source: "consensus::all_linkage::test1::test2::test3",
			},
		},
		{ // test 4 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 10}, {ID: 11}}},
					},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 9}, {ID: 11}}},
					},
				},
				{
					ID:   "test3",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 31}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 11}}},
					},
				},
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "DYS518", Alleles: []Allele{{ID: 9}, {ID: 10}, {ID: 11}}},
					{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
				},
				Source: "composite::all_linkage::test1::test2::test3",
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "DYS518", Alleles: []Allele{{ID: 11}}},
					{ID: "FGA", Alleles: []Allele{{ID: 31}}},
				},
				Source: "consensus::all_linkage::test1::test2::test3",
			},
		},
		{ // test 5 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 10}, {ID: 11}}},
						{ID: "DYS393", Alleles: []Allele{{ID: 6}, {ID: 7}}},
					},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 9}, {ID: 11}}},
					},
				},
				{
					ID:   "test3",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 31}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 7}, {ID: 11}}},
						{ID: "DYS393", Alleles: []Allele{{ID: 6}, {ID: 9.3}}},
					},
				},
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "DYS393", Alleles: []Allele{{ID: 6}, {ID: 7}, {ID: 9.3}}},
					{ID: "DYS518", Alleles: []Allele{{ID: 7}, {ID: 9}, {ID: 10}, {ID: 11}}},
					{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
				},
				Source: "composite::all_linkage::test1::test2::test3",
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "DYS393"},
					{ID: "DYS518", Alleles: []Allele{{ID: 11}}},
					{ID: "FGA", Alleles: []Allele{{ID: 31}}},
				},
				Source: "consensus::all_linkage::test1::test2::test3",
			},
		},
		{ // test 6 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 10}, {ID: 11}}},
					},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
						{ID: "DYS518", Alleles: []Allele{{ID: 9}, {ID: 11}}},
					},
				},
				{
					ID:   "test3",
					Info: make(map[string]string),
					Loci: []Locus{
						{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 31}}},
					},
				},
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "DYS518", Alleles: []Allele{{ID: 9}, {ID: 10}, {ID: 11}}},
					{ID: "FGA", Alleles: []Allele{{ID: 29}, {ID: 30}, {ID: 31}, {ID: 33}, {ID: 34}}},
				},
				Source: "composite::all_linkage::test1::test2::test3",
			},
			Sample{
				ID:   "test1::test2::test3",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "DYS518"},
					{ID: "FGA", Alleles: []Allele{{ID: 31}}},
				},
				Source: "consensus::all_linkage::test1::test2::test3",
			},
		},
		{ // test 7 --------------------------------------------------------
			[]Sample{},
			Sample{},
			Sample{},
		},
		{ // test 8 --------------------------------------------------------
			[]Sample{
				{
					ID:   "test1",
					Info: make(map[string]string),
					Loci: []Locus{{ID: "SE33"}},
				},
				{
					ID:   "test2",
					Info: make(map[string]string),
					Loci: []Locus{{ID: "VWA"}},
				},
			},
			Sample{
				ID:     "test1::test2",
				Info:   make(map[string]string),
				Source: "composite::all_linkage::test1::test2",
				Loci:   []Locus{{ID: "SE33"}, {ID: "VWA"}},
			},
			Sample{
				ID:     "test1::test2",
				Info:   make(map[string]string),
				Source: "consensus::all_linkage::test1::test2",
				Loci:   []Locus{{ID: "SE33"}, {ID: "VWA"}},
			},
		},
	}

	// TODO: add tests for y und x linkage

	for i, tc := range tests {
		resComp := Composite(tc.inSamples, ALLLINKAGE)
		resComp.sortLocusIDs()
		if !reflect.DeepEqual(tc.wantComp, resComp) {
			t.Fatalf("test %d (composite): expected: %v, got: %v", i+1,
				tc.wantComp, resComp)
		}

		resCons := Consensus(tc.inSamples, ALLLINKAGE)
		resCons.sortLocusIDs()
		if !reflect.DeepEqual(tc.wantCons, resCons) {
			t.Fatalf("test %d (consensus): expected: %v, got: %v", i+1,
				tc.wantCons, resCons)
		}
	}
}

// apparently we need this for the comparability of the tests.
func (s Sample) sortLocusIDs() {
	sort.Slice(s.Loci, func(i, j int) bool {
		return s.Loci[i].ID < s.Loci[j].ID
	})
}

// =============================================================================
func Test_concatID(t *testing.T) {

	type test struct {
		inSamples []Sample
		wantID    string
	}

	tests := []test{
		{ // test 1 --------------------------------------------------------
			[]Sample{
				{ID: "A"},
				{ID: "B"},
				{ID: "C"},
			},
			"A::B::C",
		},
		{ // test 1 --------------------------------------------------------
			[]Sample{
				{ID: "A"},
				{ID: "A"},
				{ID: "A"},
			},
			"A",
		},
	}

	for i, tc := range tests {
		r := concatID(tc.inSamples)
		if tc.wantID != r {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.wantID, r)
		}
	}
}
