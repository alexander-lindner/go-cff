package cff

import (
	"regexp"
)

var regexISSN *regexp.Regexp

func init() {
	regexISSN = regexp.MustCompile(`^\d{4}-\d{3}[\dxX]$`)
}

//UnmarshalYAML parses a cff string into an ISSN object
func (j *ISSN) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return CommonUnmarshalYAML(
		func(cffString string) {
			*j = ISSN(cffString)
		},
		regexISSN,
		unmarshal,
		"ISSN",
	)
}

//MarshalYAML serializes an ISSN object into a cff string
func (j ISSN) MarshalYAML() (interface{}, error) {
	return CommonMarshalYAML(string(j), regexISSN, "ISSN")
}
