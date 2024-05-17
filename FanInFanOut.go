package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Job represents a job to be processed, which is a URL
type Job struct {
	url string
}

// Result represents the result of a processed job
type Result struct {
	job    Job
	status int
	body   string
}

// worker fetches the content from the URLs in the jobs channel and sends results to the results channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		resp, err := http.Get(job.url)
		if err != nil {
			fmt.Printf("Worker %d failed to fetch %s: %v\n", id, job.url, err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Worker %d failed to read response body for %s: %v\n", id, job.url, err)
			continue
		}

		result := Result{job, resp.StatusCode, string(body)}
		fmt.Printf("Worker %d processed job %s with status %d\n", id, job.url, resp.StatusCode)
		results <- result

		// Simulate job processing time
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

// fanIn collects results from the results channel and processes them
func fanIn(results <-chan Result, done chan<- struct{}) {
	for result := range results {
		fmt.Printf("Result: job %s = status %d, body length %d\n", result.job.url, result.status, len(result.body))
	}
	done <- struct{}{}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
		"https://httpbin.org/get",
		"https://jsonplaceholder.typicode.com/todos/1",
	}

	const numWorkers = 3

	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))
	done := make(chan struct{})

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the jobs channel
	for _, url := range urls {
		jobs <- Job{url: url}
	}
	close(jobs)

	// Start a goroutine to close the results channel once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Start a goroutine to collect results (fan-in)
	go fanIn(results, done)

	// Wait for all results to be processed
	<-done
	fmt.Println("All jobs processed")
}
