package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 100
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	fmt.Println("finished add job.")
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// fmt.Println(<-results)
		<-results
	}
}
