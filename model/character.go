package model

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
