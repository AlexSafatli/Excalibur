package cli

import (
	"errors"
	"fmt"
	"github.com/AlexSafatli/Excalibur/sheet"
	"github.com/AlexSafatli/Excalibur/template"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var createCharacterCmd = &cobra.Command{
	Use:   "create <name> <player> <path>",
	Short: "Create a new empty character and write as a JSON",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("need a character, player, and a JSON path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		c := sheet.NewEmptyCharacter(args[0], args[1])
		if err := sheet.WriteCharacterToJSON(c, args[2]); err != nil {
			panic(err)
		}
		fmt.Printf("Generated new character ('%s') and saved to '%s'",
			args[0], args[2])
	},
}

var writeCharacterCmd = &cobra.Command{
	Use:   "write <json_path>...",
	Short: "Write any number of existing character JSONs to other formats",
	Run: func(cmd *cobra.Command, args []string) {
		var libs = []*sheet.Layout{sheet.ImportRelativeLayout("base")}
		if len(layouts) > 0 {
			for _, ll := range strings.Split(layouts, ",") {
				var lib *sheet.Layout
				if sheet.IsRelativeLayout(ll) {
					lib = sheet.ImportRelativeLayout(ll)
				} else {
					lib = sheet.ImportLayout(ll)
				}
				libs = append(libs, lib)
			}
		}
		for _, arg := range args {
			dat, err := ioutil.ReadFile(arg)
			if err != nil {
				panic(err)
			}
			c := sheet.ImportCharacterFromJSON(dat)
			n := strings.TrimSuffix(arg, filepath.Ext(arg))
			p := fmt.Sprintf("%s.%s", n, sheet.FormatHtml)
			if err := template.WriteSheetToFile(
				&template.CharacterSheet{Title: "Character Sheet",
					Character: c},
				p, libs...); err != nil {
				panic(err)
			}
			fmt.Printf("Wrote '%s' to '%s'\n", arg, p)
		}
	},
}

func init() {
	writeCharacterCmd.Flags().StringVar(&layouts, "layouts", "", "a comma-separated list of defined libraries to draw from for skills, traits, etc. when generating and rendering")
	writeCharacterCmd.Flags().BoolVar(&h, "html", true, "output to printable HTML")
}
