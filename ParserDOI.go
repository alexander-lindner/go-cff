package cff

import (
	"errors"
)

//parseDoi creates a DOI object from a string
func parseDoi(doi string) (general string, directoryIndicator string, registrantCode string, err error) {
	state := 0

	for _, c := range doi {
		if state == 0 && c == '.' {
			state++
			continue
		} else if state == 1 && c == '/' {
			state++
			continue
		}

		switch state {
		case 0:
			general += string(c)
		case 1:
			directoryIndicator += string(c)
		case 2:
			registrantCode += string(c)
		}
	}

	if general != "10" {
		err = errors.New("DOI does not start with 10")
		return
	} else if len(directoryIndicator) == 0 || len(registrantCode) == 0 {
		err = errors.New("directory indicator or registrant code was empty")
		return
	}
	return
}

//UnmarshalYAML parses a cff doi string into a DOI object
func (d *DOI) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var doi string
	err := unmarshal(&doi)
	if err != nil {
		return err
	}

	general, directoryIndicator, registrantCode, err := parseDoi(doi)
	if err != nil {
		return err
	}
	d.General = general
	d.DirectoryIndicator = directoryIndicator
	d.RegistrantCode = registrantCode

	return err
}

//IsValid checks if a DOI is valid
func (d DOI) IsValid() bool {
	if d.General != "10" {
		return false
	} else if len(d.DirectoryIndicator) == 0 || len(d.RegistrantCode) == 0 {
		return false
	}

	return true
}

//MarshalYAML serializes a DOI object into a cff doi string
func (d DOI) MarshalYAML() (interface{}, error) {
	if d.IsValid() {
		return d.General + "." + d.DirectoryIndicator + "/" + d.RegistrantCode, nil
	}
	return "", errors.New("DOI is invalid")
}

//MakeDoi creates a DOI object from a basic string
func MakeDoi(doi string) DOI {
	general, directoryIndicator, registrantCode, err := parseDoi(doi)
	if err != nil {
		return DOI{}
	}
	return DOI{
		General:            general,
		DirectoryIndicator: directoryIndicator,
		RegistrantCode:     registrantCode,
	}
}
