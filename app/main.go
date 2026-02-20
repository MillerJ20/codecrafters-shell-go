package main

import (
	"fmt"
	"os"
)

func main() {
	for true{
		fmt.Print("$ ")

		var command string
		fmt.Scanln(&command)

		if command == "exit" {
			os.Exit(0)
		}
			
		fmt.Printf("%s: command not found \n", command)
	}	
}
