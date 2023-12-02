package data_structures

import (
	"errors"
	"fmt"
)


// A FIFO datastructure that provides Pop(), Push(), Peak() and IsEmpty() methods.
type Queue[T any] struct {
	data []T
}

// Returns an empty Queue
func NewQueue[T any]() Queue[T] {

	return Queue[T]{
		data: []T{},
	}
}

// Checks if the Queue is empty
func (q *Queue[T]) IsEmpty() bool {

	return len(q.data) == 0
	
}

// Pushes a new element into the Queue
func (q *Queue[T]) Push(element T) {

	q.data = append(q.data, element)
}

// Removes and returns the front element in the Queue
func (q *Queue[T]) Pop() (T, error) {

	if q.IsEmpty() {
		var emptyResult T
		fmt.Println("emptyResult:", emptyResult)
		return emptyResult, errors.New("Pop() called on an empty Queue.")
	}
	firstElement := q.data[0]
	q.data = q.data[1:]
	return firstElement, nil
}
