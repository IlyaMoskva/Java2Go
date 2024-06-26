package main

import (
	"fmt"
)

func writeDestinations(dc chan string) {
	dest := []string{"London", "Paris", "New York", "Milan", "Sydney"}
	for _, d := range dest {
		fmt.Printf("[Writer]: Writing %v to channel.\n", d)
		dc <- d
	}
	close(dc)
	fmt.Printf("[Writer]: Shuting down.\n")
}

func main() {
	dc := make(chan string)
	go writeDestinations(dc)

	for d := range dc {
		fmt.Printf("[Main]: Hello from %v!\n", d)
	}
	fmt.Printf("[Main]: Shuting down.")
}
