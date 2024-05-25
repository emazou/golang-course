package main

import (
	"fmt"
	"github.com/emazou/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, error := greetings.Hellos(names)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(messages)
}
