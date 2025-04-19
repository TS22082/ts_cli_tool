package handlers

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var userJiraIdMap = map[string]string{
	"thomas":    "712020%3Ab256ae7c-e78c-44e5-9341-47ada7036c08",
	"raag":      "5b1aec7a9595132534bc38e9",
	"chris":     "600819e558d5a20069227b7c",
	"parker":    "6260af987be65e006935088b",
	"christian": "62326fce72d9010069e3df2c",
	"sylvia":    "6260af990f5cf500697f6d85",
	"josymar":   "62a86b2c188d08006fe157ac",
	"mario":     "712020%3Af97c69e4-fdf4-4023-acd1-d8d331b5dc9f",
	"konan":     "62543b5fed2b3e0074fbb7c4",
	"kate":      "712020%3A3d43df3b-01a0-4469-88fe-6350e9f747b8",
}

func JBLHandler(cmd *cobra.Command, args []string) {
	user := cmd.Flag("user").Value.String()

	if userJiraIdMap[user] == "" {
		fmt.Println("❌ User not found")
		return
	}

	openJiraCmd := exec.Command("open", "https://estatespace.atlassian.net/jira/software/projects/ES/boards/26/backlog?assignee="+userJiraIdMap[user])

	if err := openJiraCmd.Run(); err != nil {
		fmt.Println("❌ Error opening Jira ticket")
		return
	}

	fmt.Println("✅ Opened Jira ticket in browser")
}
