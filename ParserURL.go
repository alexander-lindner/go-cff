package cff

import "net/url"

//UnmarshalYAML parses a cff url string into a URL object
func (j *URL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}
	parsedUrl, err := url.Parse(s)
	j.URL = parsedUrl
	return err
}

//MarshalYAML serializes a URL object into a cff url string
func (j URL) MarshalYAML() (interface{}, error) {
	return j.String(), nil
}

//MakeUrl creates a URL object from a basic string
func MakeUrl(s string) URL {
	u, _ := url.Parse(s)
	return URL{URL: u}
}
