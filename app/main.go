package main

import (
	"fmt"
)

func main() {
	fmt.Print("$ ")

	var command string
	
	fmt.Scanln(&command)
	fmt.Printf("%s: command not found", command)
}

