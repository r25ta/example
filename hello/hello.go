package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//Set properties of the predefined Logger, including
	//the log entry prefix and a flag to disable printing
	//the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Ronaldo", "Olivia", "Diego", "Camila"}

	//Request a greeting message Hello function
	//message, err := greetings.Hello("Ronaldo")

	//Request a greeting message Hellos function
	messages, err := greetings.Hellos(names)

	//if an error was returned, print it to the console and
	//exit the program
	if err != nil {
		log.Fatal(err)
	}

	//if no error was returned, print the returned message
	//to the console
	fmt.Println(messages)
}
