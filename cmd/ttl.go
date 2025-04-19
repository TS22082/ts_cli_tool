package cmd

import (
	"github.com/TS22082/ts_cli_tool/descriptions"
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "ttl",
	Short: descriptions.TtlShort,
	Long:  descriptions.TtlLong,
	Run:   handlers.TTLHandler,
}

func init() {
	rootCmd.AddCommand(ttlCmd)
}
