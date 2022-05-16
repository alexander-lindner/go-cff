# go-cff 

[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.6551082.svg)](https://doi.org/10.5281/zenodo.6551082)
[![Audit](https://github.com/alexander-lindner/go-cff/actions/workflows/tests.yml/badge.svg)](https://github.com/alexander-lindner/go-cff/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/alexander-lindner/go-cff/badge.svg?branch=main)](https://coveralls.io/github/alexander-lindner/go-cff?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexander-lindner/go-cff)](https://goreportcard.com/report/github.com/alexander-lindner/go-cff)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexander-lindner/go-cff.svg)](https://pkg.go.dev/github.com/alexander-lindner/go-cff)
![License](https://img.shields.io/github/license/alexander-lindner/go-cff?style=flat-square)
![GitHub top language](https://img.shields.io/github/languages/top/alexander-lindner/go-cff)


`go-cff` is an easy-to-use library for working with the [Citation File Format (CFF)](https://citation-file-format.github.io/) format in GO.


```go
parsedCFF, err := cff.ParseFile("./CITATION.cff")
if err != nil {
    fmt.Println("Error: ", err)
    return
}
for _, reference := range parsedCFF.References {
    fmt.Println(reference.Title)
}
```



See [Documentation](https://alexander-lindner.github.io/go-cff/) for more information.

## ToDo's

* Update documentation
* Directory structure

