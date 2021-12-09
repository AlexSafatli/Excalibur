package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var createCharacterCmd = &cobra.Command{
	Use:   "character <name>",
	Short: "Create a new character and write as a JSON",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("need only a name for the character JSON file")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generated new character and saved to %s.json\n", args[0])
	},
}
