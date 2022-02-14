package forge

import (
	"reflect"
	"testing"
)

// provedIT data was downloaded from https://lftdi.camden.rutgers.edu/provedit/
var provedITheader = []string{"Sample File", "Marker", "Dye", "Allele 1",
	"Size 1", "Height 1", "Allele 2", "Size 2", "Height 2", "Allele 3",
	"Size 3", "Height 3", "Allele 4", "Size 4", "Height 4", "Allele 5",
	"Size 5", "Height 5", "Allele 6", "Size 6", "Height 6", "Allele 7",
	"Size 7", "Height 7", "Allele 8", "Size 8", "Height 8", "Allele 9",
	"Size 9", "Height 9", "Allele 10", "Size 10", "Height 10", "Allele 11",
	"Size 11", "Height 11", "Allele 12", "Size 12", "Height 12", "Allele 13",
	"Size 13", "Height 13", "Allele 14", "Size 14", "Height 14", "Allele 15",
	"Size 15", "Height 15", "Allele 16", "Size 16", "Height 16", "Allele 17",
	"Size 17", "Height 17", "Allele 18", "Size 18", "Height 18", "Allele 19",
	"Size 19", "Height 19", "Allele 20", "Size 20", "Height 20", "Allele 21",
	"Size 21", "Height 21", "Allele 22", "Size 22", "Height 22", "Allele 23",
	"Size 23", "Height 23", "Allele 24", "Size 24", "Height 24", "Allele 25",
	"Size 25", "Height 25", "Allele 26", "Size 26", "Height 26", "Allele 27",
	"Size 27", "Height 27", "Allele 28", "Size 28", "Height 28", "Allele 29",
	"Size 29", "Height 29", "Allele 30", "Size 30", "Height 30", "Allele 31",
	"Size 31", "Height 31", "Allele 32", "Size 32", "Height 32", "Allele 33",
	"Size 33", "Height 33", "Allele 34", "Size 34", "Height 34", "Allele 35",
	"Size 35", "Height 35", "Allele 36", "Size 36", "Height 36", "Allele 37",
	"Size 37", "Height 37", "Allele 38", "Size 38", "Height 38", "Allele 39",
	"Size 39", "Height 39", "Allele 40", "Size 40", "Height 40", "Allele 41",
	"Size 41", "Height 41", "Allele 42", "Size 42", "Height 42", "Allele 43",
	"Size 43", "Height 43", "Allele 44", "Size 44", "Height 44", "Allele 45",
	"Size 45", "Height 45", "Allele 46", "Size 46", "Height 46", "Allele 47",
	"Size 47", "Height 47", "Allele 48", "Size 48", "Height 48", "Allele 49",
	"Size 49", "Height 49", "Allele 50", "Size 50", "Height 50", "Allele 51",
	"Size 51", "Height 51", "Allele 52", "Size 52", "Height 52", "Allele 53",
	"Size 53", "Height 53", "Allele 54", "Size 54", "Height 54", "Allele 55",
	"Size 55", "Height 55", "Allele 56", "Size 56", "Height 56", "Allele 57",
	"Size 57", "Height 57", "Allele 58", "Size 58", "Height 58", "Allele 59",
	"Size 59", "Height 59", "Allele 60", "Size 60", "Height 60", "Allele 61",
	"Size 61", "Height 61", "Allele 62", "Size 62", "Height 62", "Allele 63",
	"Size 63", "Height 63", "Allele 64", "Size 64", "Height 64", "Allele 65",
	"Size 65", "Height 65", "Allele 66", "Size 66", "Height 66", "Allele 67",
	"Size 67", "Height 67", "Allele 68", "Size 68", "Height 68", "Allele 69",
	"Size 69", "Height 69", "Allele 70", "Size 70", "Height 70", "Allele 71",
	"Size 71", "Height 71", "Allele 72", "Size 72", "Height 72", "Allele 73",
	"Size 73", "Height 73", "Allele 74", "Size 74", "Height 74", "Allele 75",
	"Size 75", "Height 75", "Allele 76", "Size 76", "Height 76", "Allele 77",
	"Size 77", "Height 77", "Allele 78", "Size 78", "Height 78", "Allele 79",
	"Size 79", "Height 79", "Allele 80", "Size 80", "Height 80", "Allele 81",
	"Size 81", "Height 81", "Allele 82", "Size 82", "Height 82", "Allele 83",
	"Size 83", "Height 83", "Allele 84", "Size 84", "Height 84", "Allele 85",
	"Size 85", "Height 85", "Allele 86", "Size 86", "Height 86", "Allele 87",
	"Size 87", "Height 87", "Allele 88", "Size 88", "Height 88", "Allele 89",
	"Size 89", "Height 89", "Allele 90", "Size 90", "Height 90", "Allele 91",
	"Size 91", "Height 91", "Allele 92", "Size 92", "Height 92", "Allele 93",
	"Size 93", "Height 93", "Allele 94", "Size 94", "Height 94", "Allele 95",
	"Size 95", "Height 95", "Allele 96", "Size 96", "Height 96", "Allele 97",
	"Size 97", "Height 97", "Allele 98", "Size 98", "Height 98", "Allele 99",
	"Size 99", "Height 99", "Allele 100", "Size 100", "Height 100"}

