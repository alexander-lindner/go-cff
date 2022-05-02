package cff

import (
	"gopkg.in/yaml.v3"
	"os"
)

func SaveFile(file string, content Cff) error {
	parsedContent, err := yaml.Marshal(content)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, parsedContent, 0644)
	return err
}
