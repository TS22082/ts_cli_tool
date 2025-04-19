package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var rfcCmd = &cobra.Command{
	Use:   "rfc",
	Short: descriptions.RfcShort,
	Long:  descriptions.RfcLong,
	Run:   handlers.RFCHandler,
}

func init() {
	rootCmd.AddCommand(rfcCmd)
}
