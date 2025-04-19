package handlers

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func GCLHandler(cmd *cobra.Command, args []string) {
	getBranchesCmd := exec.Command("git", "branch", "--list", "ES-*")

	branches, err := getBranchesCmd.Output()

	if err != nil {
		fmt.Println("❌ Error getting local branches", err.Error())
		return
	}

	numberOfBranches := strings.Count(string(branches), "\n")

	if numberOfBranches == 0 {
		fmt.Println("❌ No matching local branches to delete")
		return
	}

	for _, branch := range strings.Split(string(branches), "\n") {

		branchString := strings.TrimSpace(string(branch))

		if len(branchString) == 0 {
			continue
		}

		fmt.Println("Deleting branch", branch)
		deleteBranchCmd := exec.Command("git", "branch", "-D", branchString)
		if err := deleteBranchCmd.Run(); err != nil {
			fmt.Println("❌ Error deleting branch", err.Error())
			return
		}
	}

	fmt.Println("✅ Complete: Deleted", numberOfBranches, "local branches")
	return
}
