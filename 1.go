// Assignment 1: Goroutine with Channel Problem: Write a Go program that calculates the
//  sum of numbers from 1 to N concurrently using goroutines and channels.
//   The program should take the value of N as input from the user.

package main

import (
	"fmt"
	"sync"
)

func sum(n int, resultChan chan<- int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	resultChan <- sum
}

func main() {
	var n int
	fmt.Print("Enter the value of N: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resultChan := make(chan int)
	var wg sync.WaitGroup

	// Launch goroutine to calculate sum
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum(n, resultChan)
	}()

	// Receive result from channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	total := 0
	for res := range resultChan {
		total += res
	}

	fmt.Println("Sum of numbers from 1 to", n, "is:", total)
}
