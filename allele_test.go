// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"reflect"
	"testing"
)

// =============================================================================
func Test_NewAllele(t *testing.T) {

	type test struct {
		inID float64
		want Allele
	}

	tests := []test{
		{9.3, Allele{ID: 9.3}},
		{29, Allele{ID: 29}},
	}

	for i, tc := range tests {
		a := NewAllele(tc.inID)
		if !reflect.DeepEqual(a, tc.want) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, a)
		}
	}
}

// =============================================================================
func Test_A2Float(t *testing.T) {

	type test struct {
		inID string
		want float64
	}

	tests := []test{
		{"X", -2},
		{"x", -2},
		{"Y", -1},
		{"y", -1},
		{"", 0},
		{"23.2", 23.2},
		{"OL", -999},
	}

	for i, tc := range tests {
		res := A2Float(tc.inID)
		if tc.want != res {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}

// =============================================================================
func Test_A2String(t *testing.T) {

	type test struct {
		inID float64
		want string
	}

	tests := []test{
		{-2, "X"},
		{-1, "Y"},
		{0, ""},
		{23.2, "23.2"},
		{-999, "NaN"},
	}

	for i, tc := range tests {
		res := A2String(tc.inID)
		if tc.want != res {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, res)
		}
	}
}
