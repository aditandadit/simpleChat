package main

import (
	"flag"
	"fmt"
	"os"
	"./lib" // ./DirectoryName gives access to the Package
)

func main() {
	var isHost bool;

	// Send the Variable that gets the Value when Parsed
	flag.BoolVar(&isHost, "listen", false, "Listen on Specified Ip");
	// if go run main.go -h -> the last Arg in flag is Help Documentation

	// Parse Command Line Args
	flag.Parse();

	if isHost {
		// go run main.go -listen <ip>
		// Args 0 -> main.go 1 ->-listen 2 -> ip

		ip := os.Args[2];
		fmt.Println("Is Host")
		lib.RunHost(ip);
	} else {
		// go run main.go ip
		ip := os.Args[1]
		fmt.Println(" Is Guest")
		lib.RunGuest(ip);
	}
}
