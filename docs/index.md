---
title: Home
icon: home
order: -10
---

`go-cff` is a simple library for working with the [CFF](https://citation-file-format.github.io/) format.

## Use it
```shell
go get github.com/alexander-lindner/go-cff
```
## Examples

```go
package main

import (
        "fmt"
        log "github.com/sirupsen/logrus"
        "github.com/alexander-lindner/go-cff"
)

func main() {
        file := "./CITATION.cff"
        parser, _ := cff.ParseFile(file)

        for _, item := range parser.Authors {
                if item.IsPerson {
                        fmt.Println(item.Person.GivenNames)
                } else {
                        fmt.Println(item.Entity.Address)
                        fmt.Println(item.Entity.Email)
                }
        }

        content := cff.Cff{
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
                Version: "sd",
        }
        file = "./CITATION2.cff"
        err := cff.SaveFile(file, content)
        if err != nil {
                log.Fatal(err)
        }
}
```