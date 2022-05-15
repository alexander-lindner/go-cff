package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

//SaveFile saves the cff object to the given path in the cff format
func SaveFile(file string, content Cff) error {
	parsedContent, err := Save(content)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, parsedContent, 0644)
	return err
}

//Save is used to parse the cff object to a cff format string
func Save(content Cff) ([]byte, error) {
	content.CffVersion = "1.2.0"
	if len(content.Authors) == 0 {
		return nil, errors.New("author is required")
	}
	if len(content.Message) == 0 {
		return nil, errors.New("message is required")
	}
	if len(content.Title) == 0 {
		return nil, errors.New("title is required")
	}
	if len(content.References) > 0 {
		for _, reference := range content.References {
			if len(reference.Title) == 0 {
				return nil, errors.New("reference.title is required")
			}
			if len(reference.ReferenceType) == 0 {
				return nil, errors.New("reference.type is required")
			}
			if len(reference.Authors) == 0 {
				return nil, errors.New("reference.author is required")
			}
		}
	}
	parsedContent, err := yaml.Marshal(content)
	return parsedContent, err
}
