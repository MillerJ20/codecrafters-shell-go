package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")

		scanner.Scan()

		command, args := parseCommand(strings.TrimSpace(scanner.Text()))

		switch command {
		case "":
			fmt.Print("\r")
		case "..":
			os.Chdir("../")
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Printf("%s \n", strings.Join(args, " "))
		case "type":
			calculateTypes(args)
		case "pwd":
			currDir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(currDir)
		case "cd":
			path, err := resolvePath(args)
			if err != nil {
				fmt.Printf("%s: %s: %s \n", command, args[0], err.Error())
				break
			}

			os.Chdir(path)
		default:
			_, wasFound := findCommandInPATH(command)
			if wasFound {
				if len(args) > 0 {
					cmd := exec.Command(command, args...)
					cmd.Stderr = os.Stderr
					cmd.Stdout = os.Stdout
					cmd.Run()
					break
				}
				cmd := exec.Command(command)
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout
				cmd.Run()
				break
			}
			fmt.Printf("%s: command not found \n", command)
		}
	}	
}

func parseCommand(line string) (command string, args []string){
	command = strings.Fields(line)[0]
	unparsedArgs := strings.TrimPrefix(line, command)
	
	inQuote := false
	escaped := false
	word := strings.Builder{}

	flush := func () {
		if len(word.String()) > 0 {
			args = append(args, word.String())
			word.Reset()
		}
	}

	for _, c := range unparsedArgs {
		switch  {
		case escaped:
			word.WriteRune(c)
			escaped = false
		case c == '\\':
			escaped = true
		case c == '\'': 
			inQuote = !inQuote
		case c == ' ' && !inQuote:
			flush()
		default:
			word.WriteRune(c)
		}
	}

	if escaped {
		word.WriteRune('\\')
	}

	flush()
	return command, args
}

func resolvePath(parts []string) (path string, err error){
	if len(parts) != 1 {
		return "", errors.New("Command only accepts a single parameter of directory to change to")
	}

	switch parts[0] {
	case "~":
		dirname, err := os.UserHomeDir()
		if err!= nil {
			return "", errors.New("Failed to traverse to directory")
		}
		return dirname, nil

	default:
		file, err := os.Stat(parts[0])
		if err != nil {
			return "", errors.New("No such file or directory")
		}

		if file.IsDir() {
			return parts[0], nil
		}

		return "", errors.New("Path entered is not a directory")
	}
}

func calculateTypes(parts []string) {
	builtins := []string{"exit", "echo", "type", "pwd", "cd"}
	if len(parts) > 0 {
		commands := strings.Fields(parts[0])
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
