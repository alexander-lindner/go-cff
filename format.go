//Package cff provides a parser for the CFF format.
//Go to https://alexander-lindner.github.io/go-cff/ for a full documentation.
package cff

import (
	"net/url"
	"time"
)

//URL represents a basic url. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsurl for more information.
type URL struct {
	*url.URL
}

//Date represents a date
type Date time.Time

//DOI represents a CFF DOI. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsdoi for more information.
type DOI struct {
	General            string
	DirectoryIndicator string
	RegistrantCode     string
}

//Entity represents an CFF entity. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsentity for more information.
type Entity struct {
	Address   string `yaml:"address,omitempty"`
	Alias     string `yaml:"alias,omitempty"`
	City      string `yaml:"city,omitempty"`
	Country   string `yaml:"country,omitempty"`
	DateEnd   Date   `yaml:"date-end,omitempty"`
	DateStart Date   `yaml:"date-start,omitempty"`
	Email     string `yaml:"email,omitempty"`
	Fax       string `yaml:"fax,omitempty"`
	Location  string `yaml:"location,omitempty"`
	Name      string `yaml:"name,omitempty"`
	Orcid     string `yaml:"orcid,omitempty"`
	PostCode  int    `yaml:"post-code,omitempty"`
	Region    string `yaml:"region,omitempty"`
	Tel       string `yaml:"tel,omitempty"`
	Website   URL    `yaml:"website,omitempty"`
}

//Person represents an CFF person. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsperson for more information.
type Person struct {
	Address      string `yaml:"address,omitempty"`
	Affiliation  string `yaml:"affiliation,omitempty"`
	Alias        string `yaml:"alias,omitempty"`
	City         string `yaml:"city,omitempty"`
	Country      string `yaml:"country,omitempty"`
	Email        string `yaml:"email,omitempty"`
	Family       string `yaml:"family-names,omitempty"`
	Fax          string `yaml:"fax,omitempty"`
	GivenNames   string `yaml:"given-names,omitempty"`
	NameParticle string `yaml:"name-particle,omitempty"`
	NameSuffix   string `yaml:"name-suffix,omitempty"`
	Orcid        string `yaml:"orcid,omitempty"`
	PostCode     int    `yaml:"post-code,omitempty"`
	Region       string `yaml:"region,omitempty"`
	Tel          string `yaml:"tel,omitempty"`
	Website      URL    `yaml:"website,omitempty"`
}

//PersonEntity a proxy struct which holds either a Person or an Entity.
//Use the bool field to determine which one is set.
type PersonEntity struct {
	Person   Person
	Entity   Entity
	IsEntity bool
	IsPerson bool
}

//Identifier represents an CFF identifier. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsidentifier for more information.
type Identifier struct {
	IdentifiersType string `yaml:"type,omitempty"`
	Value           string `yaml:"value,omitempty"`
	Description     string `yaml:"description,omitempty"`
}

