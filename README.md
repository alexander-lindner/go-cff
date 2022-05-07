# go-cff
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

* ~~definitions.country~~
* ~~definitions.email~~
* ~~definitions.fax~~
* ~~definitions.license-enum~~
* ~~definitions.orcid~~
* ~~definitions.reference.isbn~~
* ~~definitions.reference.issn~~
* ~~definitions.reference.languages~~
* ~~definitions.reference.pmcid~~
* ~~definitions.reference.status~~
* ~~definitions.reference.type~~
* ~~definitions.swh-identifier~~
* ~~definitions.url~~
* ~~definitions.version~~
