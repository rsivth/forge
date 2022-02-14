// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"strconv"
)

// Allele defines the fdo Allele struct.
type Allele struct {
	ID     float64 // name of the allele (e.g. 9.3)
	Area   float64 // peak area size
	Height float64 // signal strength, i.e. peak height in rfu
	Size   float64 // fragment length
}

// NewAllele returns an Allele object wit ID id.
func NewAllele(id float64) Allele {
	return Allele{
		ID: id,
	}
}

// A2Float converts the allele ID from string to float.
func A2Float(a string) float64 {
	allele, err := strconv.ParseFloat(a, 64)
	if err != nil {
		switch a {
		case "":
			allele = 0
		case "X", "x":
			allele = -2.0
		case "Y", "y":
			allele = -1.0
		default: // nonsense id (e.g. 'OL')
			allele = -999
		}
	}

	return allele
}

// A2String converts the allele ID from float64 to string.
func A2String(a float64) string {
	switch a {
	case 0:
		return ""
	case -1:
		return "Y"
	case -2:
		return "X"
	case -999:
		return "NaN"
	default:
		return strconv.FormatFloat(a, 'f', -1, 64)
	}
}
