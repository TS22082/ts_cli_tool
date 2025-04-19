package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var gacCmd = &cobra.Command{
	Use:   "gac",
	Short: descriptions.GacShort,
	Long:  descriptions.GacLong,
	Run:   handlers.GACHandler,
}

func init() {
	rootCmd.AddCommand(gacCmd)
}
