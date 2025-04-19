package handlers

import (
	"fmt"
	"github.com/TS22082/ts_cli_tool/templates"
	"github.com/TS22082/ts_cli_tool/utils"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func RFCHandler(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		fmt.Println("❌ Error: No component name provided, run with --help for more information")
		return
	}

	componentName := args[0]
	presentationHookFileName := fmt.Sprintf("use%sData.js", componentName)
	storyFilename := fmt.Sprintf("%s.stories.jsx", componentName)
	testFileName := fmt.Sprintf("%s.test.js", componentName)

	componentContents := fmt.Sprintf(templates.Component, componentName, componentName, componentName, componentName, componentName, componentName)
	presentationHookContents := fmt.Sprintf(templates.Hook, componentName, componentName)
	storyContents := fmt.Sprintf(templates.Story, componentName, componentName, componentName, componentName)

	getPathCmd := exec.Command("pwd")
	pathCmdOutput, err := getPathCmd.Output()

	if err != nil {
		fmt.Println("❌ Cannot get path:", err)
		return
	}

	storiesDepth := utils.CountStoriesDepth(string(pathCmdOutput))
	testFileContents := fmt.Sprintf(templates.ComponentTest, utils.AddDots(storiesDepth), componentName, componentName, componentName, componentName)

	if err := os.WriteFile("index.jsx", []byte(componentContents), 0644); err != nil {
		fmt.Println("❌ Error creating component file:", err)
	}

	if err := os.WriteFile(presentationHookFileName, []byte(presentationHookContents), 0644); err != nil {
		fmt.Println("❌ Error creating presentation hook file:", err)
	}

	if err := os.WriteFile(storyFilename, []byte(storyContents), 0644); err != nil {
		fmt.Println("❌ Error creating component story file:", err)
	}

	if err := os.WriteFile(testFileName, []byte(testFileContents), 0644); err != nil {
		fmt.Println("❌ Error creating component test file:", err)
	}

	fmt.Println("✅ Component with story and presentation hook created")
}
