/**
  Worker pool
  ------------------------------------------------------------
  keywords: concurrency, thread pool, queue
  ------------------------------------------------------------
  What is worker pool?

  Worker pool is a model which a fixed number of m workers
  work their way through n tasks in a task queue. Task stays in
  a queue until a worker finished up its current task and pulls
  a new one off.
  ------------------------------------------------------------
  Why we need Worker pool model?

  1. To amortize the costs of creating new threads
  2. To control the number of concurrently running tasks (if program
     spawn too many goroutine, your machine will quickly run out of
     memory and the CPU)
  ------------------------------------------------------------
  Implementation

  Use goroutine to spawn the worker and implement queue by channel
*/

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id  int
	Err error
	f   func() error
}

// Do a task
func (t *Task) Do() error {
	return t.f()
}

type WorkerPool struct {
	PoolSize    int
	tasksChan   chan Task
	resultsChan chan Task
	quit        chan bool
	workersWg   sync.WaitGroup
}

// Submit submits a task into tasks channel
func (pool *WorkerPool) Submit(task Task) {
	pool.tasksChan <- task
}

func NewWorkerPool(numWorker int) *WorkerPool {
	const buffSize = 100
	wp := &WorkerPool{
		PoolSize:    numWorker,
		tasksChan:   make(chan Task, buffSize),
		resultsChan: make(chan Task, buffSize),
		quit:        make(chan bool),
	}

	return wp
}

// Results Retrieve results of the tasks
func (pool *WorkerPool) Results() <-chan Task {
	return pool.resultsChan
}

// worker keeps consuming tasks from `tasksChan`, and sends the result to results channel
func (pool *WorkerPool) worker() {
	// reduce the waiting number of workers if the worker exists
	defer pool.workersWg.Done()

	for {
		select {
		case <-pool.quit:
			return
		case task, ok := <-pool.tasksChan:
			if !ok {
				return // worker exits if tasksChan is closed
			}
			task.Err = task.Do()
			pool.resultsChan <- task
		}
	}
}

// Start create worker goroutines with size of PoolSize
func (pool *WorkerPool) Start() {
	// set the counter for WaitGroup
	pool.workersWg.Add(pool.PoolSize)
	for _ = range pool.PoolSize {
		go pool.worker()
	}
}

func (pool *WorkerPool) Stop() {
	close(pool.tasksChan)

	pool.workersWg.Wait() // waits for all workers exists
	close(pool.resultsChan)
	close(pool.quit)
}

// Implement a worker pool using goroutines and channels
func main() {
	t := time.Now()
	tasks := []Task{
		{Id: 0, f: func() error { fmt.Printf("executing task %d\n", 0); time.Sleep(10 * time.Second); return nil }},
		{Id: 1, f: func() error { fmt.Printf("executing task %d\n", 1); time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
		{Id: 2, f: func() error { fmt.Printf("executing task %d\n", 2); return errors.New("error") }},
	}

	// create a work pool that includes the tasks and [numWorkers] of workers(pool size)
	const numWorkers = 2
	wp := NewWorkerPool(numWorkers)
	wp.Start()

	// submit tasks
	for _, task := range tasks {
		wp.Submit(task)
	}

	go func() {
		wp.Stop()
	}()

	results := wp.Results()
	for r := range results {
		fmt.Printf("result of task %d is %v\n", r.Id, r.Err)
	}

	fmt.Printf("All tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
}
