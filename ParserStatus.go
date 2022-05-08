package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
)

//MarshalYAML serialises the STATUS to YAML and validates the data
func (t STATUS) MarshalYAML() (interface{}, error) {
	if !STATUSExists(string(t)) {
		return nil, errors.New("invalid Status: " + string(t))
	}
	return string(t), nil
}

//UnmarshalYAML checks if the data is a valid STATUS
func (t *STATUS) UnmarshalYAML(v *yaml.Node) error {
	if !STATUSExists(v.Value) {
		return errors.New("STATUS does not exist")
	}
	return nil
}

//STATUSExists checks if a STATUS exists
func STATUSExists(country string) bool {
	switch STATUS(country) {
	case STATUS_abstract:
		return true
	case STATUS_advanceOnline:
		return true
	case STATUS_inPreparation:
		return true
	case STATUS_inPress:
		return true
	case STATUS_preprint:
		return true
	case STATUS_submitted:
		return true
	default:
		return false
	}
}
