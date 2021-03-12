package main

import (
	"fmt"
	"log"

	greetings "example.com/greeting"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Coz", "Glitch", "Amber"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
