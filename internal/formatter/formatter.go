package formatter

import "strings"

// FormatCommand formats the raw command input into an API request.
// Precondition: command is a non-empty string.
// Postcondition: returns a structured API request string.
func FormatCommand(command string) string {
	// Simple example of converting the command to lowercase and trimming spaces.
	return strings.TrimSpace(strings.ToLower(command))
}

func FormatYaml(filePath string) {

}
