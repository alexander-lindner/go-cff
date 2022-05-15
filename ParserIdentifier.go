package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
	"strings"
)

//EmptySWH represents an empty Identifier of type SWH
var EmptySWH IdentifierSWH

//EmptyURL represents an empty Identifier of type URL
var EmptyURL IdentifierURL

//EmptyDOI represents an empty Identifier of type DOI
var EmptyDOI IdentifierDOI

//EmptyOther represents an empty Identifier of type Other
var EmptyOther IdentifierOther

//Get is used to retrieve a DOI Identifier or a URL Identifier or a SWH Identifier or an Other Identifier
func (t *Identifier) Get() (IdentifierDOI, IdentifierURL, IdentifierSWH, IdentifierOther) {
	return t.DOI, t.URL, t.SWH, t.Other
}

//MarshalYAML is used to serialise a PersonEntity
func (t Identifier) MarshalYAML() (interface{}, error) {
	trueArray := []bool{
		t.DOI != EmptyDOI,
		t.URL != EmptyURL,
		t.SWH != EmptySWH,
		t.Other != EmptyOther,
	}
	counter := 0
	for _, element := range trueArray {
		if element {
			counter++
		}
	}
	if counter == 0 {
		return nil, errors.New("identifier is empty")
	} else if counter > 1 {
		return nil, errors.New("identifier contains more than one type")
	}
	if t.IsSWH || (t.SWH != EmptySWH) {
		t.SWH.IdentifiersType = "swh"
		return t.SWH, nil
	} else if t.IsDOI || (t.DOI != EmptyDOI) {
		t.DOI.IdentifiersType = "doi"
		return t.DOI, nil
	} else if t.IsURL || (t.URL != EmptyURL) {
		t.URL.IdentifiersType = "url"
		return t.URL, nil
	} else {
		return t.Other, nil
	}
}

//UnmarshalYAML creates a PersonEntity from a cff Person or Entity
func (t *Identifier) UnmarshalYAML(v *yaml.Node) error {

	t.IsDOI = false
	t.IsURL = false
	t.IsSWH = false
	var test IdentifierTest
	err := v.Decode(&test)
	if err != nil {
		return err
	}
	test.IdentifiersType = strings.ToUpper(test.IdentifiersType)
	if test.IdentifiersType == "DOI" {
		t.IsDOI = true

		var doi IdentifierDOI
		err := v.Decode(&doi)
		if err != nil {
			return err
		}
		t.DOI = doi
		return nil
	} else if test.IdentifiersType == "URL" {
		t.IsURL = true

		var url IdentifierURL
		err := v.Decode(&url)
		if err != nil {
			return err
		}
		t.URL = url
		return nil
	} else if test.IdentifiersType == "SWH" {
		t.IsSWH = true

		var SWH IdentifierSWH
		err := v.Decode(&SWH)
		if err != nil {
			return err
		}
		t.SWH = SWH
		return nil
	} else if test.IdentifiersType == "OTHER" {
		t.IsOther = true
		var other IdentifierOther
		err := v.Decode(&other)
		if err != nil {
			return err
		}
		t.Other = other
		return nil
	}

	return errors.New("unknown Identifier")
}
