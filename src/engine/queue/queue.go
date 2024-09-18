package queue

import "log"

type node struct {
	next *node
	prev *node
	item interface{}
}

type Queue struct {
	head *node
	tail *node
	len  int
}

func New(items ...interface{}) *Queue {
	queue := &Queue{nil, nil, len(items)}
	if queue.len == 0 {
		return queue
	}

	queue.head = &node{nil, nil, items[0]}
	prev := queue.head
	for i := 1; i < len(items); i++ {
		node := &node{nil, prev, items[i]}
		prev.next = node
		prev = node
	}

	queue.tail = prev
	return queue
}

func (q *Queue) Push(item interface{}) {
	if q == nil {
		log.Fatal("Trying to push in a queue wich is nil")
		return
	}

	q.len++

	if q.head == nil {
		q.head = &node{nil, nil, item}
		q.tail = q.head
		return
	}

	n := &node{nil, q.tail, item}
	q.tail.next = n
	q.tail = n
}

func (q *Queue) Pop() interface{} {
	if q == nil || q.head == nil {
		return nil
	}

	head := q.head
	q.head = head.next

	if q.tail == head {
		q.tail = nil
	}

	q.len--
	return head.item
}

func (q *Queue) Len() int {
	return q.len
}
