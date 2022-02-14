// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ExportCSV writes a sample s as CSV to a file of name f, separated by sep.
// If collate is true, the alleles of a marker will be sorted as Allele 1,
// Height 1, ..., Allele 2, Height 2, ... etc. If collate is false, they will be
// sorted as Allele 1, Allele 2, ..., Height 1, Height 2, ... etc.
func (s Sample) ExportCSV(f string, sep rune, collate bool) error {

	d, err := buildCSV(s, collate)
	if err != nil {
		return fmt.Errorf("buildCSV fails: %v", err)
	}

	if err := write2CSV(d, f, sep); err != nil {
		return fmt.Errorf("write2CSV fails: %v", err)
	}

	return nil
}

// buildCSV builds a slice of rows from sample s to be written to a CSV file. It
// can write s in collated (collated = true) or uncollated from (collated =
// false).
func buildCSV(s Sample, collated bool) ([][]string, error) {

	var header []string
	if collated {
		header = collatedHeader(s)
	} else {
		header = uncollatedHeader(s)
	}

	csvData := [][]string{header}

	for _, l := range s.Loci {
		var locus []string
		if collated {
			locus = collatedRow(l, s.MaxAlleles(), s.alleleFields())
		} else {
			locus = uncollatedRow(l, s.MaxAlleles(), s.alleleFields())
		}
		row := append([]string{s.ID, l.ID}, locus...)

		csvData = append(csvData, row)
	}

	return csvData, nil
}

// collatedHeader returns a collated header, i.e. the first line, of the CSV
//file (i.e. in the form of A1, S1, H1, A2, S2, H2).
func collatedHeader(s Sample) []string {
	a := []string{"Sample Name", "Marker"}

	for i := 0; i < s.MaxAlleles(); i++ {
		for _, f := range s.alleleFields() {
			// i+1 to start Allele 1 and not Allele 0
			a = append(a, f+" "+strconv.Itoa(i+1))
		}
	}
	return a
}

// uncollatedHeader returns a uncollated header, i.e. the first line, of the CSV
// file (i.e. in the form of A1, A2, A3, H1, H2, H3).
func uncollatedHeader(s Sample) []string {
	a := []string{"Sample Name", "Marker"}

	for _, f := range s.alleleFields() {
		for i := 0; i < s.MaxAlleles(); i++ {
			// i+1 to start Allele 1 and not Allele 0
			a = append(a, f+" "+strconv.Itoa(i+1))
		}
	}

	return a
}

// (e.g. A1, S1, H1, A2, S2, H2)
func collatedRow(l Locus, max int, fields alleleFields) []string {

	var row []string // 1 line per locus/marker
	for _, a := range l.Alleles {
		for _, f := range fields {
			switch f {
			case "Allele":
				row = append(row, A2String(a.ID))
			case "Height":
				row = append(row, A2String(a.Height))
			case "Area":
				row = append(row, A2String(a.Area))
			case "Size":
				row = append(row, A2String(a.Size))
			}

		}
	}

	var emptyFields []string
	for i := 0; i < len(fields); i++ {
		emptyFields = append(emptyFields, "")
	}

	// fill
	for i := 0; max-(len(l.Alleles)+i) > 0; i++ {
		row = append(row, emptyFields...)
	}

	return row
}

// (e.g. A1, A2, A3, H1, H2, H3)
func uncollatedRow(l Locus, max int, fields alleleFields) []string {

	var row []string
	for _, f := range fields {
		for _, a := range l.Alleles {
			switch f {
			case "Allele":
				row = append(row, A2String(a.ID))
			case "Height":
				row = append(row, A2String(a.Height))
			case "Size":
				row = append(row, A2String(a.Size))
			case "Area":
				row = append(row, A2String(a.Area))
			}
		}

		// fill
		for i := 0; max-(len(l.Alleles)+i) > 0; i++ {
			row = append(row, "")
		}
	}

	return row
}

// alleleFields
type alleleFields []string

// alleleFields returns the allele field names present in a Sample object s.
func (s Sample) alleleFields() alleleFields {

	var f alleleFields
	for _, l := range s.Loci {
		for _, a := range l.Alleles {
			if a.ID != 0 { // can be -2 or -1 in AMEL
				f = append(f, "Allele")
			}
			if a.Height > 0 {
				f = append(f, "Height")
			}
			if a.Area > 0 {
				f = append(f, "Area")
			}
			if a.Size > 0 {
				f = append(f, "Size")
			}

			return f
		}
	}

	return nil
}

// write2CSV writes the 2d data d to a file with name f.
func write2CSV(d [][]string, f string, sep rune) error {

	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
		// TODO: handle error
	}(file)

	w := csv.NewWriter(file)
	w.Comma = sep
	if err := w.WriteAll(d); err != nil {
		return err
	}

	return nil
}
