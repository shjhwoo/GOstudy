package main

import (
	"fmt"
	"sync"
)

type Job interface {
	Do()
}

type SQJob struct {
	Cost int
}

func (j *SQJob) Do() {
	fmt.Println(j.Cost * j.Cost)
}

func main() {
	var JobList [10]SQJob
	for i := 0; i < 10; i++ {
		JobList[i] = SQJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := JobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
