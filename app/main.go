package main

import (
	"bufio"
	"fmt"
	"os"
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
		if len(parts) > 1{
			fmt.Printf(" %s \n", parts[1])
		}
	default:
		fmt.Printf("%s: command not found \n", command)
	}
}
