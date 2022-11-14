package sheet

import (
	"encoding/json"
	"io/ioutil"
)

type Character struct {
	headers
	characteristics
}

type headers struct {
	Title          string
	CampaignName   string
	Implementation string
}

type characteristics struct {
	Fields     []*Field
	Attributes []*Attribute
	Skills     []*Skill
	Abilities  []*Trait
	Traits     []*Trait
	Equipment  []*Item
	Vehicles   []*Item
}

type Field struct {
	Name  string
	Value string
}

type Attribute struct {
	Name      string
	ValueName string
	Value     int
	Color     string
}

func NewEmptyCharacter(name, player string) *Character {
	return &Character{
		characteristics: characteristics{Fields: []*Field{
			{
				Name:  "Name",
				Value: name,
			},
			{
				Name:  "Player",
				Value: player,
			},
		}},
	}
}

func WriteCharacterToJSON(c *Character, path string) error {
	dat, _ := json.MarshalIndent(c, "", " ")
	return ioutil.WriteFile(path, dat, 0644)
}

func (c *Character) HasField(f *Field) bool {
	for _, field := range c.Fields {
		if field.Name == f.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetField(f *Field) *Field {
	for _, field := range c.Fields {
		if field.Name == f.Name {
			return field
		}
	}
	return nil
}

func (c *Character) Name() string {
	f := c.GetField(&Field{Name: "Name"})
	if f == nil {
		return ""
	}
	return f.Value
}

func (c *Character) HasAttribute(a *Attribute) bool {
	for _, attr := range c.Attributes {
		if attr.Name == a.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetAttribute(a *Attribute) *Attribute {
	for _, attr := range c.Attributes {
		if attr.Name == a.Name {
			return attr
		}
	}
	return nil
}

func (c *Character) HasSkill(s *Skill) bool {
	for _, skill := range c.Skills {
		if skill.Name == s.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetSkill(s *Skill) *Skill {
	for _, skill := range c.Skills {
		if skill.Name == s.Name {
			return skill
		}
	}
	return nil
}

func (c *Character) HasAbility(t *Trait) bool {
	for _, trait := range c.Abilities {
		if trait.Name == t.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetAbility(t *Trait) *Trait {
	for _, trait := range c.Abilities {
		if trait.Name == t.Name {
			return trait
		}
	}
	return nil
}

func (c *Character) HasTrait(t *Trait) bool {
	for _, trait := range c.Traits {
		if trait.Name == t.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetTrait(t *Trait) *Trait {
	for _, trait := range c.Traits {
		if trait.Name == t.Name {
			return trait
		}
	}
	return nil
}

func (c *Character) HasItem(i *Item) bool {
	for _, item := range c.Equipment {
		if item.Name == i.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetItem(i *Item) *Item {
	for _, item := range c.Equipment {
		if item.Name == i.Name {
			return item
		}
	}
	return nil
}

func (c *Character) HasVehicle(i *Item) bool {
	for _, item := range c.Vehicles {
		if item.Name == i.Name {
			return true
		}
	}
	return false
}

func (c *Character) GetVehicle(i *Item) *Item {
	for _, item := range c.Vehicles {
		if item.Name == i.Name {
			return item
		}
	}
	return nil
}
