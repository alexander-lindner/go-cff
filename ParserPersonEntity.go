package cff

import (
	"gopkg.in/yaml.v3"
)

//EmptyPerson represents an empty person
var EmptyPerson Person

//Get is used to retrieve a person and an entity
func (t *PersonEntity) Get() (Person, Entity) {
	return t.Person, t.Entity
}

//MarshalYAML is used to serialise a PersonEntity
func (t PersonEntity) MarshalYAML() (interface{}, error) {
	if t.IsPerson || (t.Person != EmptyPerson) {
		return t.Person, nil
	} else {
		return t.Entity, nil
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
