package handlers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func SassyHandler(cmd *cobra.Command, args []string) {

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	srcDir := dir + "/static/src/css/style.scss"
	distDir := dir + "/static/dist/css/style.css"

	sassCompileCmd := exec.Command("sass", srcDir, distDir)

	_, err = sassCompileCmd.Output()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Src Dir:", srcDir)
		fmt.Println("Dist Dir:", distDir)
		return
	}

	fmt.Println("Done")
}
