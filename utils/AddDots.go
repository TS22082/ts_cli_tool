package utils

import "fmt"

func AddDots(depth int) string {
	var result string

	for i := 1; i < depth; i++ {
		result += "../"
	}

	return fmt.Sprintf("%sstories/tests/setupTests", result)
}
