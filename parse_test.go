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
	countries := []Country{
		COUNTRY_AD,
		COUNTRY_AE,
		COUNTRY_AF,
		COUNTRY_AG,
		COUNTRY_AI,
		COUNTRY_AL,
		COUNTRY_AM,
		COUNTRY_AO,
		COUNTRY_AQ,
		COUNTRY_AR,
		COUNTRY_AS,
		COUNTRY_AT,
		COUNTRY_AU,
		COUNTRY_AW,
		COUNTRY_AX,
		COUNTRY_AZ,
		COUNTRY_BA,
		COUNTRY_BB,
		COUNTRY_BD,
		COUNTRY_BE,
		COUNTRY_BF,
		COUNTRY_BG,
		COUNTRY_BH,
		COUNTRY_BI,
		COUNTRY_BJ,
		COUNTRY_BL,
		COUNTRY_BM,
		COUNTRY_BN,
		COUNTRY_BO,
		COUNTRY_BQ,
		COUNTRY_BR,
		COUNTRY_BS,
		COUNTRY_BT,
		COUNTRY_BV,
		COUNTRY_BW,
		COUNTRY_BY,
		COUNTRY_BZ,
		COUNTRY_CA,
		COUNTRY_CC,
		COUNTRY_CD,
		COUNTRY_CF,
		COUNTRY_CG,
		COUNTRY_CH,
		COUNTRY_CI,
		COUNTRY_CK,
		COUNTRY_CL,
		COUNTRY_CM,
		COUNTRY_CN,
		COUNTRY_CO,
		COUNTRY_CR,
		COUNTRY_CU,
		COUNTRY_CV,
		COUNTRY_CW,
		COUNTRY_CX,
		COUNTRY_CY,
		COUNTRY_CZ,
		COUNTRY_DE,
		COUNTRY_DJ,
		COUNTRY_DK,
		COUNTRY_DM,
		COUNTRY_DO,
		COUNTRY_DZ,
		COUNTRY_EC,
		COUNTRY_EE,
		COUNTRY_EG,
		COUNTRY_EH,
		COUNTRY_ER,
		COUNTRY_ES,
		COUNTRY_ET,
		COUNTRY_FI,
		COUNTRY_FJ,
		COUNTRY_FK,
		COUNTRY_FM,
		COUNTRY_FO,
		COUNTRY_FR,
		COUNTRY_GA,
		COUNTRY_GB,
		COUNTRY_GD,
		COUNTRY_GE,
		COUNTRY_GF,
		COUNTRY_GG,
		COUNTRY_GH,
		COUNTRY_GI,
		COUNTRY_GL,
		COUNTRY_GM,
		COUNTRY_GN,
		COUNTRY_GP,
		COUNTRY_GQ,
		COUNTRY_GR,
		COUNTRY_GS,
		COUNTRY_GT,
		COUNTRY_GU,
		COUNTRY_GW,
		COUNTRY_GY,
		COUNTRY_HK,
		COUNTRY_HM,
		COUNTRY_HN,
		COUNTRY_HR,
		COUNTRY_HT,
		COUNTRY_HU,
		COUNTRY_ID,
		COUNTRY_IE,
		COUNTRY_IL,
		COUNTRY_IM,
		COUNTRY_IN,
		COUNTRY_IO,
		COUNTRY_IQ,
		COUNTRY_IR,
		COUNTRY_IS,
		COUNTRY_IT,
		COUNTRY_JE,
		COUNTRY_JM,
		COUNTRY_JO,
		COUNTRY_JP,
		COUNTRY_KE,
		COUNTRY_KG,
		COUNTRY_KH,
		COUNTRY_KI,
		COUNTRY_KM,
		COUNTRY_KN,
		COUNTRY_KP,
		COUNTRY_KR,
		COUNTRY_KW,
		COUNTRY_KY,
		COUNTRY_KZ,
		COUNTRY_LA,
		COUNTRY_LB,
		COUNTRY_LC,
		COUNTRY_LI,
		COUNTRY_LK,
		COUNTRY_LR,
		COUNTRY_LS,
		COUNTRY_LT,
		COUNTRY_LU,
		COUNTRY_LV,
		COUNTRY_LY,
		COUNTRY_MA,
		COUNTRY_MC,
		COUNTRY_MD,
		COUNTRY_ME,
		COUNTRY_MF,
		COUNTRY_MG,
		COUNTRY_MH,
		COUNTRY_MK,
		COUNTRY_ML,
		COUNTRY_MM,
		COUNTRY_MN,
		COUNTRY_MO,
		COUNTRY_MP,
		COUNTRY_MQ,
		COUNTRY_MR,
		COUNTRY_MS,
		COUNTRY_MT,
		COUNTRY_MU,
		COUNTRY_MV,
		COUNTRY_MW,
		COUNTRY_MX,
		COUNTRY_MY,
		COUNTRY_MZ,
		COUNTRY_NA,
		COUNTRY_NC,
		COUNTRY_NE,
		COUNTRY_NF,
		COUNTRY_NG,
		COUNTRY_NI,
		COUNTRY_NL,
		COUNTRY_NO,
		COUNTRY_NP,
		COUNTRY_NR,
		COUNTRY_NU,
		COUNTRY_NZ,
		COUNTRY_OM,
		COUNTRY_PA,
		COUNTRY_PE,
		COUNTRY_PF,
		COUNTRY_PG,
		COUNTRY_PH,
		COUNTRY_PK,
		COUNTRY_PL,
		COUNTRY_PM,
		COUNTRY_PN,
		COUNTRY_PR,
		COUNTRY_PS,
		COUNTRY_PT,
		COUNTRY_PW,
		COUNTRY_PY,
		COUNTRY_QA,
		COUNTRY_RE,
		COUNTRY_RO,
		COUNTRY_RS,
		COUNTRY_RU,
		COUNTRY_RW,
		COUNTRY_SA,
		COUNTRY_SB,
		COUNTRY_SC,
		COUNTRY_SD,
		COUNTRY_SE,
		COUNTRY_SG,
		COUNTRY_SH,
		COUNTRY_SI,
		COUNTRY_SJ,
		COUNTRY_SK,
		COUNTRY_SL,
		COUNTRY_SM,
		COUNTRY_SN,
		COUNTRY_SO,
		COUNTRY_SR,
		COUNTRY_SS,
		COUNTRY_ST,
		COUNTRY_SV,
		COUNTRY_SX,
		COUNTRY_SY,
		COUNTRY_SZ,
		COUNTRY_TC,
		COUNTRY_TD,
		COUNTRY_TF,
		COUNTRY_TG,
		COUNTRY_TH,
		COUNTRY_TJ,
		COUNTRY_TK,
		COUNTRY_TL,
		COUNTRY_TM,
		COUNTRY_TN,
		COUNTRY_TO,
		COUNTRY_TR,
		COUNTRY_TT,
		COUNTRY_TV,
		COUNTRY_TW,
		COUNTRY_TZ,
		COUNTRY_UA,
		COUNTRY_UG,
		COUNTRY_UM,
		COUNTRY_US,
		COUNTRY_UY,
		COUNTRY_UZ,
		COUNTRY_VA,
		COUNTRY_VC,
		COUNTRY_VE,
		COUNTRY_VG,
		COUNTRY_VI,
		COUNTRY_VN,
		COUNTRY_VU,
		COUNTRY_WF,
		COUNTRY_WS,
		COUNTRY_YE,
		COUNTRY_YT,
		COUNTRY_ZA,
		COUNTRY_ZM,
		COUNTRY_ZW,
	}
	for _, country := range countries {
		newDocument.Authors[0].Person.Country = country
		_, err = Save(newDocument)
		if err != nil {
			t.Error("Country is valid, however an error was thrown")
		}
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
	if document.Authors[0].Person.Orcid != "https://orcid.org/0000-0003-4925-7248" {
		t.Error("ORCID not parsed correctly, should be https://orcid.org/0000-0003-4925-7248, is currently", document.CffType)
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

	licenses := []License{
		LICENSE_0BSD,
		LICENSE_AAL,
		LICENSE_Abstyles,
		LICENSE_Adobe2006,
		LICENSE_AdobeGlyph,
		LICENSE_ADSL,
		LICENSE_AFL11,
		LICENSE_AFL12,
		LICENSE_AFL20,
		LICENSE_AFL21,
		LICENSE_AFL30,
		LICENSE_Afmparse,
		LICENSE_AGPL10,
		LICENSE_AGPL10only,
		LICENSE_AGPL10orlater,
		LICENSE_AGPL30only,
		LICENSE_AGPL30orlater,
		LICENSE_AGPL30,
		LICENSE_Aladdin,
		LICENSE_AMDPLPA,
		LICENSE_AML,
		LICENSE_AMPAS,
		LICENSE_ANTLRPD,
		LICENSE_ANTLRPDfallback,
		LICENSE_Apache10,
		LICENSE_Apache11,
		LICENSE_Apache20,
		LICENSE_APAFML,
		LICENSE_APL10,
		LICENSE_Apps2p,
		LICENSE_APSL10,
		LICENSE_APSL11,
		LICENSE_APSL12,
		LICENSE_APSL20,
		LICENSE_Artistic10,
		LICENSE_Artistic10cl8,
		LICENSE_Artistic10Perl,
		LICENSE_Artistic20,
		LICENSE_Bahyph,
		LICENSE_Barr,
		LICENSE_Beerware,
		LICENSE_BitTorrent10,
		LICENSE_BitTorrent11,
		LICENSE_blessing,
		LICENSE_BlueOak100,
		LICENSE_Borceux,
		LICENSE_BSD1Clause,
		LICENSE_BSD2Clause,
		LICENSE_BSD2ClausePatent,
		LICENSE_BSD2ClauseViews,
		LICENSE_BSD2ClauseFreeBSD,
		LICENSE_BSD2ClauseNetBSD,
		LICENSE_BSD3Clause,
		LICENSE_BSD3ClauseAttribution,
		LICENSE_BSD3ClauseClear,
		LICENSE_BSD3ClauseLBNL,
		LICENSE_BSD3ClauseModification,
		LICENSE_BSD3ClauseNoMilitaryLicense,
		LICENSE_BSD3ClauseNoNuclearLicense,
		LICENSE_BSD3ClauseNoNuclearLicense2014,
		LICENSE_BSD3ClauseNoNuclearWarranty,
		LICENSE_BSD3ClauseOpenMPI,
		LICENSE_BSD4Clause,
		LICENSE_BSD4ClauseShortened,
		LICENSE_BSD4ClauseUC,
		LICENSE_BSDProtection,
		LICENSE_BSDSourceCode,
		LICENSE_BSL10,
		LICENSE_BUSL11,
		LICENSE_bzip2105,
		LICENSE_bzip2106,
		LICENSE_CUDA10,
		LICENSE_CAL10,
		LICENSE_CAL10CombinedWorkException,
		LICENSE_Caldera,
		LICENSE_CATOSL11,
		LICENSE_CCBY10,
		LICENSE_CCBY20,
		LICENSE_CCBY25,
		LICENSE_CCBY25AU,
		LICENSE_CCBY30,
		LICENSE_CCBY30AT,
		LICENSE_CCBY30DE,
		LICENSE_CCBY30NL,
		LICENSE_CCBY30US,
		LICENSE_CCBY40,
		LICENSE_CCBYNC10,
		LICENSE_CCBYNC20,
		LICENSE_CCBYNC25,
		LICENSE_CCBYNC30,
		LICENSE_CCBYNC30DE,
		LICENSE_CCBYNC40,
		LICENSE_CCBYNCND10,
		LICENSE_CCBYNCND20,
		LICENSE_CCBYNCND25,
		LICENSE_CCBYNCND30,
		LICENSE_CCBYNCND30DE,
		LICENSE_CCBYNCND30IGO,
		LICENSE_CCBYNCND40,
		LICENSE_CCBYNCSA10,
		LICENSE_CCBYNCSA20,
		LICENSE_CCBYNCSA20FR,
		LICENSE_CCBYNCSA20UK,
		LICENSE_CCBYNCSA25,
		LICENSE_CCBYNCSA30,
		LICENSE_CCBYNCSA30DE,
		LICENSE_CCBYNCSA30IGO,
		LICENSE_CCBYNCSA40,
		LICENSE_CCBYND10,
		LICENSE_CCBYND20,
		LICENSE_CCBYND25,
		LICENSE_CCBYND30,
		LICENSE_CCBYND30DE,
		LICENSE_CCBYND40,
		LICENSE_CCBYSA10,
		LICENSE_CCBYSA20,
		LICENSE_CCBYSA20UK,
		LICENSE_CCBYSA21JP,
		LICENSE_CCBYSA25,
		LICENSE_CCBYSA30,
		LICENSE_CCBYSA30AT,
		LICENSE_CCBYSA30DE,
		LICENSE_CCBYSA40,
		LICENSE_CCPDDC,
		LICENSE_CC010,
		LICENSE_CDDL10,
		LICENSE_CDDL11,
		LICENSE_CDL10,
		LICENSE_CDLAPermissive10,
		LICENSE_CDLAPermissive20,
		LICENSE_CDLASharing10,
		LICENSE_CECILL10,
		LICENSE_CECILL11,
		LICENSE_CECILL20,
		LICENSE_CECILL21,
		LICENSE_CECILLB,
		LICENSE_CECILLC,
		LICENSE_CERNOHL11,
		LICENSE_CERNOHL12,
		LICENSE_CERNOHLP20,
		LICENSE_CERNOHLS20,
		LICENSE_CERNOHLW20,
		LICENSE_ClArtistic,
		LICENSE_CNRIJython,
		LICENSE_CNRIPython,
		LICENSE_CNRIPythonGPLCompatible,
		LICENSE_COIL10,
		LICENSE_CommunitySpec10,
		LICENSE_Condor11,
		LICENSE_copyleftnext030,
		LICENSE_copyleftnext031,
		LICENSE_CPAL10,
		LICENSE_CPL10,
		LICENSE_CPOL102,
		LICENSE_Crossword,
		LICENSE_CrystalStacker,
		LICENSE_CUAOPL10,
		LICENSE_Cube,
		LICENSE_curl,
		LICENSE_DFSL10,
		LICENSE_diffmark,
		LICENSE_DLDEBY20,
		LICENSE_DOC,
		LICENSE_Dotseqn,
		LICENSE_DRL10,
		LICENSE_DSDP,
		LICENSE_dvipdfm,
		LICENSE_ECL10,
		LICENSE_ECL20,
		LICENSE_eCos20,
		LICENSE_EFL10,
		LICENSE_EFL20,
		LICENSE_eGenix,
		LICENSE_Elastic20,
		LICENSE_Entessa,
		LICENSE_EPICS,
		LICENSE_EPL10,
		LICENSE_EPL20,
		LICENSE_ErlPL11,
		LICENSE_etalab20,
		LICENSE_EUDatagrid,
		LICENSE_EUPL10,
		LICENSE_EUPL11,
		LICENSE_EUPL12,
		LICENSE_Eurosym,
		LICENSE_Fair,
		LICENSE_FDKAAC,
		LICENSE_Frameworx10,
		LICENSE_FreeBSDDOC,
		LICENSE_FreeImage,
		LICENSE_FSFAP,
		LICENSE_FSFUL,
		LICENSE_FSFULLR,
		LICENSE_FTL,
		LICENSE_GD,
		LICENSE_GFDL11,
		LICENSE_GFDL11invariantsonly,
		LICENSE_GFDL11invariantsorlater,
		LICENSE_GFDL11noinvariantsonly,
		LICENSE_GFDL11noinvariantsorlater,
		LICENSE_GFDL11only,
		LICENSE_GFDL11orlater,
		LICENSE_GFDL12invariantsonly,
		LICENSE_GFDL12invariantsorlater,
		LICENSE_GFDL12noinvariantsonly,
		LICENSE_GFDL12noinvariantsorlater,
		LICENSE_GFDL12only,
		LICENSE_GFDL12orlater,
		LICENSE_GFDL12,
		LICENSE_GFDL13invariantsonly,
		LICENSE_GFDL13invariantsorlater,
		LICENSE_GFDL13noinvariantsonly,
		LICENSE_GFDL13noinvariantsorlater,
		LICENSE_GFDL13only,
		LICENSE_GFDL13orlater,
		LICENSE_GFDL13,
		LICENSE_Giftware,
		LICENSE_GL2PS,
		LICENSE_Glide,
		LICENSE_Glulxe,
		LICENSE_GLWTPL,
		LICENSE_gnuplot,
		LICENSE_GPL10,
		LICENSE_GPL10only,
		LICENSE_GPL10orlater,
		LICENSE_GPL10PLUS,
		LICENSE_GPL20,
		LICENSE_GPL20only,
		LICENSE_GPL20orlater,
		LICENSE_GPL20withautoconfexception,
		LICENSE_GPL20withbisonexception,
		LICENSE_GPL20withclasspathexception,
		LICENSE_GPL20withfontexception,
		LICENSE_GPL20withGCCexception,
		LICENSE_GPL20PLUS,
		LICENSE_GPL30,
		LICENSE_GPL30only,
		LICENSE_GPL30orlater,
		LICENSE_GPL30withautoconfexception,
		LICENSE_GPL30withGCCexception,
		LICENSE_GPL30PLUS,
		LICENSE_gSOAP13b,
		LICENSE_HaskellReport,
		LICENSE_Hippocratic21,
		LICENSE_HPND,
		LICENSE_HPNDsellvariant,
		LICENSE_HTMLTIDY,
		LICENSE_IBMpibs,
		LICENSE_ICU,
		LICENSE_IJG,
		LICENSE_ImageMagick,
		LICENSE_iMatix,
		LICENSE_Imlib2,
		LICENSE_InfoZIP,
		LICENSE_Intel,
		LICENSE_IntelACPI,
		LICENSE_Interbase10,
		LICENSE_IPA,
		LICENSE_IPL10,
		LICENSE_ISC,
		LICENSE_Jam,
		LICENSE_JasPer20,
		LICENSE_JPNIC,
		LICENSE_JSON,
		LICENSE_LAL12,
		LICENSE_LAL13,
		LICENSE_Latex2e,
		LICENSE_Leptonica,
		LICENSE_LGPL20,
		LICENSE_LGPL20PLUS,
		LICENSE_LGPL20only,
		LICENSE_LGPL20orlater,
		LICENSE_LGPL21,
		LICENSE_LGPL21only,
		LICENSE_LGPL21PLUS,
		LICENSE_LGPL21orlater,
		LICENSE_LGPL30,
		LICENSE_LGPL30only,
		LICENSE_LGPL30orlater,
		LICENSE_LGPL30PLUS,
		LICENSE_LGPLLR,
		LICENSE_Libpng,
		LICENSE_libpng20,
		LICENSE_libselinux10,
		LICENSE_libtiff,
		LICENSE_LiLiQP11,
		LICENSE_LiLiQR11,
		LICENSE_LiLiQRplus11,
		LICENSE_Linuxmanpagescopyleft,
		LICENSE_LinuxOpenIB,
		LICENSE_LPL10,
		LICENSE_LPL102,
		LICENSE_LPPL10,
		LICENSE_LPPL11,
		LICENSE_LPPL12,
		LICENSE_LPPL13a,
		LICENSE_LPPL13c,
		LICENSE_MakeIndex,
		LICENSE_MirOS,
		LICENSE_MIT,
		LICENSE_MIT0,
		LICENSE_MITadvertising,
		LICENSE_MITCMU,
		LICENSE_MITenna,
		LICENSE_MITfeh,
		LICENSE_MITModernVariant,
		LICENSE_MITopengroup,
		LICENSE_MITNFA,
		LICENSE_Motosoto,
		LICENSE_mpich2,
		LICENSE_MPL10,
		LICENSE_MPL11,
		LICENSE_MPL20,
		LICENSE_MPL20nocopyleftexception,
		LICENSE_MSPL,
		LICENSE_MSRL,
		LICENSE_MTLL,
		LICENSE_MulanPSL10,
		LICENSE_MulanPSL20,
		LICENSE_Multics,
		LICENSE_Mup,
		LICENSE_NAIST2003,
		LICENSE_NASA13,
		LICENSE_Naumen,
		LICENSE_NBPL10,
		LICENSE_NCGLUK20,
		LICENSE_NCSA,
		LICENSE_NetSNMP,
		LICENSE_NetCDF,
		LICENSE_Newsletr,
		LICENSE_NGPL,
		LICENSE_NISTPD,
		LICENSE_NISTPDfallback,
		LICENSE_NLOD10,
		LICENSE_NLOD20,
		LICENSE_Nunit,
		LICENSE_NLPL,
		LICENSE_Nokia,
		LICENSE_NOSL,
		LICENSE_Noweb,
		LICENSE_NPL10,
		LICENSE_NPL11,
		LICENSE_NPOSL30,
		LICENSE_NRL,
		LICENSE_NTP,
		LICENSE_NTP0,
		LICENSE_OUDA10,
		LICENSE_OCCTPL,
		LICENSE_OCLC20,
		LICENSE_ODbL10,
		LICENSE_ODCBy10,
		LICENSE_OFL10,
		LICENSE_OFL10noRFN,
		LICENSE_OFL10RFN,
		LICENSE_OFL11,
		LICENSE_OFL11noRFN,
		LICENSE_OFL11RFN,
		LICENSE_OGC10,
		LICENSE_OGDLTaiwan10,
		LICENSE_OGLCanada20,
		LICENSE_OGLUK10,
		LICENSE_OGLUK20,
		LICENSE_OGLUK30,
		LICENSE_OGTSL,
		LICENSE_OLDAP11,
		LICENSE_OLDAP12,
		LICENSE_OLDAP13,
		LICENSE_OLDAP14,
		LICENSE_OLDAP20,
		LICENSE_OLDAP201,
		LICENSE_OLDAP21,
		LICENSE_OLDAP22,
		LICENSE_OLDAP221,
		LICENSE_OLDAP222,
		LICENSE_OLDAP23,
		LICENSE_OLDAP24,
		LICENSE_OLDAP25,
		LICENSE_OLDAP26,
		LICENSE_OLDAP27,
		LICENSE_OLDAP28,
		LICENSE_OML,
		LICENSE_OpenSSL,
		LICENSE_OPL10,
		LICENSE_OPUBL10,
		LICENSE_OSETPL21,
		LICENSE_OSL10,
		LICENSE_OSL11,
		LICENSE_OSL20,
		LICENSE_OSL21,
		LICENSE_OSL30,
		LICENSE_Parity600,
		LICENSE_Parity700,
		LICENSE_PDDL10,
		LICENSE_PHP30,
		LICENSE_PHP301,
		LICENSE_Plexus,
		LICENSE_PolyFormNoncommercial100,
		LICENSE_PolyFormSmallBusiness100,
		LICENSE_PostgreSQL,
		LICENSE_PSF20,
		LICENSE_psfrag,
		LICENSE_psutils,
		LICENSE_Python20,
		LICENSE_Qhull,
		LICENSE_QPL10,
		LICENSE_Rdisc,
		LICENSE_RHeCos11,
		LICENSE_RPL11,
		LICENSE_RPL15,
		LICENSE_RPSL10,
		LICENSE_RSAMD,
		LICENSE_RSCPL,
		LICENSE_Ruby,
		LICENSE_SAXPD,
		LICENSE_Saxpath,
		LICENSE_SCEA,
		LICENSE_SchemeReport,
		LICENSE_StandardMLNJ,
		LICENSE_Sendmail,
		LICENSE_Sendmail823,
		LICENSE_SGIB10,
		LICENSE_SGIB11,
		LICENSE_SGIB20,
		LICENSE_SHL05,
		LICENSE_SHL051,
		LICENSE_SimPL20,
		LICENSE_SISSL,
		LICENSE_SISSL12,
		LICENSE_Sleepycat,
		LICENSE_SMLNJ,
		LICENSE_SMPPL,
		LICENSE_SNIA,
		LICENSE_Spencer86,
		LICENSE_Spencer94,
		LICENSE_Spencer99,
		LICENSE_SPL10,
		LICENSE_SSHOpenSSH,
		LICENSE_SSHshort,
		LICENSE_SSPL10,
		LICENSE_SugarCRM113,
		LICENSE_SWL,
		LICENSE_TAPROHL10,
		LICENSE_TCL,
		LICENSE_TCPwrappers,
		LICENSE_TMate,
		LICENSE_TORQUE11,
		LICENSE_TOSL,
		LICENSE_TUBerlin10,
		LICENSE_TUBerlin20,
		LICENSE_UCL10,
		LICENSE_UnicodeDFS2015,
		LICENSE_UnicodeDFS2016,
		LICENSE_UnicodeTOU,
		LICENSE_Unlicense,
		LICENSE_UPL10,
		LICENSE_Vim,
		LICENSE_VOSTROM,
		LICENSE_VSL10,
		LICENSE_W3C,
		LICENSE_W3C19980720,
		LICENSE_W3C20150513,
		LICENSE_Watcom10,
		LICENSE_Wsuipa,
		LICENSE_WTFPL,
		LICENSE_X11,
		LICENSE_X11distributemodificationsvariant,
		LICENSE_wxWindows,
		LICENSE_Xerox,
		LICENSE_XFree8611,
		LICENSE_xinetd,
		LICENSE_Xnet,
		LICENSE_xpp,
		LICENSE_XSkat,
		LICENSE_YPL10,
		LICENSE_YPL11,
		LICENSE_Zed,
		LICENSE_Zend20,
		LICENSE_Zimbra13,
		LICENSE_Zimbra14,
		LICENSE_Zlib,
		LICENSE_zlibacknowledgement,
		LICENSE_ZPL11,
		LICENSE_ZPL20,
		LICENSE_ZPL21,
	}
	for _, license := range licenses {
		newDocument.License = MakeLicenses(license)
		_, err = Save(newDocument)
		if err != nil {
			t.Error("License is valid, however an error was thrown")
		}
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
