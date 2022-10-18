package model

import (
	"encoding/json"
	"strconv"
)

type Implementation interface {
	Convert() *Character
}

type implementation struct {
	Implementation string `json:",omitempty"`
}

type SimpleImplementation struct {
	Title          string
	Implementation string
	Fields         []*Field
	Attributes     []*Attribute
	Skills         []*SimpleSkill
	Traits         []*SimpleTrait
	Equipment      []*SimpleEquipment
}

type SimpleSkill struct {
	Name   string
	Rating int
}

type SimpleTrait struct {
	Name        string
	Description string
}

type SimpleEquipment struct {
	Name     string
	Quantity int
}

func (i *SimpleImplementation) Convert() *Character {
	var c = &Character{}
	c.Fields = i.Fields
	c.Attributes = i.Attributes
	for _, s := range i.Skills {
		c.Skills = append(c.Skills, &Skill{
			Name:   s.Name,
			Fields: []*Field{{"Rating", strconv.Itoa(s.Rating)}},
		})
	}
	for _, t := range i.Traits {
		c.Traits = append(c.Traits, &Trait{
			Name:      t.Name,
			ValueName: "Description",
			Value:     t.Description,
		})
	}
	for _, e := range i.Equipment {
		c.Equipment = append(c.Equipment, &Item{
			Name:   e.Name,
			Fields: []*Field{{"Quantity", strconv.Itoa(e.Quantity)}},
		})
	}
	return c
}

func read(dat []byte, implementation string) Implementation {
	switch implementation {
	case "simple":
		var i SimpleImplementation
		if err := json.Unmarshal(dat, &i); err != nil {
			panic(err)
		}
		return &i
	}
	return nil
}

func convertToCharacter(dat []byte, implementation string) *Character {
	return read(dat, implementation).Convert()
}

func ImportCharacterFromJSON(dat []byte) *Character {
	var i implementation
	err := json.Unmarshal(dat, &i)
	if err != nil {
		panic(err)
	}
	return convertToCharacter(dat, i.Implementation)
}