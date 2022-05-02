package cff

import (
	"fmt"
	"time"
)

type CffDate time.Time

func (t *CffDate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var unparsedString string
	if err := unmarshal(&unparsedString); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02", unparsedString)
	if err != nil {
		return fmt.Errorf("failed to parse '%s' to time.Time: %v", unparsedString, err)
	}

	*t = CffDate(parsedTime)
	return nil
}
func (t *CffDate) Time() time.Time {
	return time.Time(*t)
}
func (t *CffDate) TimeString() string {
	return t.Time().Format("2006-01-02")
}
