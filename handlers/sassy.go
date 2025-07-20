package handlers

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func SassyHandler(cmd *cobra.Command, args []string) {

	srcDir := "./static/src/css/styles.scss"
	distDir := "./static/dist/css/style.css"

	sassCompileCmd := exec.Command("sass", srcDir, distDir)

	_, err := sassCompileCmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Src Dir:", srcDir)
		fmt.Println("Dist Dir:", distDir)
		fmt.Println(strings.Join(sassCompileCmd.Args, " "))

		return
	}

	fmt.Println("Done")
}