// =============================================================================
func Test_indexHeader_isCollated_allelePos(t *testing.T) {

	type test struct {
		sID           string
		header        []string
		info          []string
		n             int // n_th Allele, we want the column number
		wantIndex     index
		wantCollated  bool
		wantAllelePos int // zero-based col number where allele n can be found
	}

	tests := []test{
		{
			"Sample File",
			provedITheader,
			[]string{"Dye"},
			87, // n_the allele
			index{
				"Sample File":  0,
				"Marker":       1,
				"Allele 1":     3,
				"Dye":          2,
				"AreaOffset":   -1,
				"HeightOffset": 2,
				"SizeOffset":   1,
				"NoOfAlleles":  100,
			},
			true,
			261,
		},
		{
			"Sample Name",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker",
				"Allele 1", "Allele 2", "Height 1", "Height 2"},
			[]string{},
			2, // n_the allele
			index{
				"Sample Name":  1,
				"Marker":       3,
				"Allele 1":     4,
				"AreaOffset":   -1,
				"HeightOffset": 2,
				"SizeOffset":   -1,
				"NoOfAlleles":  2,
			},
			false,
			5,
		},
		{
			"Sample Name",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker",
				"Allele 1", "Allele 2"},
			[]string{"Run Name"},
			1, // n_the allele
			index{
				"Sample Name":  1,
				"Marker":       3,
				"Allele 1":     4,
				"Run Name":     2,
				"AreaOffset":   -1,
				"HeightOffset": -1,
				"SizeOffset":   -1,
				"NoOfAlleles":  2,
			},
			false,
			4,
		}, {
			"Sample Name",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker",
				"Allele 1", "Height 1", "Allele 2", "Height 2"},
			[]string{},
			2, // n_the allele
			index{
				"Sample Name":  1,
				"Marker":       3,
				"Allele 1":     4,
				"AreaOffset":   -1,
				"HeightOffset": 1,
				"SizeOffset":   -1,
				"NoOfAlleles":  2,
			},
			true,
			6,
		},
		{
			"Sample File",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker",
				"Allele 1", "Height 1"},
			[]string{"Run Name"},
			1, // n_the allele
			index{
				"Sample File":  0,
				"Marker":       3,
				"Allele 1":     4,
				"AreaOffset":   -1,
				"HeightOffset": 1,
				"SizeOffset":   -1,
				"NoOfAlleles":  1,
				"Run Name":     2,
			},
			true,
			4,
		},
		{
			"Sample Name",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker",
				"Allele 1", "Allele 2", "Size 1", "Size 2", "Height 1",
				"Height 2"},
			[]string{},
			2, // n_the allele
			index{
				"Sample Name":  1,
				"Marker":       3,
				"Allele 1":     4,
				"AreaOffset":   -1,
				"HeightOffset": 4,
				"SizeOffset":   2,
				"NoOfAlleles":  2,
			},
			false,
			5,
		},
		{
			"Sample Name",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker",
				"Allele 1", "Area 1", "Height 1", "Allele 2", "Area 2",
				"Height 2"},
			[]string{},
			2, // n_the allele
			index{
				"Sample Name":  1,
				"Marker":       3,
				"Allele 1":     4,
				"AreaOffset":   1,
				"HeightOffset": 2,
				"SizeOffset":   -1,
				"NoOfAlleles":  2,
			},
			true,
			7,
		},
		{
			"Sample File",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker", "Dye",
				"Allele 1", "Allele 2", "Allele 3", "Allele 4", "Allele 5",
				"Allele 6", "Allele 7", "Allele 8", "Allele 9", "Allele 10",
				"Size 1", "Size 2", "Size 3", "Size 4", "Size 5", "Size 6",
				"Size 7", "Size 8", "Size 9", "Size 10", "Height 1", "Height 2",
				"Height 3", "Height 4", "Height 5", "Height 6", "Height 7",
				"Height 8", "Height 9", "Height 10"},
			[]string{"Dye", "Run Name"},
			10, // n_the allele
			index{
				"Sample File":  0,
				"Marker":       3,
				"Allele 1":     5,
				"AreaOffset":   -1,
				"HeightOffset": 20,
				"SizeOffset":   10,
				"NoOfAlleles":  10,
				"Dye":          4,
				"Run Name":     2,
			},
			false,
			14,
		},
		{
			"Sample Name",
			[]string{"Sample File", "Sample Name", "Run Name", "Marker", "Dye",
				"Allele 1", "Area 1", "Height 1", "Allele 2", "Area 2",
				"Height 2", "Allele 3", "Area 3", "Height 3", "Allele 4",
				"Area 4", "Height 4", "Allele 5", "Area 5", "Height 5",
				"Allele 6", "Area 6", "Height 6", "Allele 7", "Area 7",
				"Height 7", "Allele 8", "Area 8", "Height 8"},
			[]string{"Dye"},
			8, // n_the allele
			index{
				"Sample Name":  1,
				"Marker":       3,
				"Allele 1":     5,
				"AreaOffset":   1,
				"HeightOffset": 2,
				"SizeOffset":   -1,
				"NoOfAlleles":  8,
				"Dye":          4,
			},
			true,
			26,
		},
	}

	for i, tc := range tests {

		resIndex, err := indexHeader(tc.sID, tc.header, tc.info)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(tc.wantIndex, resIndex) {
			t.Fatalf("test %d (index): expected: %v, got: %v", i+1,
				tc.wantIndex, resIndex)
		}

		resCollated := isCollated(resIndex)
		if tc.wantCollated != resCollated {
			t.Fatalf("test %d (collated): expected: %v, got: %v", i+1,
				tc.wantCollated, resCollated)
		}

		resAllelePos := allelePos(tc.n, resIndex)
		if tc.wantAllelePos != resAllelePos {
			t.Fatalf("test %d (allelePos): expected: %v, got: %v", i+1,
				tc.wantAllelePos, resAllelePos)
		}
	}
}

