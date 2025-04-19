package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var jiraCmd = &cobra.Command{
	Use:   "jtp",
	Short: descriptions.JtpShort,
	Long:  descriptions.JtpLong,
	Run:   handlers.JTPHandler,
}

func init() {
	rootCmd.AddCommand(jiraCmd)
}
