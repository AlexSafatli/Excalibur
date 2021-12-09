package main

import (
	"github.com/AlexSafatli/Excalibur/model"
	"github.com/AlexSafatli/Excalibur/template"
)

func main() {
	//cli.Execute()
	if err := template.WriteSheetToFile(&template.CharacterSheet{Title: "Hello", Character: &model.Character{}}, "test.html"); err != nil {
		panic(err)
	}
}
