package handlers

import (
	"fmt"
	"github.com/TS22082/ts_cli_tool/templates"
	"github.com/spf13/cobra"
	"os"
)

func RHFHandler(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("❌ Error: No function name provided, run with --help for more information")
		return
	}

	functionName := args[0]
	fileName := functionName + ".js"
	testFileName := functionName + ".test.js"

	fileContents := fmt.Sprintf(templates.Function, functionName, functionName, functionName)
	testFileContents := fmt.Sprintf(templates.FunctionTest, functionName, fileName, functionName)

	if err := os.WriteFile("index.js", []byte(fileContents), 0644); err != nil {
		fmt.Println("❌ Error creating function file")
		return
	}

	if err := os.WriteFile(testFileName, []byte(testFileContents), 0644); err != nil {
		fmt.Println("❌ Error creating test file:", err)
		return
	}

	fmt.Println("✅ File name:", fileName)
	fmt.Println("✅ Test file name:", testFileName)
	return
}
