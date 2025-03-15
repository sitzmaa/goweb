package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"goweb/internal/interpreter"
)

func main() {
	if len(os.Args) > 1 {
		command := strings.Join(os.Args[1:], " ")

		// Check for YAML file input
		if strings.HasSuffix(os.Args[1], ".yaml") {
			interpreter.HandleYaml(os.Args[1])
		} else {
			interpreter.HandleCommand(command)
		}
	} else {
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
