// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func Test_NewLocus(t *testing.T) {

	type test struct {
		inID string
		want Locus
	}

	tests := []test{
		{"FGA", Locus{ID: "FGA"}},
		{"", Locus{}},
	}

	for i, tc := range tests {
		s := NewLocus(tc.inID)
		if !reflect.DeepEqual(s, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, s)
		}
	}
}

// =============================================================================
func TestLocus_AddAllele(t *testing.T) {

	type test struct {
		inLocus  Locus
		inAllele Allele
		want     Locus
	}

	tests := []test{
		{ // 1
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			Allele{ID: 21},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}, {ID: 21}}},
		},
		{ // 2
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			Allele{},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
		},
		{ // 3
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			Allele{ID: 9.3},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
		},
		{ // 4
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3, Height: 100}, {ID: -999, Height: 100}}},
			Allele{ID: -999, Height: 734},
			Locus{ID: "FGA", Alleles: []Allele{{ID: -999, Height: 100}, {ID: -999, Height: 734}, {ID: 9.3, Height: 100}}},
		},
	}

	for i, tc := range tests {
		tc.inLocus.AddAllele(tc.inAllele)
		if !reflect.DeepEqual(tc.inLocus, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, tc.inLocus)
		}
	}
}

// =============================================================================
func TestLocus_RemoveAllele(t *testing.T) {

	type test struct {
		inLocus  Locus
		inAllele Allele
		want     Locus
	}

	tests := []test{
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			Allele{ID: 12},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}}},
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			Allele{},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			Allele{ID: 9.3},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 12}}},
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3, Height: 100}, {ID: -999, Height: 100}}},
			Allele{ID: -999, Height: 734},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3, Height: 100}}},
		},
	}

	for i, tc := range tests {
		tc.inLocus.RemoveAllele(tc.inAllele)
		if !reflect.DeepEqual(tc.inLocus, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, tc.inLocus)
		}
	}
}

// =============================================================================
func TestLocus_Allele(t *testing.T) {

	type test struct {
		inLocus Locus
		inID    float64
		want    Allele
	}

	tests := []test{
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9, Height: 928}, {ID: 12, Height: 212}}},
			9,
			Allele{ID: 9, Height: 928},
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9, Height: 928}, {ID: 12, Height: 212}}},
			21,
			Allele{},
		},
	}

	for i, tc := range tests {
		allele := tc.inLocus.Allele(tc.inID)
		if !reflect.DeepEqual(allele, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, allele)
		}
	}
}

// =============================================================================
func TestLocus_HasAllele(t *testing.T) {

	type test struct {
		inLocus Locus
		inID    float64
		want    bool
	}

	tests := []test{
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9, Height: 928}, {ID: 12, Height: 212}}},
			9,
			true,
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9, Height: 928}, {ID: 12, Height: 212}}},
			21,
			false,
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9, Height: 928}, {ID: 12, Height: 212}}},
			0,
			false,
		},
	}

	for i, tc := range tests {
		res := tc.inLocus.HasAllele(tc.inID)
		if res != tc.want {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestLocus_linkage(t *testing.T) {

	type test struct {
		inLocus Locus
		want    LocusLinkage
	}

	tests := []test{
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			AUTOSOMAL,
		},
		{
			Locus{ID: "DYS635", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			YLINKED,
		},
		{
			Locus{ID: "DXS10101", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			XLINKED,
		},
		{
			Locus{ID: "Noname", Alleles: []Allele{{ID: 9.3}, {ID: 12}}},
			AUTOSOMAL,
		},
	}

	for i, tc := range tests {
		res := tc.inLocus.Linkage()
		if !reflect.DeepEqual(res, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func TestLocus_SortByID(t *testing.T) {

	type test struct {
		inLocus Locus
		want    Locus
	}

	tests := []test{
		{ // 1
			Locus{ID: "FGA", Alleles: []Allele{{ID: 19, Height: 928}, {ID: 12, Height: 212}}},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 12, Height: 212}, {ID: 19, Height: 928}}},
		},
		{ // 2
			Locus{ID: "FGA", Alleles: []Allele{{ID: 19, Height: 928}, {ID: 7, Height: 2742}, {ID: 12, Height: 212}}},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 7, Height: 2742}, {ID: 12, Height: 212}, {ID: 19, Height: 928}}},
		},
	}

	for i, tc := range tests {
		tc.inLocus.SortByID()
		if !reflect.DeepEqual(tc.inLocus, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, tc.inLocus)
		}
	}
}

// =============================================================================
func TestLocus_SortByHeight(t *testing.T) {

	type test struct {
		inLocus Locus
		want    Locus
	}

	tests := []test{
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 19, Height: 928}, {ID: 12, Height: 212}}},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 19, Height: 928}, {ID: 12, Height: 212}}},
		},
		{
			Locus{ID: "FGA", Alleles: []Allele{{ID: 19, Height: 928}, {ID: 7, Height: 2742}, {ID: 12, Height: 212}}},
			Locus{ID: "FGA", Alleles: []Allele{{ID: 7, Height: 2742}, {ID: 19, Height: 928}, {ID: 12, Height: 212}}},
		},
	}

	for i, tc := range tests {
		res := tc.inLocus.SortByHeight()
		if !reflect.DeepEqual(res, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
