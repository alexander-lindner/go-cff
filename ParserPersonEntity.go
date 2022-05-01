package main

import (
	"gopkg.in/yaml.v3"
)

var EmptyPerson Person

func (t *PersonEntity) Get() (Person, Entity) {
	return t.Person, t.Entity
}

func (t PersonEntity) MarshalYAML() (interface{}, error) {
	if t.IsPerson || (t.Person != EmptyPerson) {
		return t.Person, nil
	} else {
		return t.Entity, nil
	}
}
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
