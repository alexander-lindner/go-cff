package cff

import (
	"regexp"
)

var regexPMCID *regexp.Regexp

func init() {
	regexPMCID = regexp.MustCompile(`^PMC[0-9]{7}$`)
}

//UnmarshalYAML parses a cff string into a PMCID object
func (j *PMCID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return CommonUnmarshalYAML(
		func(cffString string) {
			*j = PMCID(cffString)
		},
		regexPMCID,
		unmarshal,
		"PMC ID",
	)
}

//MarshalYAML serializes a PMCID object into a cff string
func (j PMCID) MarshalYAML() (interface{}, error) {
	return CommonMarshalYAML(string(j), regexPMCID, "PMC ID")
}
