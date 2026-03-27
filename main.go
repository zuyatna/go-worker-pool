package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TransferJob struct {
	ID     int
	From   string
	To     string
	Amount float64
}

func worker(ctx context.Context, workerID int, jobs <-chan TransferJob, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d halting: %v\n", workerID, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d shutting down smoothly: no more jobs\n", workerID)
				return
			}
			fmt.Printf("worker %d processing transfer id %d\n", workerID, job.ID)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan TransferJob, numJobs)
	var wg sync.WaitGroup

	fmt.Printf("starting workers...\n")
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	fmt.Printf("sending jobs...\n")
	for j := 1; j <= numJobs; j++ {
		jobs <- TransferJob{
			ID:     j,
			From:   "Account A",
			To:     "Account B",
			Amount: 100.0,
		}
	}
	close(jobs)

	wg.Wait()
	fmt.Printf("all workers finished, batch process completed")
}
