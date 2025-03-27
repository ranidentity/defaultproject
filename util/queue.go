package util

import "fmt"

type Queue struct {
	items []any
}

// Enqueue adds an element to the queue
func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first element from the queue
func (q *Queue) Dequeue() (any, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:] // Remove the first element
	return item, true
}

func (qq *Queue) Test() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue()) // Output: 1, true
	fmt.Println(q.Dequeue()) // Output: 2, true
}
