package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const nJobs = 5

	jobs := make(chan int, nJobs)
	results := make(chan int, nJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
		fmt.Println("w: ", w)
	}

	for j := 1; j <= nJobs; j++ {
		jobs <- j
		fmt.Println("j: ", j)
	}
	close(jobs)

	for a := 1; a <= nJobs; a++ {
		<-results
		fmt.Println("a: ", a)
	}
}
