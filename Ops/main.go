package main

import (
	"fmt"
	"log"
	"os/exec"
)

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
