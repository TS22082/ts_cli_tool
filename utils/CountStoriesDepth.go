package utils

import "strings"

func CountStoriesDepth(s string) int {
	parts := strings.Split(s, "/")
	for i, part := range parts {
		if part == "src" {
			return len(parts) - i
		}
	}
	return 0
}
