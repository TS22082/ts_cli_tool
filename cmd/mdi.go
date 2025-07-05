package cmd

import (
	"github.com/TS22082/ts_cli_tool/handlers"
	"github.com/spf13/cobra"
)

var mdiCmd = &cobra.Command{
	Use:   "mdi",
	Short: "mdi",
	Long:  `mdi`,
	Run:   handlers.MDIHandler,
}

func init() {
	rootCmd.AddCommand(mdiCmd)
}
