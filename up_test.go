// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func Test_InferUnknownPersons(t *testing.T) {

	const (
		MinHomozygous   = 1500
		MajorCompRatio  = 2.9
		HeteroImbalance = 0.67
		WeakSignal      = 500
	)

	type test struct {
		inSample Sample
		want     []Sample
	}

	tests := []test{
		{
			Sample{
				ID:   "testinferUPs",
				Info: make(map[string]string),
				Loci: []Locus{
					{ID: "L1", Alleles: []Allele{{ID: 19, Height: 15980}}},
					{ID: "L2", Alleles: []Allele{{ID: 27, Height: 3043}, {ID: 28, Height: 15953}, {ID: 29, Height: 17927}, {ID: 30, Height: 2397}}},
					{ID: "L3", Alleles: []Allele{{ID: 15, Height: 2905}, {ID: 16, Height: 14099}, {ID: 18, Height: 18792}}},
					{ID: "L4", Alleles: []Allele{{ID: 6, Height: 8670}, {ID: 7, Height: 2121}, {ID: 9, Height: 2440}, {ID: 9.3, Height: 10618}}},
					{ID: "L5", Alleles: []Allele{{ID: 21, Height: 13812}, {ID: 23, Height: 4228}, {ID: 23.2, Height: 13300}}},
					{ID: "L6", Alleles: []Allele{{ID: 14, Height: 18289}, {ID: 17, Height: 15099}, {ID: 18, Height: 1863}}},
					{ID: "L7", Alleles: []Allele{{ID: 10, Height: 1819}, {ID: 13, Height: 14423}, {ID: 15, Height: 2221}, {ID: 16, Height: 11934}}},
					{ID: "L8", Alleles: []Allele{{ID: 13, Height: 3217}, {ID: 15, Height: 13922}, {ID: 21, Height: 2417}, {ID: 22, Height: 12689}}},
				},
			},
			[]Sample{
				{
					ID:     "testinferUPs::UP1",
					Info:   make(map[string]string),
					Source: "testinferUPs::MC",
					Loci: []Locus{
						{ID: "L1", Alleles: []Allele{{ID: 19}}},
						{ID: "L2", Alleles: []Allele{{ID: 28}, {ID: 29}}},
						{ID: "L3", Alleles: []Allele{{ID: 16}, {ID: 18}}},
						{ID: "L4", Alleles: []Allele{{ID: 6}, {ID: 9.3}}},
						{ID: "L5", Alleles: []Allele{{ID: 21}, {ID: 23.2}}},
						{ID: "L6", Alleles: []Allele{{ID: 14}, {ID: 17}}},
						{ID: "L7", Alleles: []Allele{{ID: 13}, {ID: 16}}},
						{ID: "L8", Alleles: []Allele{{ID: 15}, {ID: 22}}},
					},
					Kit: Kit{
						ID:   "unknown Kit",
						STRs: []STR{{ID: "L1"}, {ID: "L2"}, {ID: "L3"}, {ID: "L4"}, {ID: "L5"}, {ID: "L6"}, {ID: "L7"}, {ID: "L8"}}},
				}, {
					ID:     "testinferUPs::UP2",
					Info:   make(map[string]string),
					Source: "testinferUPs::MC::REM::MC",
					Loci: []Locus{
						{ID: "L2", Alleles: []Allele{{ID: 27}, {ID: 30}}},
						{ID: "L3", Alleles: []Allele{{ID: 15}}},
						{ID: "L4", Alleles: []Allele{{ID: 7}, {ID: 9}}},
						{ID: "L5", Alleles: []Allele{{ID: 23}}},
						{ID: "L6", Alleles: []Allele{{ID: 18}}},
						{ID: "L7", Alleles: []Allele{{ID: 10}, {ID: 15}}},
						{ID: "L8", Alleles: []Allele{{ID: 13}, {ID: 21}}},
					},
					Kit: Kit{
						ID:   "unknown Kit",
						STRs: []STR{{ID: "L2"}, {ID: "L3"}, {ID: "L4"}, {ID: "L5"}, {ID: "L6"}, {ID: "L7"}, {ID: "L8"}}},
				},
			},
		},
	}

	for i, tc := range tests {
		res := tc.inSample.InferUnknownPersons(HeteroImbalance, MajorCompRatio, MinHomozygous, WeakSignal)
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func Test_removeAlleles(t *testing.T) {

	type test struct {
		inLocus     Locus
		personLocus Locus
		want        Locus
	}

	tests := []test{
		{ // 1
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 18, Height: 998},
				{ID: 23, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2},
				{ID: 18}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 23, Height: 7623}}},
		},
		{ // 2
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 18, Height: 998},
				{ID: 23, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 23}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 18, Height: 998}}},
		},
		{ // 3
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 23, Height: 5790},
				{ID: 27, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 23},
				{ID: 27}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820}}},
		},
		{ // 4
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 23, Height: 5790},
				{ID: 27, Height: 7623}}},
			Locus{ID: "SE33"},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 23, Height: 5790},
				{ID: 27, Height: 7623}}},
		},
		{ // 5
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 9.2, Height: 820},
				{ID: 23, Height: 5790},
				{ID: 27, Height: 7623}}},
			Locus{ID: "SE33", Alleles: []Allele{
				{ID: 21.3},
				{ID: 23},
				{ID: 27}}},
			Locus{},
		},
	}

	for i, tc := range tests {
		res := tc.inLocus.removePersonFromLocus(tc.personLocus)
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestSample_Sex(t *testing.T) {

	type test struct {
		person Sample
		want   string
	}

	tests := []test{
		{
			Sample{Loci: []Locus{{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}}}},
			"male",
		},
		{
			Sample{Loci: []Locus{{ID: "AMEL", Alleles: []Allele{{ID: -1}}}}},
			"na",
		},
		{
			Sample{Loci: []Locus{{ID: "AMEL", Alleles: []Allele{}}}},
			"na",
		},
		{
			Sample{Loci: []Locus{{ID: "AMEL", Alleles: []Allele{{ID: -2}}}}},
			"female",
		},
		{
			Sample{Loci: []Locus{{ID: "TH01", Alleles: []Allele{{ID: 9, Height: 3021}, {ID: 9.3, Height: 1023}}}}},
			"na",
		},
		{
			Sample{Loci: []Locus{{ID: "TH01", Alleles: []Allele{{ID: 7, Height: 3021}, {ID: 9, Height: 3021}, {ID: 9.3, Height: 1023}}}}},
			"na",
		},
	}

	for i, tc := range tests {
		sex := tc.person.Sex()
		if tc.want != sex {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, sex)
		}
	}
}

