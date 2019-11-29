package main

import (
	"fmt"
	"os"
	"strconv"
)

func portCheck(port string) {

	portNum, err := strconv.Atoi(port)
	if err != nil {
		// handle error
		fmt.Printf("Error specifying port %v\n", err)
		os.Exit(2)
	}
	if !portSane(portNum) {
		fmt.Printf("Invalid port %s because outside range 1025-65535\n", port)
		os.Exit(1)
	}
}

func portSane(port int) bool {

	if port <= 1024 {
		return false
	}

	if port > 65535 {
		return false
	}

	return true
}