//Reference represents an CFF references. See https://github.com/citation-file-format/citation-file-format/blob/main/schema-guide.md#definitionsreference for more information.
type Reference struct {
	Abbreviation       string         `yaml:"abbreviation,omitempty"`
	Abstract           string         `yaml:"abstract,omitempty"`
	Authors            []PersonEntity `yaml:"authors,omitempty"`
	CollectionDoi      DOI            `yaml:"collection-doi,omitempty"`
	CollectionTitle    string         `yaml:"collection-title,omitempty"`
	CollectionType     string         `yaml:"collection-type,omitempty"`
	Commit             string         `yaml:"commit,omitempty"`
	Conference         Entity         `yaml:"conference,omitempty"`
	Contact            []PersonEntity `yaml:"contact,omitempty"`
	Copyright          string         `yaml:"copyright,omitempty"`
	DataType           string         `yaml:"data-type,omitempty"`
	DatabaseProvider   Entity         `yaml:"database-provider,omitempty"`
	Database           string         `yaml:"database,omitempty"`
	DateAccessed       Date           `yaml:"date-accessed,omitempty"`
	DateDownloaded     Date           `yaml:"date-downloaded,omitempty"`
	DatePublished      Date           `yaml:"date-published,omitempty"`
	DateReleased       Date           `yaml:"date-released,omitempty"`
	Department         string         `yaml:"department,omitempty"`
	Doi                DOI            `yaml:"doi,omitempty"`
	Edition            string         `yaml:"edition,omitempty"`
	Editors            []PersonEntity `yaml:"editors,omitempty"`
	EditorsSeries      []PersonEntity `yaml:"editors-series,omitempty"`
	End                string         `yaml:"end,omitempty"`
	Entry              string         `yaml:"entry,omitempty"`
	Filename           string         `yaml:"filename,omitempty"`
	Format             string         `yaml:"format,omitempty"`
	Identifiers        []Identifier   `yaml:"identifiers,omitempty"`
	Institution        Entity         `yaml:"institution,omitempty"`
	Isbn               string         `yaml:"isbn,omitempty"`
	Issn               string         `yaml:"issn,omitempty"`
	Issue              string         `yaml:"issue,omitempty"`
	IssueDate          string         `yaml:"issue-date,omitempty"`
	IssueTitle         string         `yaml:"issue-title,omitempty"`
	Journal            string         `yaml:"journal,omitempty"`
	Keywords           string         `yaml:"keywords,omitempty"`
	Languages          string         `yaml:"languages,omitempty"`
	License            string         `yaml:"license,omitempty"`
	LicenseUrl         URL            `yaml:"license-url,omitempty"`
	LocEnd             string         `yaml:"loc-end,omitempty"`
	LocStart           string         `yaml:"loc-start,omitempty"`
	Location           Entity         `yaml:"location,omitempty"`
	Medium             string         `yaml:"medium,omitempty"`
	Month              int            `yaml:"month,omitempty"`
	Nihmsid            string         `yaml:"nihmsid,omitempty"`
	Notes              string         `yaml:"notes,omitempty"`
	Number             string         `yaml:"number,omitempty"`
	NumberVolumes      string         `yaml:"number-volumes,omitempty"`
	Pages              string         `yaml:"pages,omitempty"`
	PatentStates       string         `yaml:"patent-states,omitempty"`
	Pmcid              string         `yaml:"pmcid,omitempty"`
	Publisher          Entity         `yaml:"publisher,omitempty"`
	Recipients         PersonEntity   `yaml:"recipients,omitempty"`
	Repository         URL            `yaml:"repository,omitempty"`
	RepositoryArtifact URL            `yaml:"repository-artifact,omitempty"`
	RepositoryCode     URL            `yaml:"repository-code,omitempty"`
	Scope              string         `yaml:"scope,omitempty"`
	Section            string         `yaml:"section,omitempty"`
	Senders            PersonEntity   `yaml:"senders,omitempty"`
	Start              string         `yaml:"start,omitempty"`
	Status             string         `yaml:"status,omitempty"`
	Term               string         `yaml:"term,omitempty"`
	ThesisType         string         `yaml:"thesis-type,omitempty"`
	Title              string         `yaml:"title,omitempty"`
	Translators        PersonEntity   `yaml:"translators,omitempty"`
	ReferenceType      string         `yaml:"type,omitempty"`
	Url                URL            `yaml:"url,omitempty"`
	Version            string         `yaml:"version,omitempty"`
	Volume             int            `yaml:"volume,omitempty"`
	VolumeTitle        string         `yaml:"volume-title,omitempty"`
	Year               int            `yaml:"year,omitempty"`
	YearOriginal       int            `yaml:"year-original,omitempty"`
}

//Cff is the representation of a CFF file
type Cff struct {
	Abstract           string         `yaml:"abstract,omitempty"`
	Authors            []PersonEntity `yaml:"authors,omitempty"`
	CffVersion         string         `yaml:"cff-version,omitempty"`
	Commit             string         `yaml:"commit,omitempty"`
	Contact            []PersonEntity `yaml:"contact,omitempty"`
	DateReleased       Date           `yaml:"date-released,omitempty"`
	Doi                DOI            `yaml:"doi,omitempty"`
	Identifiers        []Identifier   `yaml:"identifiers,omitempty"`
	Keywords           []string       `yaml:"keywords,omitempty"`
	License            string         `yaml:"license,omitempty"` //todo validator using custom type
	LicenseUtl         URL            `yaml:"license-url,omitempty"`
	Message            string         `yaml:"message,omitempty"`
	PreferredCitation  Reference      `yaml:"preferred-citation,omitempty"`
	References         []Reference    `yaml:"references,omitempty"`
	Repository         URL            `yaml:"repository,omitempty"`
	RepositoryArtifact URL            `yaml:"repository-artifact,omitempty"`
	RepositoryCode     URL            `yaml:"repository-code,omitempty"`
	Title              string         `yaml:"title,omitempty"`
	CffType            string         `yaml:"type,omitempty"` //todo validator using custom type
	Url                URL            `yaml:"url,omitempty"`
	Version            string         `yaml:"version,omitempty"`
}
