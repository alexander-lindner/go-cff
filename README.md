# go-cff 

[![DOI](https://zenodo.org/badge/487247273.svg)](https://zenodo.org/badge/latestdoi/487247273)


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

