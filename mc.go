// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

// MajorComponent returns the major component in sample s, if present. It uses
// four parameters for the decision. It returns only loci that contain data.
func (s Sample) MajorComponent(heteroImbalance, majorCompRatio, minHomozygous, weakSignal float64) Sample {

	mcSamp := NewSample("MC::", s.Source)

	for _, l := range s.Loci {
		mcLoc := l.MajorComponent(heteroImbalance, majorCompRatio, minHomozygous, weakSignal)
		if len(mcLoc.Alleles) > 0 {
			mcSamp.AddLocus(mcLoc)
		}
	}

	mcSamp.UnknownKit()
	return mcSamp
}

// HasMajorComponent returns true if sample s contains a major component according to the
// parameters provided in conf.
func (s Sample) HasMajorComponent(heteroImbalance, majorCompRatio, minHomozygous, weakSignal float64) bool {
	return len(s.MajorComponent(heteroImbalance, majorCompRatio, minHomozygous, weakSignal).Loci) > 0
}

// MajorComponent attempts to infer a major component at locus l based on the
// given parameters. If no major component is found, it returns a Locus object
// with the same ID as l but without any alleles.
func (l Locus) MajorComponent(heteroImbalance, majorCompRatio, minHomozygous, weakSignal float64) Locus {

	if l.ID == "IQCS" || l.ID == "IQCL" { // TODO: maybe get list of alleles from file kits.go?
		return NewLocus(l.ID)
	}

	if len(l.Alleles) < 2 {
		// The single allele is large enough for homozygous MC.
		// Note: GO guarantees that the first condition is
		// evaluated before the second one.
		if len(l.Alleles) == 1 &&
			l.Alleles[0].Height >= minHomozygous {
			return l
		}
		return NewLocus(l.ID)
	}

	// now we have at least 2 alleles at this locus.
	// sort the alleles of the locus by height and get the first 2 largest.
	sortLoc := l.SortByHeight()
	a1 := sortLoc.Alleles[0]
	a2 := sortLoc.Alleles[1]

	// initialize locus for major component
	mcLoc := NewLocus(l.ID)

	// First allele is distinct and strong enough for homozygous MC.
	if a1.Height/a2.Height > majorCompRatio &&
		a1.Height > minHomozygous {
		mcLoc.AddAllele(a1)
		return mcLoc
	}

	// The two alleles are both strong enough and similar enough for MC.
	if a2.Height/a1.Height > heteroImbalance &&
		a2.Height > weakSignal {
		if len(l.Alleles) == 2 {
			return l
		}
		// now we have at least 3 alleles
		a3 := sortLoc.Alleles[2] // third largest allele
		// are the first two distinct enough?
		if a2.Height/a3.Height > majorCompRatio {
			mcLoc.AddAllele(a1)
			mcLoc.AddAllele(a2)
			return mcLoc
		}
	}

	return NewLocus(l.ID)
}
