package handlers

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"regexp"
)

func JTPHandler(cmd *cobra.Command, args []string) {
	getBranchNameCmd := exec.Command("git", "symbolic-ref", "--short", "-q", "HEAD")
	branchName, err := getBranchNameCmd.Output()

	if err != nil {
		fmt.Println("Error getting branch name")
		return
	}

	pattern := `ES-[0-9]+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(branchName), -1)

	if len(matches) == 0 {
		fmt.Println("❌ Branch name does not match ES-<>")
		return
	}

	jiraTicketNumber := matches[0][0]
	openJiraTicketCmd := exec.Command("open", "https://estatespace.atlassian.net/browse/"+jiraTicketNumber)

	if err := openJiraTicketCmd.Run(); err != nil {
		fmt.Println("❌ Error opening Jira ticket")
		return
	}

	fmt.Println("✅ Opened Jira ticket in browser")
}
