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

// Example: Implement a worker pool using goroutines and channels
// func example() {
// 	t := time.Now()
// 	tasks := []Task{
// 		{Id: 0, f: func() error { fmt.Printf("executing task %d\n", 0); time.Sleep(10 * time.Second); return nil }},
// 		{Id: 1, f: func() error { fmt.Printf("executing task %d\n", 1); time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
// 		{Id: 2, f: func() error { fmt.Printf("executing task %d\n", 2); return errors.New("error") }},
// 	}
//
// 	// create a work pool that includes the tasks and [numWorkers] of workers(pool size)
// 	const numWorkers = 2
// 	wp := NewWorkerPool(numWorkers)
// 	wp.Start()
//
// 	// submit tasks
// 	for _, task := range tasks {
// 		wp.Submit(task)
// 	}
//
// 	// (blocking) Stop the worker pool after all tasks are submitted
// 	wp.Stop()
//
// 	// Print results from the worker pool resultsChan
// 	for r := range wp.Results() {
// 		fmt.Printf("result of task %d is %v\n", r.Id, r.Err)
// 	}
//
// 	fmt.Printf("All tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
// }

package workerpool

import (
	"sync"
)

type Task struct {
	Id  int
	Err error
	F   func() error
}

// Do execute a task
func (t *Task) Do() error {
	return t.F()
}

type WorkerPool struct {
	PoolSize    int
	tasksChan   chan Task
	resultsChan chan Task
	quit        chan struct{}
	workersWg   sync.WaitGroup
}

// Submit submits a task into tasks channel
func (pool *WorkerPool) Submit(task Task) {
	pool.tasksChan <- task
}

// NewWorkerPool initialize a worker pool
func NewWorkerPool(numWorker int) *WorkerPool {
	const buffSize = 100
	wp := &WorkerPool{
		PoolSize:    numWorker,
		tasksChan:   make(chan Task, buffSize),
		resultsChan: make(chan Task, buffSize),
		quit:        make(chan struct{}),
	}

	return wp
}

// Results Retrieve results of the tasks
func (pool *WorkerPool) Results() <-chan Task {
	return pool.resultsChan
}

// worker keeps consuming tasks from `tasksChan`, and outputting the result to resultsChan
func (pool *WorkerPool) worker() {
	// decrement the WaitGroup counter when the worker exits
	defer pool.workersWg.Done()

	for {
		select {
		case <-pool.quit:
			//  worker exits if the quit channel is closed or a message is received
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
	close(pool.tasksChan) // close the tasksChan to make all workers exits
	pool.workersWg.Wait() // waits for all workers exists
	close(pool.resultsChan)
}
