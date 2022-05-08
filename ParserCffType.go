package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
)

//MarshalYAML serialises the CFFTYPE to YAML and validates the data
func (t CFFTYPE) MarshalYAML() (interface{}, error) {
	if !CFFTYPEExists(string(t)) {
		return nil, errors.New("invalid CFFTYPE: " + string(t))
	}
	return string(t), nil
}

//UnmarshalYAML checks if the data is a valid CFFTYPE
func (t *CFFTYPE) UnmarshalYAML(v *yaml.Node) error {
	if !CFFTYPEExists(v.Value) {
		return errors.New("CFFTYPE does not exist")
	}
	return nil
}

//CFFTYPEExists checks if a CFFTYPE exists
func CFFTYPEExists(country string) bool {
	switch CFFTYPE(country) {
	case CFFTYPE_Software:
		return true
	case CFFTYPE_Dataset:
		return true
	default:
		return false
	}
}
