// Assignment 2: Producer-Consumer with Channel Problem: Implement the producer-consumer problem using goroutines and channels. 
//The producer should generate numbers from 1 to 100 and send them to a channel, and the consumer should print those numbers.
package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	wg.Done()
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Consumed:", num)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
