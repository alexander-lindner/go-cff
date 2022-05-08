package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
)

//MarshalYAML serialises the TYPE to YAML and validates the data
func (t TYPE) MarshalYAML() (interface{}, error) {
	if !TYPEExists(string(t)) {
		return nil, errors.New("invalid type: " + string(t))
	}
	return string(t), nil
}

//UnmarshalYAML checks if the data is a valid TYPE
func (t *TYPE) UnmarshalYAML(v *yaml.Node) error {
	if !TYPEExists(v.Value) {
		return errors.New("TYPE does not exist")
	}
	return nil
}

//TYPEExists checks if a TYPE exists
func TYPEExists(country string) bool {
	switch TYPE(country) {
	case TYPE_art:
		return true
	case TYPE_article:
		return true
	case TYPE_audiovisual:
		return true
	case TYPE_bill:
		return true
	case TYPE_blog:
		return true
	case TYPE_book:
		return true
	case TYPE_catalogue:
		return true
	case TYPE_conferencePaper:
		return true
	case TYPE_conference:
		return true
	case TYPE_data:
		return true
	case TYPE_database:
		return true
	case TYPE_dictionary:
		return true
	case TYPE_editedWork:
		return true
	case TYPE_encyclopedia:
		return true
	case TYPE_filmBroadcast:
		return true
	case TYPE_generic:
		return true
	case TYPE_governmentDocument:
		return true
	case TYPE_grant:
		return true
	case TYPE_hearing:
		return true
	case TYPE_historicalWork:
		return true
	case TYPE_legalCase:
		return true
	case TYPE_legalRule:
		return true
	case TYPE_magazineArticle:
		return true
	case TYPE_manual:
		return true
	case TYPE_map:
		return true
	case TYPE_multimedia:
		return true
	case TYPE_music:
		return true
	case TYPE_newspaperArticle:
		return true
	case TYPE_pamphlet:
		return true
	case TYPE_patent:
		return true
	case TYPE_personalCommunication:
		return true
	case TYPE_proceedings:
		return true
	case TYPE_report:
		return true
	case TYPE_serial:
		return true
	case TYPE_slides:
		return true
	case TYPE_softwareCode:
		return true
	case TYPE_softwareContainer:
		return true
	case TYPE_softwareExecutable:
		return true
	case TYPE_softwareVirtualMachine:
		return true
	case TYPE_software:
		return true
	case TYPE_soundRecording:
		return true
	case TYPE_standard:
		return true
	case TYPE_statute:
		return true
	case TYPE_thesis:
		return true
	case TYPE_unpublished:
		return true
	case TYPE_video:
		return true
	case TYPE_website:
		return true
	default:
		return false
	}
}
