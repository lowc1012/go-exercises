package queue

import (
	"fmt"
	"sync"
)

type T interface{}

type RingBuffer struct {
	sync.Mutex
	q        []T
	capacity int
	head     int
	tail     int
	full     bool
}

// NewRingBuffer creates a new RingBuffer with the given capacity
func NewRingBuffer(capacity int) *RingBuffer {
    return &RingBuffer{
        q:        make([]T, capacity),
        capacity: capacity,
    }
}

// Enqueue adds an element to the end of the queue.
// If the queue is full, it returns an error.

func (r *RingBuffer) Enqueue(data T) error {
	r.Lock() // thread-safe
	defer r.Unlock()

	if r.IsFull() {
		return fmt.Errorf("The queue is full")
	}

	r.q[r.tail] = data
	r.tail = (r.tail + 1) % r.capacity
	r.full = r.head == r.tail
	return nil
}

func (r *RingBuffer) Dequeue() (T, error) {
	r.Lock() // thread-safe
	defer r.Unlock()

	if r.IsEmpty() {
		return nil, fmt.Errorf("The queue is empty")
	}

	// get data from the head of the queue
	data := r.q[r.head]

	// move head forward
	r.head = (r.head + 1) % r.capacity
	r.full = false
	return data, nil
}

func (r *RingBuffer) IsEmpty() bool {
	return r.head == r.tail && !r.full
}

func (r *RingBuffer) IsFull() bool {
	return r.full
}
