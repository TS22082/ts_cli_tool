package handlers

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func SassyHandler(cmd *cobra.Command, args []string) {

	sassCompileCmd := exec.Command("sass", "static/src/css/style.scss", "static/dist/css/style.css")

	_, err := sassCompileCmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Done")
}
