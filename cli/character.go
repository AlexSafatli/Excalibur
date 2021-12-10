package cli

import (
	"errors"
	"fmt"
	"github.com/AlexSafatli/Excalibur/model"
	"github.com/AlexSafatli/Excalibur/template"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var createCharacterCmd = &cobra.Command{
	Use:   "create <name> <player> <path>",
	Short: "Create a new empty character and write as a JSON",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("need a character and player name and a JSON path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		c := model.NewEmptyCharacter(args[0], args[1])
		if err := model.WriteCharacterToJSON(c, args[2]); err != nil {
			panic(err)
		}
		fmt.Printf("Generated new character (Name: '%s') and saved to %s", args[0], args[2])
	},
}

var writeCharacterCmd = &cobra.Command{
	Use:   "write <json_path> <html_path>",
	Short: "Write an existing character JSON to a printable HTML file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("need a JSON and output HTML path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dat, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		c := model.ImportCharacterFromJSON(dat)
		if err := template.WriteSheetToFile(&template.CharacterSheet{Title: "Character Sheet", Character: c}, args[1]); err != nil {
			panic(err)
		}
		fmt.Printf("Wrote %s to %s", args[0], args[1])
	},
}
