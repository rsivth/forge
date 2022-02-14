// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"testing"
)

// =============================================================================
/*func TestSample_InferKit(t *testing.T) {

	conf := Config{
		KitFolder: "",
	}
	type test struct {
		inSample Sample
		want     string // kit name
	}

	tests := []test{
		{
			Sample{Loci: []Locus{{ID: "SE33"}, {ID: "VWA"}, {ID: "TH01"}}},
			"PPL ESX17",
		},
		{
			Sample{Loci: []Locus{{ID: "FGA"}, {ID: "D3S1358"}, {ID: "SE33"}}},
			"NGM-Detect",
		},
		{
			Sample{Loci: []Locus{{ID: "DYS19"}, {ID: "DYS391"}, {ID: "DYS448"}}},
			"PPL Y23",
		},
		{
			Sample{Loci: []Locus{{ID: "SE33"}, {ID: "VWA"}, {ID: "DYS385"}}},
			"unknown Kit",
		},
	}

	for i, tc := range tests {
		tc.inSample.InferKit()
		if tc.inSample.Kit.ID != tc.want {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, tc.inSample.Kit.ID)
		}
	}
}
*/
// =============================================================================
func TestKit_HasSTR(t *testing.T) {

	type test struct {
		inKit Kit
		inSTR string
		want  bool
	}

	tests := []test{
		{
			Kit{
				ID: "Ifiler Plus",
				STRs: []STR{
					{ID: "D8S1179", Dye: "blue"},
					{ID: "D21S11", Dye: "blue"},
					{ID: "D7S820", Dye: "blue"},
					{ID: "CSF1PO", Dye: "blue"},
					{ID: "D3S1358", Dye: "green"},
					{ID: "TH01", Dye: "green"},
					{ID: "D13S317", Dye: "green"},
					{ID: "D16S539", Dye: "green"},
					{ID: "D2S1338", Dye: "green"},
					{ID: "D19S433", Dye: "yellow"},
					{ID: "VWA", Dye: "yellow"},
					{ID: "TPOX", Dye: "yellow"},
					{ID: "D18S51", Dye: "yellow"},
					{ID: "AMEL", Dye: "red"},
					{ID: "D5S818", Dye: "red"},
					{ID: "FGA", Dye: "red"},
				},
			},
			"D19S433",
			true,
		},
		{
			Kit{
				ID: "Ifiler Plus",
				STRs: []STR{
					{ID: "D8S1179", Dye: "blue"},
					{ID: "D21S11", Dye: "blue"},
					{ID: "D7S820", Dye: "blue"},
					{ID: "CSF1PO", Dye: "blue"},
					{ID: "D3S1358", Dye: "green"},
					{ID: "TH01", Dye: "green"},
					{ID: "D13S317", Dye: "green"},
					{ID: "D16S539", Dye: "green"},
					{ID: "D2S1338", Dye: "green"},
					{ID: "D19S433", Dye: "yellow"},
					{ID: "VWA", Dye: "yellow"},
					{ID: "TPOX", Dye: "yellow"},
					{ID: "D18S51", Dye: "yellow"},
					{ID: "AMEL", Dye: "red"},
					{ID: "D5S818", Dye: "red"},
					{ID: "FGA", Dye: "red"},
				},
			},
			"DYS437",
			false,
		},
	}

	for i, tc := range tests {
		res := tc.inKit.HasSTR(tc.inSTR)
		if res != tc.want {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
