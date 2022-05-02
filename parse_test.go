package cff

import (
	"testing"
	"time"
)

func TestCffParser(t *testing.T) {
	file := "./CITATION.cff"
	parser, err := ParseFile(file)
	if err != nil {
		t.Error(err)
	}
	if len(parser.Authors) != 1 {
		t.Error("Authors not parsed correctly, should be 1, is currently", len(parser.Authors))
	}

	if parser.Authors[0].Person.Family != "Lindner" {
		t.Error("Authors not parsed correctly: Family should be Lindner, is currently", parser.Authors[0].Person.Family)
	}
	if parser.CffVersion != "1.2.0" {
		t.Error("CffVersion not parsed correctly")
	}
	parsedTime, _ := time.Parse("2006-01-02", "2022-04-30")
	if parser.DateReleased.Time() != parsedTime {
		t.Error("release date not parsed correctly")
	}
}
