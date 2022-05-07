package cff

import (
	"regexp"
)

var regexSWHID *regexp.Regexp

func init() {
	regexSWHID = regexp.MustCompile(`^swh:1:(snp|rel|rev|dir|cnt):[0-9a-fA-F]{40}$`)
}

//UnmarshalYAML parses a cff string into a SWHID object
func (j *SWHID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return CommonUnmarshalYAML(
		func(cffString string) {
			*j = SWHID(cffString)
		},
		regexSWHID,
		unmarshal,
		"SWH Identifier",
	)
}

//MarshalYAML serializes a SWHID object into a cff string
func (j SWHID) MarshalYAML() (interface{}, error) {
	return CommonMarshalYAML(string(j), regexSWHID, "SWH Identifier")
}
