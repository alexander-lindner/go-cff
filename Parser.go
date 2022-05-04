package cff

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

// ParseFile parses a cff file to a cff.Cff struct
func ParseFile(file string) (Cff, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Couldn't read config file.", err)
	}

	return Parse(string(content))
}

// Parse parses a string in the cff format to a cff.Cff struct
func Parse(content string) (Cff, error) {
	var t Cff
	err := yaml.Unmarshal([]byte(content), &t)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return t, nil
}