// =============================================================================
func Test_parseLocus_parseAllele(t *testing.T) {

	type test struct {
		n          int      // n_th Allele, we want the column number
		line       []string // csv line to parse
		wantAllele Allele
		wantLocus  Locus
	}

	// the correctness of the indexHeader() function with these arguments was
	// tested above.
	idx, err := indexHeader("Sample File", provedITheader, []string{"Dye"})
	if err != nil {
		t.Error(err)
	}

	// provedIT data was downloaded from
	// https://lftdi.camden.rutgers.edu/provedit/
	tests := []test{
		{
			3, // third allele
			[]string{"A02_RD14-0003-40_41-1;4-M3S30-0.075IP-Q4.0_001.5sec.fsa",
				"vWA", "B",
				"OL", "114", "3",
				"7", "117.98", "2",
				"OL", "123.54", "4",
				"12", "137.64", "5",
				"13", "141.91", "40",
				"14", "146.24", "40",
				"15", "150.33", "21",
				"OL", "159.16", "1",
				"OL", "168.78", "3",
				"OL", "176.84", "3",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
				"", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			Allele{ID: -999, Area: 0, Height: 4, Size: 123.54},
			// Note: alleles are sorted by ID by default!
			Locus{ID: "VWA", Alleles: []Allele{
				{ID: -999, Area: 0, Height: 3, Size: 114},
				{ID: -999, Area: 0, Height: 4, Size: 123.54},
				{ID: -999, Area: 0, Height: 1, Size: 159.16},
				{ID: -999, Area: 0, Height: 3, Size: 168.78},
				{ID: -999, Area: 0, Height: 3, Size: 176.84},
				{ID: 7, Area: 0, Height: 2, Size: 117.98},
				{ID: 12, Area: 0, Height: 5, Size: 137.64},
				{ID: 13, Area: 0, Height: 40, Size: 141.91},
				{ID: 14, Area: 0, Height: 40, Size: 146.24},
				{ID: 15, Area: 0, Height: 21, Size: 150.33}}},
		},
	}

	for i, tc := range tests {

		resAllele, err := parseAllele(tc.n, tc.line, idx)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(tc.wantAllele, resAllele) {
			t.Fatalf("test %d (allele): expected: %v, got: %v", i+1,
				tc.wantAllele, resAllele)
		}

		resLocus, err := parseLocus(tc.line, idx)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(tc.wantLocus, resLocus) {
			t.Fatalf("test %d (locus): expected: %v, got: %v", i+1,
				tc.wantLocus, resLocus)
		}

	}
}
