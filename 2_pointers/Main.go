package main

import (
	"log"
)

func main() {
	myString := "Green"
	log.Println("My String is set to", myString)
	changeWithPointer(&myString)
	log.Println("After a function call - My String is set to", myString)

}

func changeWithPointer(s *string) {
	log.Println("s is set to ", s)
	newValue := "Red"
	*s = newValue
}
