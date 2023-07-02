package main

import (
	"fmt"
	"log"
)

func main() {
	var myString = "Green"

	fmt.Println("My string is set to ", myString)
	changeUsingPointer(&myString)
	fmt.Println("My string is set to ", myString)
}

func changeUsingPointer(s *string) {
	log.Print("s is set to ", s)
	newValue := "Red"
	*s = newValue

}

