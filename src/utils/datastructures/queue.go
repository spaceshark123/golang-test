package datastructures

import (
	"errors"
)

type QueueString struct {
	arr []string
}

func (e *QueueString) Enqueue(val string) {
	e.arr = append(e.arr, val)
}

func (e *QueueString) Dequeue() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot dequeue from an empty queue")
	}
	val := e.arr[0]
	e.arr = e.arr[1:]
	return val, nil
}

func (e *QueueString) Peek() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot peek from an empty queue")
	}
	return e.arr[0], nil
}

func (e *QueueString) Len() int {
	return len(e.arr)
}