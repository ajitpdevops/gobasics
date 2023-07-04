package main

import (
	"log"
	"net/http"

	"github.com/whoajitpatil/gobasics/pkg/handler"
)

const portNumber = ":8081"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	log.Printf("Starting the web application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
