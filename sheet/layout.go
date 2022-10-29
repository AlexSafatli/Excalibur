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

func ImportLayout(p string) *Character {
	dat, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	var l Character
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

func ImportRelativeLayout(p string) *Character {
	return ImportLayout(fmt.Sprintf("%s/%s.json", LayoutsDir, p))
}
