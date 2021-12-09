package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "excalibur",
	Short: "Excalibur is a CLI to manage and generate printable HTML character sheets from input JSONs for a variety of TTRPG systems",
	Args:  needCommandArg,
	Run:   noOpCmd,
}

func needCommandArg(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("need a command to run")
	}
	return nil
}

func noOpCmd(_ *cobra.Command, _ []string) {

}

func Execute() {
	// root
	rootCmd.AddCommand(createCharacterCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
