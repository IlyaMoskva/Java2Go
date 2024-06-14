package main

import (
	"testing"
	"time"
)

func TestWriteDestinations(t *testing.T) {
	dc := make(chan string)
	go writeDestinations(dc)

	expectedDestinations := []string{"London", "Paris", "New York", "Milan", "Sydney"}
	receivedDestinations := make([]string, 0, len(expectedDestinations))

	for d := range dc {
		receivedDestinations = append(receivedDestinations, d)
	}

	if len(receivedDestinations) != len(expectedDestinations) {
		t.Errorf("Expected %d destinations, but got %d", len(expectedDestinations), len(receivedDestinations))
	}

	for i, d := range receivedDestinations {
		if d != expectedDestinations[i] {
			t.Errorf("Expected destination %s at index %d, but got %s", expectedDestinations[i], i, d)
		}
	}
}

func TestWriteDestinationsTimeout(t *testing.T) {
	dc := make(chan string)
	go writeDestinations(dc)

	expectedDestinations := []string{"London", "Paris", "New York", "Milan", "Sydney"}
	receivedDestinations := make([]string, 0, len(expectedDestinations))

	timeout := time.After(2 * time.Second)
	for {
		select {
		case d, ok := <-dc:
			if !ok {
				if len(receivedDestinations) != len(expectedDestinations) {
					t.Errorf("Expected %d destinations, but got %d", len(expectedDestinations), len(receivedDestinations))
				}
				for i, d := range receivedDestinations {
					if d != expectedDestinations[i] {
						t.Errorf("Expected destination %s at index %d, but got %s", expectedDestinations[i], i, d)
					}
				}
				return
			}
			receivedDestinations = append(receivedDestinations, d)
		case <-timeout:
			t.Fatalf("Test timed out. Expected %d destinations, but got %d", len(expectedDestinations), len(receivedDestinations))
		}
	}
}
