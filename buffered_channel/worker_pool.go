package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id        int
	randomno  int
	worker_id int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit
		number /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(worker_id int, wg *sync.WaitGroup) {
	for job := range jobs {
		job.worker_id = worker_id
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		worker_id := i
		go worker(worker_id, &wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno, -1}
		jobs <- job
	}
	fmt.Printf("jobs channel ready to close...\n")
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d at worker %d, input random no %d , sum of digits %d\n", result.job.id, result.job.worker_id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
