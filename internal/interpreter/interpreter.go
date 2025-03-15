package interpreter

import (
	"goweb/internal/executor"
	"goweb/internal/formatter"
	"goweb/internal/receiver"
)

// HandleCommand processes the command input from the user.
// Precondition: The command is a valid string input.
// Postcondition: The command is parsed, executed, and the response is handled.
func HandleCommand(command string) {
	// Step 1: Parse the command into an API request
	apiRequest := formatter.FormatCommand(command)

	// Step 2: Execute the API request
	response := executor.ExecuteRequest(apiRequest)

	// Step 3: Handle the response
	receiver.HandleResponse(response)
}

// HandleYaml processes the YAML file input from the user.
// Precondition: The file path is a valid string input.
// Postcondition: The YAML file is parsed, executed, and the response is handled.
func HandleYaml(filePath string) {
	// Implementation for handling YAML input goes here
}
