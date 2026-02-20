package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
		scanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("$ ")

		scanner.Scan()

		parseCommand(scanner.Text())
	}	
}

func parseCommand(line string) {
	parts := strings.SplitN(line, " ", 2)
	command := parts[0]

	switch command {
	case "":
		fmt.Print("\r")
	case "exit":
		os.Exit(0)
	case "echo":
		if len(parts) > 1 {
			fmt.Printf("%s \n", parts[1])
		}
	case "type":
		builtins := []string{"exit", "echo", "type"}
		if len(parts) > 1 {
			commands := strings.Split(parts[1], " ")
			currCommand := commands[0]

			if slices.Contains(builtins, currCommand) {
					fmt.Printf("%s is a shell builtin \n", currCommand)
					break
			}

			fmt.Printf("%s: command not found \n", currCommand)
		}
	default:
		fmt.Printf("%s: command not found \n", command)
	}
}
