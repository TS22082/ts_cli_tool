package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "es",
	Short: "Geekcentric cli",
	Long: `A CLI tool for myself

Use es --help for more information, or visit https://github.com/TS22082/ts_cli_tool
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
