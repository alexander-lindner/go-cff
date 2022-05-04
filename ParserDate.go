package cff

import (
	"fmt"
	"time"
)

//UnmarshalYAML parses a cff date string into a time.Time
func (t *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var unparsedString string
	if err := unmarshal(&unparsedString); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02", unparsedString)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to time.Time: %v", unparsedString, err)
	}

	*t = Date(parsedTime)
	return nil
}

//Time returns the time.Time
func (t *Date) Time() time.Time {
	return time.Time(*t)
}

//TimeString returns the string representation of the date
func (t *Date) TimeString() string {
	return t.Time().Format("2006-01-02")
}
