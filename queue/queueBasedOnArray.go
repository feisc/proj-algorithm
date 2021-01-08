package queue

import "fmt"

type ArrayQueue struct{
	data      []interface{}
	capacity  int
	head      int
	tail      int
}

func NewArrayQueue(n int)*ArrayQueue{
	return &ArrayQueue{make([]interface{},n),n,0,0}
}

func (this *ArrayQueue)IsFull() bool{
	if this.tail == this.capacity{
		return true
	}
	return false
}

func (this *ArrayQueue)Enqueue(v interface{})bool{
	if this.IsFull(){
		return false
	}
	this.data = append(this.data,v)
	this.tail++
	return true
}

func (this *ArrayQueue)Dequeue()interface{}{
	// 空对列
	if this.head == this.tail{
		return nil
	}
	v := this.data[this.head]
	this.head++
	return v
}

func (this *ArrayQueue)String()string{
	if this.head == this.tail{
		return "empty queue"
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", this.data[i])
	}
	result += "<-tail"
	return result
}
