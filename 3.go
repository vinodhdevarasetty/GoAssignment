package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex
var sharedInt int

func main() {
	groutines := 10
	iterations := 1000 // Number of increments each goroutine will perform

	wg.Add(groutines)
	for i := 0; i < groutines; i++ {
		go incr(iterations)
	}
	wg.Wait()
	fmt.Println("Final value:", sharedInt)
}

func incr(iterations int) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		mtx.Lock()
		sharedInt++
		mtx.Unlock()
	}
}
