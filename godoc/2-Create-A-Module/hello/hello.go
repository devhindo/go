package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    message, err := greetings.Hello("devhindo")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(message)
}