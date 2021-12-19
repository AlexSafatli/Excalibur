package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Layout struct {
	Title string
	Defaults
}

type Defaults struct {
	Fields     []*Field
	Attributes []*Attribute
	Skills     []*Skill
	Traits     []*Trait
	Equipment  []*Item
}

func ImportLayout(name string) *Layout {
	dat, err := ioutil.ReadFile(fmt.Sprintf("systems/%s.json", name))
	if err != nil {
		panic(err)
	}

	var l Layout
	err = json.Unmarshal(dat, &l)
	if err != nil {
		panic(err)
	}
	return &l
}
