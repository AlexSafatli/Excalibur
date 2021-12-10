package model

import (
	"encoding/json"
	"io/ioutil"
)

type Character struct {
	Fields     []*Field
	Attributes []*Attribute
	Skills     []*Skill
	Traits     []*Trait
	Equipment  []*Item
}

type Field struct {
	Name  string
	Value string
}

func NewEmptyCharacter(name, player string) *Character {
	return &Character{
		Fields: []*Field{
			{
				Name:  "Name",
				Value: name,
			},
			{
				Name:  "Player",
				Value: player,
			},
		},
	}
}

func ImportCharacterFromJSON(dat []byte) *Character {
	var c Character
	err := json.Unmarshal(dat, &c)
	if err != nil {
		return &Character{}
	}
	return &c
}

func WriteCharacterToJSON(c *Character, path string) error {
	dat, _ := json.MarshalIndent(c, "", " ")
	return ioutil.WriteFile(path, dat, 0644)
}
