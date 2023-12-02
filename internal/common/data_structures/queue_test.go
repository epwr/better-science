package data_structures

import (
	"testing"
)


func TestANewQueueIsEmpty(t *testing.T) {

	q := NewQueue[int]()

	assertTrue(t, q.IsEmpty(), "A new Queue should be empty")
}

func TestAQueueIsFIFO(t *testing.T) {

	q := NewQueue[int]()
	q.Push(5)
	q.Push(10)

	value, _ := q.Pop()
    assertTrue(t, value == 5, "Queue should return the added elements in FIFO order")
}

func TestCallingPopOnAnEmptyQueueReturnsAnError(t *testing.T) {

	q := NewQueue[int]()
	q.Push(0)
	q.Pop()

	_, err := q.Pop()
    assertError(t, err, "Pop() should return an error if the Queue is empty")
}
