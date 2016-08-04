package main

import (
	"fmt"
)

// Capitalization decides on visibility
// outside of the package
// Capitalized -> Public
// otherwise private
func Greet(name string) {
	fmt.Println("Hello, "+ name)
}

func GreetNames(names []string, suffix string) {
	for _, n := range names {
		Greet(n + suffix)
	}
}

func main() {
	names := []string {
	"X",
	"Y",
	"Z",
	"W",
	}
	// Execute this in a brand new thread
	// Concurrently
	// GreetNames(names)
	// GreetNames(names)
	// Channel
	comm := make(chan string) // Channel of type string
        // Anonymous Function
	go func() {
		GreetNames(names, " <C> ")
		// Send a message to the channel
		comm <- "Finished greeting names concurrently...."
	}()
	//go GreetNames(names, " <C> ") // Concurrent code
	GreetNames(names, " <M> ") // Main loop
	// Listen for communication from channel
	// Pipe info out of the channel
	fmt.Println(<- comm)
}

