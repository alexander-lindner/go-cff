package cff

import (
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
	"time"
)

var document Cff

func TestMain(m *testing.M) {
	createTestDir()
	file := "./TEST_CITATION_TEST.cff"
	doc, err := ParseFile(file)
	if err == nil {
		document = doc
		errCode := m.Run()
		removeDir()
		os.Exit(errCode)
	}
	os.Exit(-1)
}
func TestCffParser(t *testing.T) {
	if len(document.Authors) != 2 {
		t.Error("Authors not parsed correctly, should be 1, is currently", len(document.Authors))
	}
	if document.Authors[0].Person.Family != "Doe" {
		t.Error("Authors not parsed correctly: Family should be Doe, is currently", document.Authors[0].Person.Family)
	}
	if document.CffVersion != "1.2.0" {
		t.Error("CffVersion not parsed correctly")
	}
	parsedTime, _ := time.Parse("2006-01-02", "2022-04-30")
	if document.DateReleased.Time() != parsedTime {
		t.Error("release date not parsed correctly, should be 2022-04-30, is currently ", document.DateReleased.Time())
	}
}
func TestCountry(t *testing.T) {
	if document.Authors[0].Person.Country != COUNTRY_DE {
		t.Error("Country not parsed correctly, should be DE, is currently ", document.Authors[0].Person.Country)
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family:  "Doe",
					Country: "YOU",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}

	if CountryExists(string(newDocument.Authors[0].Person.Country)) {
		t.Error("Country does not exist, however no error was thrown")
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("Country should not be valid, however no error was thrown")
	}
	newDocument.Authors[0].Person.Country = COUNTRY_AF
	_, err = Save(newDocument)
	if err != nil {
		t.Error("Country is valid, however an error was thrown")
	}
	err = SaveFile("tests/Country.cff", newDocument)
	if err != nil {
		t.Error("Country is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Country")
}
func TestDate(t *testing.T) {
	if document.DateReleased.Time().Format("2006-01-02") != "2022-04-30" {
		t.Error("Date not parsed correctly, should be 2022-04-30, is currently ", document.DateReleased.Time())
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Message:      "This is a test",
		Title:        "Test",
		DateReleased: Date(time.Now()),
	}

	_, err := Save(newDocument)
	if err != nil {
		t.Error("Date should be valid, however an error was thrown")
	}
	err = SaveFile("tests/Date.cff", newDocument)
	if err != nil {
		t.Error("Date is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Date")
}
func TestDOI(t *testing.T) {
	if document.Doi != MakeDoi("10.1038/123456") {
		t.Error("DOI not parsed correctly, should be 10.1038/123456, is currently ", document.Doi)
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
		Doi:     MakeDoi("111038"),
	}
	_, _, _, err := parseDoi("111038")
	if err == nil {
		t.Error("DOI should not be valid, however no error was thrown")
	}
	_, err = ParseDoi("111038")
	if err == nil {
		t.Error("DOI should not be valid, however no error was thrown")
	}
	content, err := Save(newDocument)
	if err == nil && strings.Contains(string(content), "111038") {
		t.Error("DOI should not be valid, however doi was found in the cff file")
	}
	newDocument.Doi = MakeDoi("10.1038/123456")
	_, err = Save(newDocument)
	if err != nil {
		t.Error("DOI is valid, however an error was thrown")
	}
	err = SaveFile("tests/DOI.cff", newDocument)
	if err != nil {
		t.Error("DOI is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "DOI")
}
func TestIdentifier(t *testing.T) {
	if document.Identifiers[0].IsURL == true {
		if document.Identifiers[0].IsOther {
			t.Error("Identifier not parsed correctly, should be URL, however the internal structure is wrong: isOther= ", document.Identifiers[0].IsOther)
		}
		if document.Identifiers[0].IsDOI {
			t.Error("Identifier not parsed correctly, should be URL, however the internal structure is wrong: IsDOI= ", document.Identifiers[0].IsDOI)
		}
		if document.Identifiers[0].IsSWH {
			t.Error("Identifier not parsed correctly, should be URL, however the internal structure is wrong: IsSWH= ", document.Identifiers[0].IsSWH)
		}
		if document.Identifiers[0].URL == EmptyURL {
			t.Error("Identifier not parsed correctly, should be URL, however URL is empty ")
		}
		if document.Identifiers[0].DOI != EmptyDOI {
			t.Error("Identifier not parsed correctly, should be URL, however DOI is not empty ")
		}
		if document.Identifiers[0].SWH != EmptySWH {
			t.Error("Identifier not parsed correctly, should be URL, however SWH is not empty ")
		}
		if document.Identifiers[0].Other != EmptyOther {
			t.Error("Identifier not parsed correctly, should be URL, however Other is not empty ")
		}
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Identifiers: []Identifier{
			{
				URL: IdentifierURL{
					Value: MakeUrl("https://www.google.com"),
				},
				Other: IdentifierOther{
					Value: "https://www.google.com",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("Identifiers should not be valid, however no error was thrown")
	}
	newDocument.Identifiers[0].Other = EmptyOther
	_, err = Save(newDocument)
	if err != nil {
		t.Error("Identifiers is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/Identifiers.cff", newDocument)
	if err != nil {
		t.Error("Identifiers is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Identifiers")
}
func TestISBN(t *testing.T) {
	if document.References[0].Isbn != "9781603095075" {
		t.Error("ISBN not parsed correctly, should be en, is currently ", document.References[0].Isbn)
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		References: []Reference{
			{
				Authors: []PersonEntity{
					{
						Person: Person{
							Family: "Doe",
						},
					},
				},
				Isbn:          "abc",
				ReferenceType: ReferenceType_article,
				Title:         "Test",
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}

	_, err := Save(newDocument)
	if err == nil {
		t.Error("ISBN should not be valid, however no error was thrown")
	}
	newDocument.References[0].Isbn = "9781603095076"
	_, err = Save(newDocument)
	if err != nil {
		t.Error("ISBN is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/ISBN.cff", newDocument)
	if err != nil {
		t.Error("ISBN is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "ISBN")
}
func TestIssn(t *testing.T) {
	if document.References[0].Issn != "2475-9066" {
		t.Error("ISSN not parsed correctly, should be en, is currently ", document.References[0].Issn)
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		References: []Reference{
			{
				Authors: []PersonEntity{
					{
						Person: Person{
							Family: "Doe",
						},
					},
				},
				Issn:          "abc",
				ReferenceType: ReferenceType_article,
				Title:         "Test",
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}

	_, err := Save(newDocument)
	if err == nil {
		t.Error("ISSN should not be valid, however no error was thrown")
	}
	newDocument.References[0].Issn = "2475-9077"
	_, err = Save(newDocument)
	if err != nil {
		t.Error("ISSN is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/ISSN.cff", newDocument)
	if err != nil {
		t.Error("ISSN is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "ISSN")
}
func TestLanguage(t *testing.T) {
	if document.References[0].Languages[0] != "en" {
		t.Error("Language not parsed correctly, should be en, is currently ", document.References[0].Languages[0])
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		References: []Reference{
			{
				Authors: []PersonEntity{
					{
						Person: Person{
							Family: "Doe",
						},
					},
				},
				Languages:     []Language{"DES"},
				ReferenceType: ReferenceType_article,
				Title:         "Test",
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}

	_, err := Save(newDocument)
	if err == nil {
		t.Error("Language should not be valid, however no error was thrown")
	}
	newDocument.References[0].Languages = []Language{"de"}
	_, err = Save(newDocument)
	if err != nil {
		t.Error("Language is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/Language.cff", newDocument)
	if err != nil {
		t.Error("Language is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Language")
}
func TestORCID(t *testing.T) {
	if document.Authors[0].Person.Orcid != "https://orcid.org/0000-0003-4925-7248-s" {
		t.Error("ORCID not parsed correctly, should be https://orcid.org/0000-0003-4925-7248-s, is currently", document.CffType)
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
					Orcid:  "doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("ORCID should not be valid, however no error was thrown")
	}
	newDocument.Authors[0].Person.Orcid = "https://orcid.org/0000-0003-4925-7248"
	_, err = Save(newDocument)
	if err != nil {
		t.Error("ORCID is valid, however an error was thrown")
	}
	err = SaveFile("tests/ORCID.cff", newDocument)
	if err != nil {
		t.Error("ORCID is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "ORCID")
}
func TestPerson(t *testing.T) {

	if document.Authors[0].IsPerson {
		if document.Authors[0].IsEntity {
			t.Error("Authors not parsed correctly, should be Person, however the internal structure is wrong: IsEntity= ", document.Authors[0].IsEntity)
		}
		if document.Authors[0].Person == EmptyPerson {
			t.Error("Authors not parsed correctly, should be Person, however Person is empty ")
		}
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
				Entity: Entity{
					Name: "Doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("PersonEntity should not be valid, however no error was thrown")
	}
	newDocument.Authors[0].Entity = EmptyEntity
	_, err = Save(newDocument)
	if err != nil {
		t.Error("PersonEntity is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/PersonEntity.cff", newDocument)
	if err != nil {
		t.Error("PersonEntity is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "PersonEntity")

}
func TestEntity(t *testing.T) {

	if document.Authors[1].IsEntity {
		if document.Authors[1].IsPerson {
			t.Error("Authors not parsed correctly, should be Person, however the internal structure is wrong: IsEntity= ", document.Authors[1].IsPerson)
		}
		if document.Authors[1].Entity == EmptyEntity {
			t.Error("Authors not parsed correctly, should be Entity, however Entity is empty ")
		}
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
				Entity: Entity{
					Name: "Doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("PersonEntity should not be valid, however no error was thrown")
	}
	newDocument.Authors[0].Person = EmptyPerson
	_, err = Save(newDocument)
	if err != nil {
		t.Error("PersonEntity is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/PersonEntity2.cff", newDocument)
	if err != nil {
		t.Error("PersonEntity is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "PersonEntity2")

}
func TestPmcid(t *testing.T) {
	if document.References[0].Pmcid != "PMC3134971" {
		t.Error("Pmcid not parsed correctly, should be PMC3134971, is currently", document.CffType)
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		References: []Reference{
			{
				Authors: []PersonEntity{
					{
						Person: Person{
							Family: "Doe",
						},
					},
				},
				Languages:     []Language{"en"},
				Title:         "Test",
				Pmcid:         "sdfsdf",
				ReferenceType: ReferenceType_conference,
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("Pmcid should not be valid, however no error was thrown")
	}
	newDocument.References[0].Pmcid = "PMC3134971"
	_, err = Save(newDocument)
	if err != nil {
		t.Error("Pmcid is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/Pmcid.cff", newDocument)
	if err != nil {
		t.Error("Pmcid is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Pmcid")
}
func TestReferenceType(t *testing.T) {
	if document.References[0].ReferenceType != ReferenceType_conferencePaper {
		t.Error("ReferenceType not parsed correctly, should be "+ReferenceType_conferencePaper+", is currently ", document.References[0].ReferenceType)
	}
	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		References: []Reference{
			{
				Authors: []PersonEntity{
					{
						Person: Person{
							Family: "Doe",
						},
					},
				},
				Languages:     []Language{"en"},
				ReferenceType: "test",
				Title:         "Test",
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	if ReferenceTypeExists(string(newDocument.References[0].ReferenceType)) {
		t.Error("ReferenceType does not exist, however no error was thrown")
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("ReferenceType should not be valid, however no error was thrown")
	}
	newDocument.References[0].ReferenceType = ReferenceType_blog
	_, err = Save(newDocument)
	if err != nil {
		t.Error("ReferenceType is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/ReferenceType.cff", newDocument)
	if err != nil {
		t.Error("ReferenceType is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "ReferenceType")
}
func TestStatus(t *testing.T) {
	if document.References[0].Status != STATUS_submitted {
		t.Error("Status not parsed correctly, should be PMC3134971, is currently", document.References[0].Status)
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		References: []Reference{
			{
				Authors: []PersonEntity{
					{
						Person: Person{
							Family: "Doe",
						},
					},
				},
				Languages:     []Language{"en"},
				Title:         "Test",
				Status:        "sdfsdf",
				ReferenceType: ReferenceType_conference,
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("Status should not be valid, however no error was thrown")
	}
	newDocument.References[0].Status = STATUS_inPreparation
	_, err = Save(newDocument)
	if err != nil {
		t.Error("Status is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/Status.cff", newDocument)
	if err != nil {
		t.Error("Status is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Status")
}
func TestSwh(t *testing.T) {
	if !document.References[0].Identifiers[2].IsSWH && document.References[0].Identifiers[2].SWH.Value != "swh:1:dir:bc286860f423ea7ced246ba7458eef4b4541cf2d" {
		t.Error("SWH not parsed correctly, should be swh:1:dir:bc286860f423ea7ced246ba7458eef4b4541cf2d, is currently", document.References[0].Identifiers[2].SWH.Value)
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Identifiers: []Identifier{
			{
				IsSWH: true,
				SWH: IdentifierSWH{
					Value: "d",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("SWH should not be valid, however no error was thrown")
	}
	newDocument.Identifiers[0].SWH.Value = "swh:1:dir:bc286860f423ea7ced246ba7458eef4b4541cf2d"
	_, err = Save(newDocument)
	if err != nil {
		t.Error("SWH is valid, however an error was thrown: ", err)
	}
	err = SaveFile("tests/SWH.cff", newDocument)
	if err != nil {
		t.Error("SWH is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "SWH")
}
func TestURL(t *testing.T) {
	if !reflect.DeepEqual(document.Url, MakeUrl("http://www.google.com")) {
		t.Error("Url not parsed correctly, should be "+MakeUrl("http://www.google.com").String()+", is currently", document.Url)
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
		Url:     MakeUrl("sdadas"),
	}
	_, err := Save(newDocument)
	if err != nil || newDocument.Url.String() == "" {
		t.Error("Url should not be valid, however no error was thrown")
	}
	newDocument.Url = MakeUrl("http://www.google.com")
	_, err = Save(newDocument)
	if err != nil {
		t.Error("Url is valid, however an error was thrown")
	}
	err = SaveFile("tests/Url.cff", newDocument)
	if err != nil {
		t.Error("Url is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "Url")
}
func TestCffType(t *testing.T) {
	if document.CffType != CFFTYPE_Software {
		t.Error("CffType not parsed correctly, should be CFFTYPE_Software, is currently", document.CffType)
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
		CffType: "CFFTYPE_Demo",
	}
	if CFFTYPEExists(string(newDocument.CffType)) {
		t.Error("CffType does not exist, however no error was thrown")
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("CffType should not be valid, however no error was thrown")
	}
	newDocument.CffType = CFFTYPE_Dataset
	_, err = Save(newDocument)
	if err != nil {
		t.Error("CffType is valid, however an error was thrown")
	}
	err = SaveFile("tests/CffType.cff", newDocument)
	if err != nil {
		t.Error("CffType is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "CffType")

}
func TestLicense(t *testing.T) {
	if len(document.License.Data) != 1 {
		t.Error("Licenses not parsed correctly, should be 1, is currently", len(document.License.Data))
	}
	if !LicenseExists(string(document.License.Data[0])) {
		t.Error("License not valid. However, it should be valid as no error was thrown before")
	}
	if LicenseExists("DEMO") {
		t.Error("License DEMO should not be valid")
	}
	if !LicenseExists(string(LICENSE_AdobeGlyph)) {
		t.Error("License AdobeGlyph should be valid")
	}

	newDocument := Cff{
		Authors: []PersonEntity{
			{
				Person: Person{
					Family: "Doe",
				},
			},
		},
		Message: "This is a test",
		Title:   "Test",
		License: MakeLicensesByString("DEMO"),
	}
	if LicenseExists(string(newDocument.License.Data[0])) {
		t.Error("License DEMO should not be valid")
	}
	_, err := Save(newDocument)
	if err == nil {
		t.Error("License DEMO should not be valid, however no error was thrown")
	}
	newDocument.License = MakeLicenses(LICENSE_AdobeGlyph)
	_, err = Save(newDocument)
	if err != nil {
		t.Error("License is valid, however an error was thrown")
	}
	err = SaveFile("tests/license.cff", newDocument)
	if err != nil {
		t.Error("License is valid, however an error was thrown: ", err)
	}
	validTestCffFile(t, "license")
}

//validTestCffFile runs a python executable to check if it is valid
func validTestCffFile(t *testing.T, dir string) {
	t.Log("Validating test file: ~/.local/bin/cffconvert --validate -i ./tests/" + dir + ".cff")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Error("Could not get user home directory")
	}
	cmd := exec.Command(homeDir+"/.local/bin/cffconvert", "--validate", "-i", "./tests/"+dir+".cff")
	out, err := cmd.Output()
	if err != nil {
		t.Error("Written cff file is not valid. Error: ", err)
	}
	t.Log(string(out))
}

func createTestDir() {
	if _, err := os.Stat("tests"); os.IsNotExist(err) {
		err := os.Mkdir("./tests/", os.ModePerm)
		if err != nil {
			os.Exit(-1)
			return
		}
	}
}

func removeDir() {
	err := os.RemoveAll("./tests/")
	if err != nil {
		os.Exit(-1)
		return
	}
}
