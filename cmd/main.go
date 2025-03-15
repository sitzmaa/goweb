package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"goweb/internal/interpreter"
)

func main() {
	// If there are command-line arguments, process them
	if len(os.Args) > 1 {
		// Join the command arguments into a single string
		command := strings.Join(os.Args[1:], " ")

		// Send the entire command string to the interpreter to handle it
		interpreter.HandleCommand(command)
	} else {
		// Otherwise, enter interactive mode
		interactiveMode()
	}
}

func interactiveMode() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("goweb> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		// Detect YAML input interactively
		if strings.HasSuffix(input, ".yaml") {
			interpreter.HandleYaml(input)
		} else {
			interpreter.HandleCommand(input)
		}
	}
}
