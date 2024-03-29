package sheet

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	ImplementationSimple = "simple"
)

type Implementation interface {
	Convert() *Character
}

type implementation struct {
	Implementation string `json:",omitempty"`
}

type SimpleImplementation struct {
	headers
	Fields     []*Field
	Attributes []*Attribute
	Skills     []*SimpleSkill
	Abilities  []*SimpleTrait
	Traits     []*SimpleTrait
	Equipment  []*SimpleEquipment
	Vehicles   []*SimpleEquipment
}

type SimpleSkill struct {
	Name string
	Rank int
}

type SimpleTrait struct {
	Name   string
	Effect string
}

type SimpleEquipment struct {
	Name   string
	Effect string
	Size   string
}

func (i *SimpleImplementation) Convert() *Character {
	var c = &Character{}
	c.CampaignName = i.CampaignName
	c.Fields = i.Fields
	c.Attributes = i.Attributes
	for _, s := range i.Skills {
		c.Skills = append(c.Skills, &Skill{
			Name:   s.Name,
			Fields: []*Field{{"Rank", strconv.Itoa(s.Rank)}},
		})
	}
	for _, t := range i.Abilities {
		c.Abilities = append(c.Abilities, &Trait{Name: t.Name, Value: t.Effect,
			ValueName: "Effect"})
	}
	for _, t := range i.Traits {
		c.Traits = append(c.Traits, &Trait{Name: t.Name, Value: t.Effect,
			ValueName: "Effect"})
	}
	for _, e := range i.Equipment {
		item := &Item{
			Name:   e.Name,
			Fields: []*Field{{"Effect", e.Effect}},
		}
		if e.Size != "" {
			item.Fields = append(item.Fields, &Field{"Size", e.Size})
		}
		c.Equipment = append(c.Equipment, item)
	}
	for _, e := range i.Vehicles {
		item := &Item{
			Name:   e.Name,
			Fields: []*Field{{"Effect", e.Effect}},
		}
		if e.Size != "" {
			item.Fields = append(item.Fields, &Field{"Size", e.Size})
		}
		c.Vehicles = append(c.Vehicles, item)
	}
	return c
}

func read(dat []byte, implementation string) Implementation {
	switch implementation {
	case ImplementationSimple:
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
		fmt.Println(err)
		os.Exit(1)
	}
	return convertToCharacter(dat, i.Implementation)
}
