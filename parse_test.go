package cff

import (
	"testing"
	"time"
)

func TestCffParser(t *testing.T) {
	file := "./TEST_CITATION_TEST.cff"
	parser, err := ParseFile(file)
	if err != nil {
		t.Error(err)
	}
	if len(parser.Authors) != 2 {
		t.Error("Authors not parsed correctly, should be 1, is currently", len(parser.Authors))
	}

	if parser.Authors[0].Person.Family != "Doe" {
		t.Error("Authors not parsed correctly: Family should be Doe, is currently", parser.Authors[0].Person.Family)
	}
	if parser.CffVersion != "1.2.0" {
		t.Error("CffVersion not parsed correctly")
	}
	parsedTime, _ := time.Parse("2006-01-02", "2022-04-30")
	if parser.DateReleased.Time() != parsedTime {
		t.Error("release date not parsed correctly, should be 2022-04-30, is currently", parser.DateReleased.Time())
	}
}
