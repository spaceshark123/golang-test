package datastructures

import (
	"errors"
)

type StackString struct {
	arr []string
}

func (e *StackString) Push(val string) {
	e.arr = append(e.arr, val)
}

func (e *StackString) Pop() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot pop from an empty stack")
	}
	val := e.arr[len(e.arr)-1]
	e.arr = e.arr[:len(e.arr)-1]
	return val, nil
}

func (e *StackString) Peek() (string, error) {
	if(len(e.arr) == 0) {
		return "", errors.New("Cannot peek from an empty stack")
	}
	return e.arr[len(e.arr)-1], nil
}

func (e *StackString) Len() int {
	return len(e.arr)
}