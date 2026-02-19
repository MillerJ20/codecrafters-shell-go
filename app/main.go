package main

import (
	"fmt"
)

func main() {
	fmt.Print("$ ")

	var command string
	/*
	n, err := fmt.Scanln(&command)
	if err != nil {
		log.Fatal(err)
	}
	*/

	fmt.Printf("%s: command not found", command)
}

