package datastructures

import (
	"errors"
)

type stackString struct {
	arr []string
}

func (e stackString) push(val string) {
	e.arr = append(e.arr, val)
}

func (e stackString) pop() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot pop from an empty stack")
	}
	val := e.arr[len(e.arr)-1]
	e.arr = e.arr[:len(e.arr)-1]
	return val, nil
}

func (e stackString) peek() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot peek from an empty stack")
	}
	return e.arr[len(e.arr)-1], nil
}

type queueString struct {
	arr []string
}

func (e queueString) enqueue(val string) {
	e.arr = append(e.arr, val)
}

func (e queueString) dequeue() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot dequeue from an empty queue")
	}
	val := e.arr[0]
	e.arr = e.arr[1:]
	return val, nil
}

func (e queueString) peek() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot peek from an empty queue")
	}
	return e.arr[0], nil
}