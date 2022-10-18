package template

import (
	"github.com/AlexSafatli/Excalibur/model"
	"html/template"
	"os"
)

type CharacterSheet struct {
	Title     string
	Character *model.Character
}

const templatePath = "template/sheet.html"

func render(c *CharacterSheet, name string) error {
	t := template.New("sheet.html")
	t.Funcs(template.FuncMap{
		"mod": func(i, j int) bool { return i%j == 0 },
		"skillColumns": func(c *CharacterSheet) []*model.Field {
			if len(c.Character.Skills) == 0 {
				return []*model.Field{}
			}
			return c.Character.Skills[0].Fields
		},
	})
	t, err := t.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	err = t.Execute(f, c)
	if err != nil {
		return err
	}
	return nil
}

func addDefaults(c *CharacterSheet, l *model.Layout) {
	for _, field := range l.Fields {
		if !c.Character.HasField(field) {
			c.Character.Fields = append(c.Character.Fields, field)
		}
	}

	for _, attr := range l.Attributes {
		if attr.Name == "*" {
			// Wildcard
			for i := range c.Character.Attributes {
				if len(attr.ValueName) > 0 {
					if c.Character.Attributes[i].ValueName == "" {
						c.Character.Attributes[i].ValueName = attr.ValueName
					}
				}
				if attr.Value != 0 {
					if c.Character.Attributes[i].Value == 0 {
						c.Character.Attributes[i].Value = attr.Value
					}
				}
				if len(attr.Color) > 0 {
					if c.Character.Attributes[i].Color == "" {
						c.Character.Attributes[i].Color = attr.Color
					}
				}
			}
		} else {
			if !c.Character.HasAttribute(attr) {
				c.Character.Attributes = append(c.Character.Attributes, attr)
			} else {
				a := c.Character.GetAttribute(attr)
				if len(attr.ValueName) > 0 {
					a.ValueName = attr.ValueName
				}
				if attr.Value != 0 {
					a.Value = attr.Value
				}
				if len(attr.Color) > 0 {
					a.Color = attr.Color
				}
			}
		}
	}

	for _, skill := range l.Skills {
		if skill.Name == "*" {
			// Wildcard
			for i := range c.Character.Skills {
				for _, f := range skill.Fields {
					if !c.Character.Skills[i].HasField(f) {
						c.Character.Skills[i].Fields = append(c.Character.Skills[i].Fields, f)
					}
				}
			}
		} else {
			if !c.Character.HasSkill(skill) {
				c.Character.Skills = append(c.Character.Skills, skill)
			} else {
				s := c.Character.GetSkill(skill)
				for _, f := range skill.Fields {
					if !s.HasField(f) {
						s.Fields = append(s.Fields, f)
					}
				}
			}
		}
	}

	for _, trait := range l.Traits {
		if trait.Name == "*" {
			// Wildcard
			for i := range c.Character.Traits {
				if len(trait.ValueName) > 0 {
					c.Character.Traits[i].ValueName = trait.ValueName
				}
				if len(trait.Value) > 0 {
					c.Character.Traits[i].Value = trait.Value
				}
			}
		} else {
			if !c.Character.HasTrait(trait) {
				c.Character.Traits = append(c.Character.Traits, trait)
			} else {
				t := c.Character.GetTrait(trait)
				if len(trait.ValueName) > 0 {
					t.ValueName = trait.ValueName
				}
				if len(trait.Value) > 0 {
					t.Value = trait.Value
				}
			}
		}
	}

	for _, item := range l.Equipment {
		if item.Name == "*" {
			// Wildcard
			for i := range c.Character.Equipment {
				for _, f := range item.Fields {
					if !c.Character.Equipment[i].HasField(f) {
						c.Character.Equipment[i].Fields = append(c.Character.Equipment[i].Fields, f)
					}
				}
			}
		} else {
			if !c.Character.HasItem(item) {
				c.Character.Equipment = append(c.Character.Equipment, item)
			} else {
				i := c.Character.GetItem(item)
				for _, f := range item.Fields {
					if !i.HasField(f) {
						i.Fields = append(i.Fields, f)
					}
				}
			}
		}
	}
}

func WriteSheetToFile(c *CharacterSheet, name string, l ...*model.Layout) error {
	for _, ll := range l {
		addDefaults(c, ll)
	}
	return render(c, name)
}
