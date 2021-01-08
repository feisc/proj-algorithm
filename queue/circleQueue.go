package queue

type CircleQueue struct {
	data        []interface{}
	capacity    int
	head        int
	tail        int
}

func (this *CircleQueue)NewCircleQueue(n int)*CircleQueue{
	return &CircleQueue{make([]interface{},n),n,0,0}
}

func (this *CircleQueue)EnQueue(v interface{}) bool{
	// 队列满
	if this.head == (this.tail+1)%this.capacity{
		return false
	}
	this.data[this.tail] = v
	// 更新tail 的值
	this.tail = (this.capacity + 1) % this.capacity
	return true
}

func (this *CircleQueue)DeQueue()interface{}{
	if this.head == this.tail{
		return nil
	}
	v := this.data[this.head]
	this.head = (this.head + 1) % this.capacity
	return v
}