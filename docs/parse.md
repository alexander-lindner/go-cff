---
title: Parse
icon: cpu
order: -20
---

`go-cff` provides two ways to parse the data:
* [`cff.Parse(content string)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Parse) for parsing a string
* [`cff.ParseFile(filename string)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#ParseFile) for parsing a file

The returned value is a [`cff.CFF`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Cff) object.
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

## The `PersonEntity` type

The **CFF** file format allows in several places two types at the same time:
* [`Person`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Person)
* [`Entity`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Entity)

Go doesn't allow to have a variable with several types.
Thus, we created a small proxy struct to handle this situation: [`PersonEntity`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#PersonEntity).
```go
type PersonEntity struct {
	Person   Person
	Entity   Entity
	IsEntity bool
	IsPerson bool
}
```
As you see, it has the two fields for the `Person` and the `Entity` type.
Additionally, it also has two boolean fields to know if it is a `Person` or an `Entity`.
The leftover field contains an empty struct of the corresponding type.

### Example
```go
parser, _ := cff.ParseFile("./CITATION.cff")

for _, item := range parser.Authors {
    if item.IsPerson() {
        fmt.Println(item.Person.GivenNames)
    } else {
        fmt.Println(item.Entity.Address)
    }
}
```