package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var gclCmd = &cobra.Command{
	Use:   "gcl",
	Short: descriptions.GclShort,
	Long:  descriptions.GclLong,
	Run:   handlers.GCLHandler,
}

func init() {
	rootCmd.AddCommand(gclCmd)
}
