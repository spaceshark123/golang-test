package datastructures_test

import (
	"testing"
	"github.com/spaceshark123/golang-test/src/utils/datastructures"
)

func TestStackPush(t *testing.T) {
	stack := datastructures.StackString{}
	stack.Push("hello")
	
	if stack.Len() != 1 {
		t.Errorf("expected: %d, recieved: %d", 1, stack.Len())
	}
}

func TestStackPop(t *testing.T) {
	stack := datastructures.StackString{}
	stack.Push("hello")
	stack.Push("hello2")
	stack.Push("hello3")

	val, err := stack.Pop() // should give hello3
	if val != "hello3" {
		t.Errorf("expected: %s, recieved: %s", "hello3", val)
	}
	if err != nil {
		t.Errorf("expected: %v, recieved: %v", nil, err)
	}
}

func TestStackPeek(t *testing.T) {
	stack := datastructures.StackString{}
	stack.Push("hello")
	stack.Push("hello2")

	val, err := stack.Peek() // should give hello2
	if val != "hello2" {
		t.Errorf("expected: %s, recieved: %s", "hello2", val)
	}
	if err != nil {
		t.Errorf("expected: %v, recieved: %v", nil, err)
	}
}

func TestStackEmpty(t *testing.T) {
	stack := datastructures.StackString{}
	
	val, err := stack.Peek() // should give error
	if val != "" {
		t.Errorf("expected: %s, recieved: %s", "", val)
	}
	if err == nil {
		t.Errorf("expected: %v, recieved: %v", "error", err)
	}
}