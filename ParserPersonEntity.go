package cff

import (
	"errors"
	"gopkg.in/yaml.v3"
)

//EmptyPerson represents an empty Person
var EmptyPerson Person

//EmptyEntity represents an empty Entity
var EmptyEntity Entity

//Get is used to retrieve a person and an entity
func (t *PersonEntity) Get() (Person, Entity) {
	return t.Person, t.Entity
}

//MarshalYAML is used to serialise a PersonEntity
func (t PersonEntity) MarshalYAML() (interface{}, error) {
	t.IsPerson = t.Person != EmptyPerson
	t.IsEntity = t.Entity != EmptyEntity
	if (t.IsPerson) && (!t.IsEntity) {
		return t.Person, nil
	} else if (t.IsEntity) && (!t.IsPerson) {
		return t.Entity, nil
	} else {
		return nil, errors.New("identifier is empty or both person and entity were set")
	}
}

//UnmarshalYAML creates a PersonEntity from a cff Person or Entity
func (t *PersonEntity) UnmarshalYAML(v *yaml.Node) error {
	var person Person

	isEntity := false
	for _, node := range v.Content {
		if node.Value == "address" {
			isEntity = true
		}
	}
	t.IsEntity = isEntity
	t.IsPerson = !isEntity

	if !isEntity {
		err := v.Decode(&person)
		if err != nil {
			return err
		}
		t.Person = person
		return nil
	}
	var entity Entity
	err := v.Decode(&entity)
	if err == nil {
		t.Entity = entity
	}

	return err
}
