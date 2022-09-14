package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/apicli/apicli/internal/pkg/ui"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apicli",
	Short: "Apicli is a CLI for interacting with API's",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ui.Start()
		defer ui.Stop()
		if err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
