# forge - FORensic GEnetics library

![](https://img.shields.io/badge/build-passing-success)
![](https://img.shields.io/badge/tests-passing-success)
![](https://img.shields.io/badge/go_vet-passing-success)
![](https://img.shields.io/badge/golint-passing-success)
![](https://img.shields.io/badge/test%20coverage-59%25-yellow)

This package reads and manipulates forensic genetics data such
as [short tandem repeats](https://en.wikipedia.org/wiki/STR_analysis#Forensic_uses) (STRs). forge is written
in [Go](https://golang.org/) and employs
a [test-driven development](https://en.wikipedia.org/wiki/Test-driven_development)
approach to validate the correctness of its functions. For some tests the PROVEDit dataset was used (downloaded
from [https://lftdi.camden.rutgers.edu/provedit/files/](https://lftdi.camden.rutgers.edu/provedit/files/)). The dataset
was published by [Alfonse et al.](https://www.fsigenetics.com/article/S1872-4973(17)30214-4/fulltext) (2017).

In particular, this package provides functions that:

- import STR samples from a [Genemapper](https://www.thermofisher.com/order/catalog/product/4475073) CSV file
- import lab reference profiles as exported from Genemapper
- import allele frequency information from the [STRider.online](https://www.STRider.online) XML file
- match reference profiles with stain samples
- infer profiles of unknown persons from stain samples
- export STR samples as Genemapper CSV files
- perform basic forensic statistics such as CPI and RMNE

