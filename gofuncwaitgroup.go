package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a channel
	queue := make(chan string)

	// To drain all routines before exiting
	// Wait Groups
	var w sync.WaitGroup
	w.Add(2) // Wait for two things

        // ... This feature actually blocks waiting for queue
	// We would prefer to print out queue while it is draining
	// Go Routines
	go func(queue chan string, wg *sync.WaitGroup) {
		// Wait for function to finish
		defer wg.Done()
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")
		queue <- "done func1"
	}(queue, &w) 
        
	go func(queue chan string, wg *sync.WaitGroup) {
		// Wait for function to finish
		defer wg.Done()
		fmt.Println("func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("func 2")
		queue <- "done func2"
	}(queue, &w)
	
	// An Inefficient way of catching up with go compiler
	// time.Sleep(4 * time.Second)
	// fmt.Println(<-queue)
	// fmt.Println(<-queue)
	go func(queue chan string) {
		for text := range queue {
			fmt.Println(text)
	        }
        }(queue)
	// Block execution until Wait Groups are finished
	w.Wait()
}


