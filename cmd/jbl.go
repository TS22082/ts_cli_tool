package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var jiraBlCmd = &cobra.Command{
	Use:   "jbl",
	Short: descriptions.JblShort,
	Long:  descriptions.JblLong,
	Run:   handlers.JBLHandler,
}

func init() {
	rootCmd.AddCommand(jiraBlCmd)
	jiraBlCmd.Flags().StringP("user", "u", "", "Backlog for specific user")
}
