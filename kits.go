// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Kit described a specific PCR kit.
type Kit struct {
	ID   string `json:"ID"`   // name of the kit, e.g. NGM-Detect
	STRs []STR  `json:"STRs"` // slice of STRs in the kit
}

// STR describes the name and dye of an STR for a specific kit.
type STR struct {
	ID  string `json:"ID"`  // e.g. VWA
	Dye string `json:"Dye"` // color of the dye
}

// TODO: fix the tests for this function

// InferKit infers the kit of Sample s based on the kits in folder dir.
func (s *Sample) InferKit(dir string) error {

	if len(s.Loci) < 3 || dir == "" {
		// not an error
		s.UnknownKit()
		return nil
	}

	kits, err := readKitFiles(dir)
	if err != nil {
		return fmt.Errorf("cannot infer sample kit: %v", err)
	}

	for _, kit := range kits {
		if kit.matchSample(*s) {
			s.Kit = kit
			return nil
		}
	}

	// if no kit matched the sample
	s.UnknownKit()
	return nil
}

// AssignKit assigns a kit k to sample s. It will override any previously
// assigned kit to s.
func (s *Sample) AssignKit(k Kit) {
	s.Kit = k
}

// readKitFiles reads json files with PCR-kit information from folder dir.
func readKitFiles(dir string) ([]Kit, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return []Kit{}, fmt.Errorf("cannot read kit folder %v: %v", dir, err)
	}

	kits := make(map[string]Kit)
	for _, f := range files {
		fName := dir + f.Name()
		kit, err := readKitFile(fName)
		if err != nil {
			return []Kit{}, fmt.Errorf("cannot read kit file: %v", err)
		}

		// avoid multiple entries for the same kit ID
		if _, ok := kits[kit.ID]; ok {
			return []Kit{}, fmt.Errorf("more than on file with kit ID %v | delete this entry", kit.ID)
		}

		kits[kit.ID] = kit
	}

	var kitSlice []Kit
	for _, k := range kits {
		kitSlice = append(kitSlice, k)
	}
	return kitSlice, nil
}

// readKitFile
func readKitFile(f string) (Kit, error) {

	kFile, err := os.Open(f)
	if err != nil {
		return Kit{}, fmt.Errorf("cannot open kit json file:%v", err)
	}
	defer func(kFile *os.File) {
		err := kFile.Close()
		if err != nil {
			// TODO: handle error
		}
	}(kFile)

	var kit Kit
	jP := json.NewDecoder(kFile)
	if err = jP.Decode(&kit); err != nil {
		return Kit{}, fmt.Errorf("cannot decode kit file:%v", err)
	}

	// avoid vWA/VWA issues like in parser.go
	var strs []STR
	for _, str := range kit.STRs {
		strs = append(strs, STR{strings.ToUpper(str.ID), str.Dye})
	}

	return Kit{
		ID:   kit.ID,
		STRs: strs,
	}, nil
}

// matchSample determines whether sample s is of kit k.
func (k Kit) matchSample(s Sample) bool {
	for i, locus := range s.Loci {
		// If the sample has more loci than the kit or the IDs
		// don't match this is not the kit we are looking for.
		if i >= len(k.STRs) || locus.ID != k.STRs[i].ID {
			return false
		}
	}

	return true
}

// UnknownKit returns a Kit object for an unknown kit. TODO: add to tests
func (s *Sample) UnknownKit() {
	var strs []STR
	for _, l := range s.Loci {
		strs = append(strs, STR{ID: l.ID})
	}

	s.Kit = Kit{
		ID:   "unknown Kit",
		STRs: strs,
	}
}

// IsOfUnknownKit determines whether the kit of sample s is known.
func (s Sample) IsOfUnknownKit() bool {
	return s.Kit.ID == "unknown Kit"
}

// HasSTR checks whether a PCR kit k has an STR of name str.
func (k Kit) HasSTR(str string) bool {

	for _, s := range k.STRs {
		if s.ID == str {
			return true
		}
	}

	return false
}
