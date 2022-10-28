package sheet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

func IsRelativeLayout(p string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s/%s.json",
		LayoutsDir, p)); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func ImportRelativeLayout(p string) *Layout {
	return ImportLayout(fmt.Sprintf("%s/%s.json", LayoutsDir, p))
}
