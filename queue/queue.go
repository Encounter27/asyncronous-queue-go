package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func New() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)

	return queue
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
}

func (q *Queue) Empty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len == 0
}

func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, item)
	q.len++
}

func (q *Queue) Pop() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == 0 {
		return nil, fmt.Errorf("Queue is empty")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]
	q.len--

	return item, nil
}

func (q *Queue) Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.len == 0 {
		return nil
	}

	item := q.queue[0]

	return item
}
