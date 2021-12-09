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

var templatePath = "template/sheet.html"

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

func WriteSheetToFile(c *CharacterSheet, name string) error {
	return render(c, name)
}
