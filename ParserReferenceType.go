package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
)

//MarshalYAML serialises the ReferenceType to YAML and validates the data
func (t ReferenceType) MarshalYAML() (interface{}, error) {
	if !ReferenceTypeExists(string(t)) {
		return nil, errors.New("invalid reference type: " + string(t))
	}
	return string(t), nil
}

//UnmarshalYAML checks if the data is a valid ReferenceType
func (t *ReferenceType) UnmarshalYAML(v *yaml.Node) error {
	if !ReferenceTypeExists(v.Value) {
		return errors.New("ReferenceType does not exist")
	}
	*t = ReferenceType(v.Value)
	return nil
}

//ReferenceTypeExists checks if a ReferenceType exists
func ReferenceTypeExists(referenceType string) bool {
	switch ReferenceType(referenceType) {
	case ReferenceType_art:
		return true
	case ReferenceType_article:
		return true
	case ReferenceType_audiovisual:
		return true
	case ReferenceType_bill:
		return true
	case ReferenceType_blog:
		return true
	case ReferenceType_book:
		return true
	case ReferenceType_catalogue:
		return true
	case ReferenceType_conferencePaper:
		return true
	case ReferenceType_conference:
		return true
	case ReferenceType_data:
		return true
	case ReferenceType_database:
		return true
	case ReferenceType_dictionary:
		return true
	case ReferenceType_editedWork:
		return true
	case ReferenceType_encyclopedia:
		return true
	case ReferenceType_filmBroadcast:
		return true
	case ReferenceType_generic:
		return true
	case ReferenceType_governmentDocument:
		return true
	case ReferenceType_grant:
		return true
	case ReferenceType_hearing:
		return true
	case ReferenceType_historicalWork:
		return true
	case ReferenceType_legalCase:
		return true
	case ReferenceType_legalRule:
		return true
	case ReferenceType_magazineArticle:
		return true
	case ReferenceType_manual:
		return true
	case ReferenceType_map:
		return true
	case ReferenceType_multimedia:
		return true
	case ReferenceType_music:
		return true
	case ReferenceType_newspaperArticle:
		return true
	case ReferenceType_pamphlet:
		return true
	case ReferenceType_patent:
		return true
	case ReferenceType_personalCommunication:
		return true
	case ReferenceType_proceedings:
		return true
	case ReferenceType_report:
		return true
	case ReferenceType_serial:
		return true
	case ReferenceType_slides:
		return true
	case ReferenceType_softwareCode:
		return true
	case ReferenceType_softwareContainer:
		return true
	case ReferenceType_softwareExecutable:
		return true
	case ReferenceType_softwareVirtualMachine:
		return true
	case ReferenceType_software:
		return true
	case ReferenceType_soundRecording:
		return true
	case ReferenceType_standard:
		return true
	case ReferenceType_statute:
		return true
	case ReferenceType_thesis:
		return true
	case ReferenceType_unpublished:
		return true
	case ReferenceType_video:
		return true
	case ReferenceType_website:
		return true
	default:
		return false
	}
}