// =============================================================================
func TestSample_SamePerson(t *testing.T) {

	type test struct {
		person1 Sample
		person2 Sample
		want    bool
	}

	tests := []test{
		{ // ===================== test 1 =====================
			Sample{Loci: []Locus{ // this is a UP
				{ID: "D21S11", Alleles: []Allele{{ID: 30, Height: 4049}, {ID: 31, Height: 5266}}},
				{ID: "VWA", Alleles: []Allele{{ID: 16, Height: 9300}, {ID: 17, Height: 7095}}},
				{ID: "TH01", Alleles: []Allele{{ID: 7, Height: 4705}, {ID: 9.3, Height: 6034}}},
				{ID: "FGA", Alleles: []Allele{{ID: 21, Height: 26191}}},
				{ID: "D3S1358", Alleles: []Allele{{ID: 14, Height: 9407}, {ID: 16, Height: 9737}}},
				{ID: "D8S1179", Alleles: []Allele{{ID: 10, Height: 15370}, {ID: 13, Height: 15918}}},
				{ID: "D18S51", Alleles: []Allele{{ID: 14, Height: 22161}}},
				{ID: "D1S1656", Alleles: []Allele{{ID: 11, Height: 6226}}},
				{ID: "D2S441", Alleles: []Allele{{ID: 11, Height: 3576}}},
				{ID: "D22S1045", Alleles: []Allele{{ID: 16, Height: 8607}}},
				{ID: "D16S539", Alleles: []Allele{{ID: 12, Height: 23888}}},
				{ID: "D2S1338", Alleles: []Allele{{ID: 17, Height: 15251}}},
				{ID: "D19S433", Alleles: []Allele{{ID: 14, Height: 12675}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2, Height: 12683}, {ID: -1, Height: 15568}}},
			}},
			Sample{Loci: []Locus{ // this is a reference profile
				{ID: "SE33", Alleles: []Allele{{ID: 16, Height: 759}, {ID: 18, Height: 738}}},
				{ID: "D21S11", Alleles: []Allele{{ID: 30, Height: 5227}}},
				{ID: "VWA", Alleles: []Allele{{ID: 14, Height: 3202}, {ID: 16, Height: 2517}}},
				{ID: "TH01", Alleles: []Allele{{ID: 6, Height: 1137}, {ID: 9.3, Height: 1119}}},
				{ID: "FGA", Alleles: []Allele{{ID: 20, Height: 905}, {ID: 22, Height: 884}}},
				{ID: "D3S1358", Alleles: []Allele{{ID: 16, Height: 3199}}},
				{ID: "D8S1179", Alleles: []Allele{{ID: 11, Height: 817}, {ID: 13, Height: 775}}},
				{ID: "D18S51", Alleles: []Allele{{ID: 16, Height: 1468}, {ID: 20, Height: 1450}}},
				{ID: "D1S1656", Alleles: []Allele{{ID: 12, Height: 1156}, {ID: 19.3, Height: 1133}}},
				{ID: "D2S441", Alleles: []Allele{{ID: 10, Height: 1512}, {ID: 14, Height: 1718}}},
				{ID: "D10S1248", Alleles: []Allele{{ID: 13, Height: 2167}, {ID: 14, Height: 2036}}},
				{ID: "D12S391", Alleles: []Allele{{ID: 16, Height: 1120}, {ID: 18, Height: 1042}}},
				{ID: "D22S1045", Alleles: []Allele{{ID: 15, Height: 1423}, {ID: 16, Height: 1283}}},
				{ID: "D16S539", Alleles: []Allele{{ID: 9, Height: 1280}, {ID: 12, Height: 1325}}},
				{ID: "D2S1338", Alleles: []Allele{{ID: 17, Height: 2107}, {ID: 25, Height: 1825}}},
				{ID: "D19S433", Alleles: []Allele{{ID: 15, Height: 635}, {ID: 15.2, Height: 646}}},
				{ID: "AMEL", Alleles: []Allele{{ID: -2, Height: 12683}}},
			}},
			false,
		},
		{ // ===================== test 2 =====================
			Sample{Loci: []Locus{ // this is a UP
				{ID: "D22S1045", Alleles: []Allele{{ID: 16, Height: 8607}, {ID: 17, Height: 8607}, {ID: 18, Height: 8607}}},
			}},
			Sample{Loci: []Locus{ // this is a reference profile
				{ID: "D22S1045", Alleles: []Allele{{ID: 16, Height: 8607}, {ID: 17, Height: 8607}}},
			}},
			false,
		},
	}

	for i, tc := range tests {
		res := SamePerson(tc.person1, tc.person2)
		if tc.want != res {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestSample_UniteUPs(t *testing.T) {

	type test struct {
		UPs  []Sample
		want Sample
	}

	tests := []test{
		{ // ----- 1
			[]Sample{
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 1203}, {ID: 15, Height: 1129}}},
						{ID: "VWA", Alleles: []Allele{{ID: 6, Height: 1203}}},
					},
				},
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 344}, {ID: 15, Height: 545}}},
						{ID: "VWA", Alleles: []Allele{{ID: 6, Height: 454}}},
					},
				},
			},
			Sample{Info: make(map[string]string),
				Loci: []Locus{
					{ID: "SE33", Alleles: []Allele{{ID: 13}, {ID: 15}}},
					{ID: "VWA", Alleles: []Allele{{ID: 6}}},
				},
				Kit: Kit{
					ID: "unknown Kit",
					STRs: []STR{
						{ID: "SE33"},
						{ID: "VWA"},
					},
				},
			},
		},
		{ // ----- 2
			[]Sample{
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 1203}, {ID: 15, Height: 1129}}},
						{ID: "VWA", Alleles: []Allele{{ID: 6, Height: 1203}}},
					},
				},
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 344}, {ID: 15, Height: 545}}},
						{ID: "VWA", Alleles: []Allele{{ID: 6, Height: 454}}},
					},
				},
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 344}}},
					},
				},
			},
			Sample{Info: make(map[string]string),
				Loci: []Locus{
					{ID: "SE33", Alleles: []Allele{{ID: 13}, {ID: 15}}},
					{ID: "VWA", Alleles: []Allele{{ID: 6}}},
				},
				Kit: Kit{
					ID: "unknown Kit",
					STRs: []STR{
						{ID: "SE33"},
						{ID: "VWA"},
					},
				},
			},
		},
		{ // ----- 3
			[]Sample{
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 1203}, {ID: 15, Height: 1129}}},
					}},
				{
					Loci: []Locus{
						{ID: "VWA", Alleles: []Allele{{ID: 6, Height: 454}}},
					}},
			},
			Sample{Info: make(map[string]string),
				Loci: []Locus{
					{ID: "SE33", Alleles: []Allele{{ID: 13}, {ID: 15}}},
					{ID: "VWA", Alleles: []Allele{{ID: 6}}},
				},
				Kit: Kit{
					ID: "unknown Kit",
					STRs: []STR{
						{ID: "SE33"},
						{ID: "VWA"},
					},
				},
			},
		},
		{ // ----- 4
			[]Sample{
				{
					Loci: []Locus{
						{ID: "SE33", Alleles: []Allele{{ID: 13, Height: 1203}, {ID: 14, Height: 1003}, {ID: 15, Height: 1129}}},
					},
				},
				{
					Loci: []Locus{
						{ID: "VWA", Alleles: []Allele{{ID: 6, Height: 454}}},
					},
				},
			},
			Sample{Info: make(map[string]string),
				Loci: []Locus{
					{ID: "VWA", Alleles: []Allele{{ID: 6}}},
				},
				Kit: Kit{
					ID: "unknown Kit",
					STRs: []STR{
						{ID: "VWA"},
					},
				},
			},
		},
		{ // ----- 5
			[]Sample{},
			Sample{},
		},
		{ // ----- 6
			[]Sample{
				{ID: "Sample 1"},
			},
			Sample{ID: "", Info: make(map[string]string), Source: "Sample 1",
				Kit: Kit{
					ID: "unknown Kit",
				},
			},
		},
		{ // ----- 7
			[]Sample{
				{ID: "Sample 1", Loci: []Locus{{ID: "SE33", Alleles: []Allele{{ID: 9.3}}}}},
			},
			Sample{Loci: []Locus{{ID: "SE33", Alleles: []Allele{{ID: 9.3}}}},
				Info: make(map[string]string), Source: "Sample 1",
				Kit: Kit{
					ID: "unknown Kit",
					STRs: []STR{
						{ID: "SE33"},
					},
				},
			},
		},
	}

	for i, tc := range tests {
		res := UniteUPs(tc.UPs, "")
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
