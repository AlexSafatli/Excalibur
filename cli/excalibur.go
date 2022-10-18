package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	layouts string
)

var rootCmd = &cobra.Command{
	Use:   "excalibur",
	Short: "Excalibur is a CLI to manage and generate printable HTML character sheets from input JSONs for the VOID TTRPG system",
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
	rootCmd.AddCommand(writeCharacterCmd)
	rootCmd.Flags().StringVar(&layouts, "layouts", "", "a comma-separated list of defined libraries to draw from for skills, traits, etc. when generating and rendering")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
