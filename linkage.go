// Copyright (c) 2017-2022 Roland Schultheiß. All rights reserved.
// License information can be found in the LICENSE file.

package forge

type LocusLinkage int

const (
	AUTOSOMAL LocusLinkage = iota
	YLINKED
	XLINKED
	ALLLINKAGE
)

// String returns the locus linkage as string.
func (l LocusLinkage) String() string {
	switch l {
	case AUTOSOMAL:
		return "autosomal"
	case YLINKED:
		return "Y-linked"
	case XLINKED:
		return "X-linked"
	default: // ALLLINKAGE
		return "all_linkage"
	}
}

// Linkage returns linkage of locus l.
func (l Locus) Linkage() LocusLinkage {
	switch {
	case LocusCoordinates(l.ID).Chr == -1:
		return YLINKED
	case LocusCoordinates(l.ID).Chr == -2:
		return XLINKED
	default:
		return AUTOSOMAL
	}
}

// GenomicCoordinates holds the genomic coordinates of a locus.
type GenomicCoordinates struct {
	Chr   int // Chromosome number (X = -2, Y = -1)
	Start int
	End   int
}

// LocusCoordinates returns the genomic coordinates of locus with id id.
// Genomic coordinates info: https://strbase.nist.gov//chrom.htm
func LocusCoordinates(id string) GenomicCoordinates {
	switch id {
	case "D1S1656":
		return GenomicCoordinates{Chr: 1, Start: 228972000}
	case "TPOX":
		return GenomicCoordinates{Chr: 2, Start: 1472000}
	case "D2S441":
		return GenomicCoordinates{Chr: 2, Start: 68213613}
	case "D2S1338":
		return GenomicCoordinates{Chr: 2, Start: 218705000}
	case "D3S1358":
		return GenomicCoordinates{Chr: 3, Start: 45557000}
	case "FGA":
		return GenomicCoordinates{Chr: 4, Start: 155866000}
	case "CSF1PO":
		return GenomicCoordinates{Chr: 5, Start: 149436000}
	case "D5S818":
		return GenomicCoordinates{Chr: 5, Start: 123139000}
	case "SE33":
		return GenomicCoordinates{Chr: 6, Start: 89043000}
	case "D7S820":
		return GenomicCoordinates{Chr: 7, Start: 83433000}
	case "D8S1179":
		return GenomicCoordinates{Chr: 8, Start: 125976000}
	case "D10S1248":
		return GenomicCoordinates{Chr: 10, Start: 130566908}
	case "TH01":
		return GenomicCoordinates{Chr: 11, Start: 2149000}
	case "VWA":
		return GenomicCoordinates{Chr: 12, Start: 5963000}
	case "D12S391":
		return GenomicCoordinates{Chr: 12, Start: 12341000}
	case "D13S317":
		return GenomicCoordinates{Chr: 13, Start: 81620000}
	case "D14S1434":
		return GenomicCoordinates{Chr: 14, Start: 93298432}
	case "Penta E":
		return GenomicCoordinates{Chr: 15, Start: 95175000}
	case "D16S539":
		return GenomicCoordinates{Chr: 16, Start: 84944000}
	case "D18S51":
		return GenomicCoordinates{Chr: 18, Start: 59100000}
	case "D19S433":
		return GenomicCoordinates{Chr: 19, Start: 35109000}
	case "Penta D":
		return GenomicCoordinates{Chr: 21, Start: 43880000}
	case "D21S11":
		return GenomicCoordinates{Chr: 21, Start: 19476000}
	case "D22S1045":
		return GenomicCoordinates{Chr: 22, Start: 35779368}
	case "DYS576":
		return GenomicCoordinates{Chr: -1}
	case "DYS389I":
		return GenomicCoordinates{Chr: -1}
	case "DYS635":
		return GenomicCoordinates{Chr: -1}
	case "DYS389II":
		return GenomicCoordinates{Chr: -1}
	case "DYS627":
		return GenomicCoordinates{Chr: -1}
	case "DYS460":
		return GenomicCoordinates{Chr: -1}
	case "DYS458":
		return GenomicCoordinates{Chr: -1}
	case "DYS19":
		return GenomicCoordinates{Chr: -1}
	case "YGATAH4":
		return GenomicCoordinates{Chr: -1}
	case "DYS448":
		return GenomicCoordinates{Chr: -1}
	case "DYS391":
		return GenomicCoordinates{Chr: -1}
	case "DYS456":
		return GenomicCoordinates{Chr: -1}
	case "DYS390":
		return GenomicCoordinates{Chr: -1}
	case "DYS438":
		return GenomicCoordinates{Chr: -1}
	case "DYS392":
		return GenomicCoordinates{Chr: -1}
	case "DYS518":
		return GenomicCoordinates{Chr: -1}
	case "DYS570":
		return GenomicCoordinates{Chr: -1}
	case "DYS437":
		return GenomicCoordinates{Chr: -1}
	case "DYS385 a/b": // TODO: a/b will be upper letter, right? see NewLocus()
		return GenomicCoordinates{Chr: -1}
	case "DYS449":
		return GenomicCoordinates{Chr: -1}
	case "DYS393":
		return GenomicCoordinates{Chr: -1}
	case "DYS439":
		return GenomicCoordinates{Chr: -1}
	case "DYS481":
		return GenomicCoordinates{Chr: -1}
	case "DYF387S1":
		return GenomicCoordinates{Chr: -1}
	case "DYS533":
		return GenomicCoordinates{Chr: -1}
	case "YINDEL": //TODO: check whether correct name
		return GenomicCoordinates{Chr: -1}
	case "DXS10101":
		return GenomicCoordinates{Chr: -2}
	default:
		return GenomicCoordinates{}
	}
}
