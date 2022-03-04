/**
  Worker pool
  ------------------------------------------------------------
  keywords: concurrency, thread pool, queue
  ------------------------------------------------------------
  What is worker pool?

  Worker pools are a model in which a fixed number of m workers
  work their way through n tasks in a task queue. Task stays in
  a queue until a worker finished up its current task and pulls
  a new one off.
  ------------------------------------------------------------
  Why we need worker pool?

  1. To amortize the costs of creating new threads
  2. To control the number of concurrently running tasks (if spawn
     too many goroutine, your machine will quickly run out of memory
     and the CPU)
  ------------------------------------------------------------
  Implementation

  Use goroutine to spawn the worker and implement queue by channel
*/

package main

import (
	"errors"
	"fmt"
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
	tasksSize   int
	tasksChan   chan Task
	resultsChan chan Task
	Results     func() []Task
}

func NewWorkerPool(tasks []Task, size int) *WorkerPool {
	tasksChan := make(chan Task, len(tasks))
	resultsChan := make(chan Task, len(tasks))

	// create tasks channel
	for _, t := range tasks {
		tasksChan <- t
	}
	close(tasksChan) // close the channel to indicate that’s all the work we have.

	wp := &WorkerPool{
		PoolSize:    size,
		tasksSize:   len(tasks),
		tasksChan:   tasksChan,
		resultsChan: resultsChan,
	}

	wp.Results = wp.results
	return wp
}

// Do a task and receive result(err) from task. Finally, send the result to results channel
func (pool *WorkerPool) worker() {
	for t := range pool.tasksChan {
		t.Err = t.Do()
		pool.resultsChan <- t
	}
}

// Retrieve results of the tasks
func (pool *WorkerPool) results() []Task {
	tasks := make([]Task, pool.tasksSize)
	for i := 0; i < pool.tasksSize; i++ {
		tasks[i] = <-pool.resultsChan
	}

	return tasks
}

func (pool *WorkerPool) Start() {
	for i := 0; i < pool.PoolSize; i++ {
		go pool.worker()
	}
}

// Implement a worker pool using goroutines and channels
func main() {
	t := time.Now()

	const numWorkers = 2

	tasks := []Task{
		{Id: 0, f: func() error { time.Sleep(10 * time.Second); fmt.Println(0); return nil }},
		{Id: 1, f: func() error { time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
		{Id: 2, f: func() error { fmt.Println(2); return errors.New("error") }},
	}

	// create a work pool that includes the tasks and [numWorkers] of workers(pool size)
	wp := NewWorkerPool(tasks, numWorkers)
	wp.Start()

	tasks = wp.Results()
	fmt.Printf("All tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
	for _, t := range tasks {
		fmt.Printf("result of task %d is %v\n", t.Id, t.Err)
	}

}
