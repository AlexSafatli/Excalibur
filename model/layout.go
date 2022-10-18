package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	LayoutsDir = "layouts"
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

func ImportLayout(p string) *Layout {
	dat, err := ioutil.ReadFile(p)
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

func ImportRelativeLayout(p string) *Layout {
	return ImportLayout(fmt.Sprintf("%s/%s.json", LayoutsDir, p))
}
