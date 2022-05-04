package cff

import (
	"gopkg.in/yaml.v3"
	"os"
)

//SaveFile saves the cff object to the given path in the cff format
func SaveFile(file string, content Cff) error {
	parsedContent, err := yaml.Marshal(content)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, parsedContent, 0644)
	return err
}

//Save is used to parse the cff object to a cff format string
func Save(content Cff) ([]byte, error) {
	parsedContent, err := yaml.Marshal(content)
	return parsedContent, err
}
