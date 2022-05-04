---
title: Write
icon: pencil
order: -30
---

`go-cff` provides two ways to decode a [`cff.CFF`](https://github.com/alexander-lindner/go-cff/blob/main/format.go) object:
* `cff.Save(content Cff)` for getting a string back
* `cff.SaveFile(file string, content Cff) ` for saving as a file

 [`cff.CFF`](https://github.com/alexander-lindner/go-cff/blob/main/format.go) is a struct,
 so you can easly create a new object from that struct or extend a read in object.
Empty fields are skipped.
## Helper

For providing a better experience, we use special data structures when parsing a file, for 
example `url.URL`.
However, when creating a field with such a data structure, it is often really messy.
Hence, you can use the following helper functions as alegant shortcuts:
* `cff.MakeUrl(string)`
* `cff.MakeDoi(string)`

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
