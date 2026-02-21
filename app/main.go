package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
	parts := strings.Fields(line)
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
		calculateTypes(parts)
	default:
		_, wasFound := findCommandInPATH(command)
		if wasFound {
			if len(parts) > 1 {
				cmd := exec.Command(command, parts[1:]...)
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
			cmd := exec.Command(command)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		fmt.Printf("%s: command not found \n", command)
	}
}

func calculateTypes(parts []string) {
	builtins := []string{"exit", "echo", "type"}
	if len(parts) > 1 {
		commands := strings.Fields(parts[1])
		currCommand := commands[0]

		if slices.Contains(builtins, currCommand) {
			fmt.Printf("%s is a shell builtin \n", currCommand)
			return
		}
		
		fullPath, wasFound := findCommandInPATH(currCommand)
		if wasFound {
			fmt.Printf("%s is %s\n", currCommand, fullPath)
			return
		}

		fmt.Printf("%s: not found \n", currCommand)
	}
}

func findCommandInPATH(command string) (fullPath string, wasFound bool){
	path, err := exec.LookPath(command)
	if err != nil {
		fullPath = ""
		wasFound = false
		return 
	}
	
	fullPath = path
	wasFound = true
	return
}
