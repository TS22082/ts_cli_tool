package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var sassyCmd = &cobra.Command{
	Use:   "sassy <path> <path>",
	Short: descriptions.SassyShort,
	Long:  descriptions.SassyLong,
	Run:   handlers.SassyHandler,
}

func init() {
	rootCmd.AddCommand(sassyCmd)
}
