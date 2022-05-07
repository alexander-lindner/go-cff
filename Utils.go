package cff

import (
	"errors"
	"regexp"
)

func CommonMarshalYAML(cffString string, regexPattern *regexp.Regexp, typeOfString string) (interface{}, error) {
	var err error
	if regexPattern.MatchString(cffString) {
		return cffString, err
	} else {
		return "", errors.New("invalid " + typeOfString + ": " + cffString)
	}
}

func CommonUnmarshalYAML(matchAction func(string), regexPattern *regexp.Regexp, unmarshal func(interface{}) error, typeOfString string) error {
	var cffString string
	err := unmarshal(&cffString)
	if err != nil {
		return err
	}

	if regexPattern.MatchString(cffString) {
		matchAction(cffString)
	} else {
		err = errors.New("invalid " + typeOfString + ": " + cffString)
	}
	return err
}
