---
title: Parse
icon: cpu
order: -20
---

`go-cff` provides two ways to parse the data:
* `cff.Parse(content string)` for parsing a string
* `cff.ParseFile(filename string)` for parsing a file

The returned value is a [`cff.CFF`](https://github.com/alexander-lindner/go-cff/blob/main/format.go) object.
This object contains all the data from the CFF file, not available fields are empty.

## Example

```go
parsedCFF, err := cff.ParseFile("./CITATION.cff")
if err != nil {
    fmt.Println("Error: ", err)
    return
}
fmt.Println("CFF File Version: ", parsedCFF.CffVersion)
```

### with arrays

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
