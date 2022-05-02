package cff

import "net/url"

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
func (j URL) MarshalYAML() (interface{}, error) {
	return j.String(), nil
}
func MakeUrl(s string) URL {
	u, _ := url.Parse(s)
	return URL{URL: u}
}
