package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()

	if !q.Empty() ||
		q.len != 0 ||
		q.Len() != 0 {
		t.Error()
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.queue[0] != 1 ||
		q.queue[1] != 2 ||
		q.queue[2] != 3 {
		fmt.Println(q.queue)
		t.Error()
	}

	if q.Len() != 3 {
		t.Error()
	}

	a, _ := q.Pop()

	if a != 1 || q.Len() != 2 {
		t.Error()
	}

	b := q.Front()

	if b != 2 {
		t.Error()
	}
}
