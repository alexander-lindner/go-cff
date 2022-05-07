package cff

import (
	"regexp"
)

var regexLanguage *regexp.Regexp

func init() {
	regexLanguage = regexp.MustCompile(`^[a-z]{2,3}$`)
}

//UnmarshalYAML parses a cff string into a Language object
func (j *Language) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return CommonUnmarshalYAML(
		func(cffString string) {
			*j = Language(cffString)
		},
		regexLanguage,
		unmarshal,
		"language",
	)
}

//MarshalYAML serializes a Language object into a cff string
func (j Language) MarshalYAML() (interface{}, error) {
	return CommonMarshalYAML(string(j), regexLanguage, "language")
}
