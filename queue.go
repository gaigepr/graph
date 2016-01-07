package graph

import ()

//
type node struct {
	data *Vertex
	next *node
}

//
type Queue struct {
	head  *node
	tail  *node
	count int
}

//
func NewQueue() *Queue {
	return &Queue{}
}

//
func (q *Queue) Length() int {
	return q.count
}

//
func (q *Queue) Push(v *Vertex) {
	n := &node{data: v}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

//
func (q *Queue) Poll() *Vertex {
	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}

	q.count--

	return n.data
}

//
func (q *Queue) Peek() *Vertex {
	if q.head == nil {
		return nil
	}

	return q.head
}
