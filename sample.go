// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

// Sample defines the fdo Sample struct. It holds the forensic genetic
// information for a single sample.
type Sample struct {
	// unique identifier of the sample (e.g. Stain-01)
	ID string
	// Info provides any additional information to the sample.
	Info Info
	// Kit describes PCR kit (will be inferred from the data)
	Kit Kit
	// Source from where the sample was obtained, e.g. file name.
	Source string
	// Loci constitutes the number of loci constituting for this stain.
	Loci []Locus
}

// Info may hold additional information on the sample as read from
// the Genemapper file.
type Info map[string]string

// NewSample generates a Sample object.
func NewSample(id, source string) Sample {
	return Sample{
		ID:     id,
		Info:   make(map[string]string),
		Source: source,
	}
}

// AddLocus adds a Locus l to Sample s. It will only add the locus if the l
// has an ID and there is not already a locus of the same id present. An empty
// (of alleles) locus will still be sampled so we get the locus for kit
// inference etc.
func (s *Sample) AddLocus(l Locus) {
	if l.ID != "" && !s.HasLocus(l.ID) { // loci must be unique
		s.Loci = append(s.Loci, l)
	}
}

// Locus returns the locus of name id. If no such locus is found it returns an
// empty struct.
func (s Sample) Locus(id string) Locus {
	for _, l := range s.Loci {
		if id == l.ID {
			return l
		}
	}

	return Locus{}
}

// HasLocus returns true if Sample s contains a Locus id with alleles.
// Otherwise, it returns false.
func (s Sample) HasLocus(id string) bool {
	return len(s.Locus(id).Alleles) > 0
}

// MaxAlleles returns the maximum number of alleles amongst all loci of a
// sample s.
func (s Sample) MaxAlleles() int {
	var max int
	for _, l := range s.Loci {
		if len(l.Alleles) > max {
			max = len(l.Alleles)
		}
	}
	return max
}

// TotalNumberOfAlleles returns the total number of alleles in sample s.
func (s Sample) TotalNumberOfAlleles() int {
	var r int
	for _, l := range s.Loci {
		r += len(l.Alleles)
	}
	return r
}

// MinContributor returns the minimum number of people that could have
// contributed to sample s.
func (s Sample) MinContributor() int {

	if s.MaxAlleles()%2 == 0 {
		return s.MaxAlleles() / 2
	}

	// integer division rounds down, so we need to add 1 to
	// even max alleles numbers.
	return (s.MaxAlleles() + 1) / 2
}

// ChangeID
func (s *Sample) ChangeID(id string) {
	s.ID = id
}
