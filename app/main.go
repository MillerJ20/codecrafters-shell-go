package main

import (
	"fmt";
	"log"
)

func main() {
	fmt.Print("$ ")

	var command string
	
	n, err := fmt.Scanln(&command)
	if err != nil {
		log.Fatal(err)
	}
	
	if n < 0 {
		log.Fatal("Wow bad number!")
	}

	fmt.Printf("%s: command not found", command)
}

