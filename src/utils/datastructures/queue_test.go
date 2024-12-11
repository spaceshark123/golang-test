package datastructures_test

import (
	"testing"
	"github.com/spaceshark123/golang-test/src/utils/datastructures"
)

func TestQueueEnqueue(t *testing.T) {
	queue := datastructures.QueueString{}
	queue.Enqueue("hello")

	if queue.Len() != 1 {
		t.Errorf("expected: %d, recieved: %d", 1, queue.Len())
	}
}

func TestQueueDequeue(t *testing.T) {
	queue := datastructures.QueueString{}
	queue.Enqueue("hello")
	queue.Enqueue("hello2")
	queue.Enqueue("hello3")

	val, err := queue.Dequeue() // should give hello
	if val != "hello" {
		t.Errorf("expected: %s, recieved: %s", "hello", val)
	}
	if err != nil {
		t.Errorf("expected: %v, recieved: %v", nil, err)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := datastructures.QueueString{}
	queue.Enqueue("hello")
	queue.Enqueue("hello2")

	val, err := queue.Peek() // should give hello
	if val != "hello" {
		t.Errorf("expected: %s, recieved: %s", "hello", val)
	}
	if err != nil {
		t.Errorf("expected: %v, recieved: %v", nil, err)
	}
}

func TestQueueEmpty(t *testing.T) {
	queue := datastructures.QueueString{}
	
	val, err := queue.Peek() // should give error
	if val != "" {
		t.Errorf("expected: %s, recieved: %s", "", val)
	}
	if err == nil {
		t.Errorf("expected: %v, recieved: %v", "error", err)
	}
}