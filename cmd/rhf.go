package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "rhf",
	Short: descriptions.RhfShort,
	Long:  descriptions.RhfLong,
	Run:   handlers.RHFHandler,
}

func init() {
	rootCmd.AddCommand(genCmd)
}
