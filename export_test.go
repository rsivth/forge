// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func Test_header(t *testing.T) {

	type test struct {
		inSample       Sample
		wantCollated   []string
		wantUnCollated []string
	}

	tests := []test{
		{
			Sample{
				ID: "TestID",
				Loci: []Locus{
					{ID: "VWA", Alleles: []Allele{
						{ID: 17, Height: 2039, Size: 111.2},
						{ID: 21, Height: 1993, Size: 122.1}}},
					{ID: "FGA", Alleles: []Allele{
						{ID: 20.1, Height: 322, Size: 343},
						{ID: 31, Height: 8923, Size: 422.3},
						{ID: -999, Height: 3432, Size: 723}}},
				},
			},
			[]string{"Sample Name", "Marker",
				"Allele 1", "Height 1", "Size 1",
				"Allele 2", "Height 2", "Size 2",
				"Allele 3", "Height 3", "Size 3"},
			[]string{"Sample Name", "Marker",
				"Allele 1", "Allele 2", "Allele 3",
				"Height 1", "Height 2", "Height 3",
				"Size 1", "Size 2", "Size 3"},
		},
		{
			Sample{
				ID: "TestID",
				Loci: []Locus{
					{ID: "VWA", Alleles: []Allele{{ID: 17}, {ID: 21}}},
					{ID: "FGA", Alleles: []Allele{{ID: 20.1}, {ID: 31}}}},
			},
			[]string{"Sample Name", "Marker", "Allele 1", "Allele 2"},
			[]string{"Sample Name", "Marker", "Allele 1", "Allele 2"},
		},
	}

	for i, tc := range tests {
		resCollated := collatedHeader(tc.inSample)
		if !reflect.DeepEqual(tc.wantCollated, resCollated) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1,
				tc.wantCollated, resCollated)
		}

		resUnCollated := uncollatedHeader(tc.inSample)
		if !reflect.DeepEqual(tc.wantUnCollated, resUnCollated) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1,
				tc.wantUnCollated, resUnCollated)
		}
	}
}

// =============================================================================
func Test_alleleFields(t *testing.T) {

	type test struct {
		inSample Sample
		want     alleleFields
	}

	tests := []test{
		{
			Sample{
				ID: "TestID",
				Loci: []Locus{
					{ID: "SE33", Alleles: []Allele{}},
					{ID: "VWA", Alleles: []Allele{
						{ID: 17, Height: 2039, Size: 111.2},
						{ID: 21, Height: 1993, Size: 122.1}}},
					{ID: "FGA", Alleles: []Allele{
						{ID: 20.1, Height: 322, Size: 343},
						{ID: 31, Height: 8923, Size: 422.3}}},
				},
			},
			alleleFields{"Allele", "Height", "Size"},
		},
		{
			Sample{
				ID: "TestID",
				Loci: []Locus{
					{ID: "AMEL", Alleles: []Allele{{ID: -2}, {ID: -1}}},
					{ID: "FGA", Alleles: []Allele{{ID: 20.1}, {ID: 31}}}},
			},
			alleleFields{"Allele"},
		},
		{
			Sample{
				ID: "TestID",
				Loci: []Locus{
					{ID: "VWA", Alleles: []Allele{
						{ID: 17, Height: 232, Area: 3432.7, Size: 567.1}}},
					{ID: "FGA", Alleles: []Allele{
						{ID: 9.3, Height: 675, Area: 432.37, Size: 2324}}},
				},
			},
			alleleFields{"Allele", "Height", "Area", "Size"},
		},
	}

	for i, tc := range tests {
		res := tc.inSample.alleleFields()
		if !reflect.DeepEqual(tc.want, res) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func Test_collatedCSV_uncollatedCSV(t *testing.T) {

	type test struct {
		inSample       Sample
		wantCollated   [][]string
		wantUnCollated [][]string
	}

	tests := []test{
		{
			Sample{
				ID: "TestID",
				Loci: []Locus{
					{ID: "VWA", Alleles: []Allele{
						{ID: 17, Height: 2039, Size: 111.2},
						{ID: 27.1, Height: 122, Size: 326}}},
					{ID: "FGA", Alleles: []Allele{
						{ID: 20.1, Height: 322, Size: 343},
						{ID: 21, Height: 1993, Size: 382.1},
						{ID: -999, Height: 8923, Size: 433.3}}},
				},
			},
			[][]string{
				{"Sample Name", "Marker",
					"Allele 1", "Height 1", "Size 1",
					"Allele 2", "Height 2", "Size 2",
					"Allele 3", "Height 3", "Size 3",
				},
				{"TestID", "VWA",
					"17", "2039", "111.2",
					"27.1", "122", "326",
					"", "", "",
				},
				{"TestID", "FGA",
					"20.1", "322", "343",
					"21", "1993", "382.1",
					"NaN", "8923", "433.3",
				},
			},
			[][]string{
				{"Sample Name", "Marker",
					"Allele 1", "Allele 2", "Allele 3",
					"Height 1", "Height 2", "Height 3",
					"Size 1", "Size 2", "Size 3",
				},
				{"TestID", "VWA",
					"17", "27.1", "",
					"2039", "122", "",
					"111.2", "326", "",
				},
				{"TestID", "FGA",
					"20.1", "21", "NaN",
					"322", "1993", "8923",
					"343", "382.1", "433.3",
				},
			},
		},
	}

	for i, tc := range tests {
		resCollated, err := buildCSV(tc.inSample, true)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(tc.wantCollated, resCollated) {
			t.Fatalf("test %d (collated): expected: %v, got: %v", i+1,
				tc.wantCollated, resCollated)
		}

		resUnCollated, err := buildCSV(tc.inSample, false)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(tc.wantUnCollated, resUnCollated) {
			t.Fatalf("test %d (uncollated): expected: %v, got: %v", i+1,
				tc.wantUnCollated, resUnCollated)
		}
	}
}
