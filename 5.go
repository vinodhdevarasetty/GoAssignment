// Assignment 5: Pipeline Pattern Problem: Implement a pipeline pattern in Go where one goroutine
// generates numbers, another squares them, and a third prints the squared numbers.
package main

import (
	"fmt"
	"sync"
)

// Generator generates numbers from 1 to n and sends them to the out channel.
func generator(n int, out chan<- int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out)
}

// Squarer reads numbers from the in channel, squares them, and sends them to the out channel.
func squarer(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// Printer reads squared numbers from the in channel and prints them.
func printer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	// Define the channels for each stage of the pipeline
	nums := make(chan int)
	squares := make(chan int)

	// Define a WaitGroup to wait for the printer to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the goroutines for each stage of the pipeline
	go generator(10, nums) // Adjust the number to generate more or fewer numbers
	go squarer(nums, squares)
	go printer(squares, &wg)

	// Wait for the printer to finish
	wg.Wait()
}
