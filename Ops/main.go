package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./main.go <arguments>")
		os.Exit(1)
	}

	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("Invalid URL %v", err)
	}

	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Response Code: %d\nHttp Body: %s\n", response.StatusCode, string(body))

}

// Run from Terminal -
// go run main.go argument1 argument2 argument3
// echo $? "To print the last exit code"

/*
func main() {

	// Run platform commands or sub-processes

	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

	ipaddr, err := exec.Command("ifconfig", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", ipaddr)
	// Until this point here - it is still very much operating system specific coding.

}
*/
