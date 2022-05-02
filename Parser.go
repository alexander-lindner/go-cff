package cff

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseFile(file string) (Cff, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Couldn't read config file.", err)
	}

	return Parse(string(content))
}

func Parse(content string) (Cff, error) {
	var t Cff
	err := yaml.Unmarshal([]byte(content), &t)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return t, nil
}
