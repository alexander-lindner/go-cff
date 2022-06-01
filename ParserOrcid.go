package cff

import (
	"regexp"
)

var regexORCID *regexp.Regexp

func init() {
	regexORCID = regexp.MustCompile(`^https://orcid\.org/[0-9]{4}-[0-9]{4}-[0-9]{4}-[0-9]{3}[0-9X]{1}$`)
}

//UnmarshalYAML parses a cff string into an ORCID object
func (j *ORCID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return CommonUnmarshalYAML(
		func(cffString string) {
			*j = ORCID(cffString)
		},
		regexORCID,
		unmarshal,
		"ORCID",
	)
}

//MarshalYAML serializes an ORCID object into a cff string
func (j ORCID) MarshalYAML() (interface{}, error) {
	parsedItem, err := CommonMarshalYAML(string(j), regexORCID, "ORCID")
	return parsedItem, err
}
