// Assignment 4: Goroutine with Timeout Problem: Write a Go program that performs a task in a
//  goroutine and waits for it to finish. However, if the task takes more than 3 seconds, the
//  program should print a timeout message and exit.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a channel to signal completion of the task
	done := make(chan bool)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Start the goroutine
	go func() {
		// Simulate a long-running task
		// Generate a random sleep duration between 1 and 6 seconds
		randomDuration := time.Duration(r.Intn(3)+1) * time.Second
		time.Sleep(randomDuration)
		done <- true
	}()

	// Create a timeout channel
	timeout := time.After(3 * time.Second)

	// Use select to wait for either the task to complete or the timeout
	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-timeout:
		fmt.Println("Timeout: Task took too long to complete")
	}
}
