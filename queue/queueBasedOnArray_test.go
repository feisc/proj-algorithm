package queue

import "testing"

func TestArrayQueueEnqueue(t *testing.T){
	q := NewArrayQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
}