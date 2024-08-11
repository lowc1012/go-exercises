package workerpool

import (
	"fmt"
	"testing"
)

func TestTask_Do(t *testing.T) {
	task := Task{
		Id: 1, f: func() error {
			return fmt.Errorf("test error")
		},
	}

	if err := task.Do(); err == nil || err.Error() != "test error" {
		t.Fatalf("Expected error 'test error', got %v", err)
	}
}

func TestNewWorkerPool(t *testing.T) {

}

func TestWorkerPool_Start(t *testing.T) {
}
