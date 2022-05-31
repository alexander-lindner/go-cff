---
title: Write
icon: pencil
order: -30
---

`go-cff` provides two ways to serialize a [`cff.CFF`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Cff) object as the CFF format:
* [`cff.Save(content Cff)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Save) for getting a string back
* [`cff.SaveFile(file string, content Cff)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#SaveFile) for saving as a file

 [`cff.CFF`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#Cff) is a struct,
 so you can easly create a new object from that struct or extend an already read in object.
Empty and not set fields are skipped during serialisation.
## Helper

For providing a better experience, we use special data structures when parsing a file, for 
example `url.URL`.
However, when creating a field with such a data structure, it is often really messy.
Hence, you can use the following helper functions as alegant shortcuts:
* [`cff.MakeUrl(string)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#MakeUrl)
* [`cff.MakeDoi(string)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#MakeDoi)
* [`cff.MakeLicensesByString(...string)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#MakeLicensesByString)
* [`cff.MakeLicenses(...License)`](https://pkg.go.dev/github.com/alexander-lindner/go-cff#MakeLicenses)
```go
content := cff.Cff{
    Version: "1.0",
    Doi:     cff.MakeDoi("10.1038/123456"),
}
content2 := cff.Cff{
    Version: "1.0",
    Doi: cff.DOI{
        RegistrantCode:     "123456",
        General:            "10",
        DirectoryIndicator: "1038",
    },
}
fmt.Println("Equal: ", reflect.DeepEqual(content, content2))
```

## The `PersonEntity` type
!!! 
For an exmplenation see [Parse](parse.md#the-personentity-type). 
!!!
The construction of a `PersonEntity` is very simple, use it like this:
```go
cff.Cff{
  Authors: []cff.PersonEntity{
      {
          Person: cff.Person{
              GivenNames: "John",
              Family:     "Doe",
          },
      },
      {
          Entity: cff.Entity{
              Address: "123 Main St",
              Email:   "test@l.de",
          },
      },
  },
}
```
Only the not empty fields are serialized, so setting `Person` and `Entity` at the same time is useless.
## Example

### Save as a string

```go
content := cff.Cff{
    Version: "1.0",
    Url:     cff.MakeUrl("http://www.google.com"),
    Doi:     cff.MakeDoi("10.1038/123456"),
}

cffContent, err := cff.Save(content)
if err != nil {
    fmt.Println("Error: ", err)
}
fmt.Println(string(cffContent))
```
Results in:
```yaml
doi: 10.1038/123456
url: http://www.google.com
version: 1.0
```
### Save as a file
```go
content := cff.Cff{
    Version: "1.0",
    Url:     cff.MakeUrl("http://www.google.com"),
    Doi:     cff.MakeDoi("10.1038/123456"),
}

err := cff.SaveFile("./CITATION.cff", content)
if err != nil {
    fmt.Println("Error: ", err)
}
```
