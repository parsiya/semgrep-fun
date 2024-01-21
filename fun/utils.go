package fun

import (
	"os"
	"strings"
)

// WriteFile writes the given data to the given file.
func WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// Returns the top-level package. Split the input by `/` and return the last
// part.
func TopLevelPackageOld(input string) string {
	parts := strings.Split(input, "/")
	return parts[len(parts)-1]
}

// Returns the top-level package. Start from the end and find the first `/`.
// Should be a bit faster than the above, not that it really matters.
func TopLevelPackage(input string) string {
	i := len(input) - 1
	for i >= 0 && input[i] != '/' {
		i--
	}
	return input[i+1:]
}
