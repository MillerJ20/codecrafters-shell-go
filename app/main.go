package main

import (
	"fmt"
)

func main() {
	for true{
		fmt.Print("$ ")

		var command string
		fmt.Scanln(&command)
		fmt.Printf("%s: command not found \n", command)
	}	
}

