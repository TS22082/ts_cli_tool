package utils

import "strings"

func CleanAIResponse(s string) string {
	s = strings.TrimSpace(s)
	// Remove triple backticks and anything after (e.g. ```json)
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```") // fallback for plain code blocks
	s = strings.TrimSuffix(s, "```")
	s = strings.TrimSpace(s)
	return s
}
