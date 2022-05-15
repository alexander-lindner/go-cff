package cff

import (
	"regexp"
)

var regexISBN *regexp.Regexp

func init() {
	regexISBN = regexp.MustCompile(`^[0-9\- ]{10,17}X?$`)
}

//UnmarshalYAML parses a cff string into an ISBN object
func (j *ISBN) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return CommonUnmarshalYAML(
		func(cffString string) {
			*j = ISBN(cffString)
		},
		regexISBN,
		unmarshal,
		"ISBN",
	)
}

//MarshalYAML serializes an ISBN object into a cff string
func (j ISBN) MarshalYAML() (interface{}, error) {
	parsedItem, err := CommonMarshalYAML(string(j), regexISBN, "ISBN")
	j = ISBN(parsedItem.(string))
	return parsedItem, err
}
