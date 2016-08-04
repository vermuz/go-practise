package main

import (
	"fmt"
)

func main() {
	//name := "Mani"
	//var name string = "Mani"
	//fmt.Println("Hello, " + name)

	// Array of Strings
	names := []string {
	"Mani",
	"Ali",
	"Usman",
	"Faiz",
	}
	// Range should give us index and the name
	// of that index
	// n -> name
	// _ -> throw away the index value
	// _ -> Go's way of letting us capture something and 
	// throw it away
	for _, n := range names {
		fmt.Println("Hello, " + n)
	}
}
