package stack

import "fmt"

type ArrayStack struct {
	// 栈的数据
	data []interface{}
	// 栈顶指针
	top  int
}

func NewArrayStack()*ArrayStack{
	return &ArrayStack{
		data:  make([]interface{},0,32),
		top:   -1,
	}
}

func (this *ArrayStack)IsEmpty()bool{
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack)Push(v interface{}){
	this.top += 1

	//超过栈的长度，扩容
	if this.top > len(this.data) - 1 {
		this.data = append(this.data,v)
	}else{
		this.data[this.top] = v
	}
}


func (this *ArrayStack)Pop()interface{}{
	if this.IsEmpty(){
		return nil
	}
	v := this.data[this.top]
	this.top--
	return v
}

func (this *ArrayStack)Top()interface{}{
	if this.IsEmpty(){
		return nil
	}
	return this.data[this.top]
}

func (this *ArrayStack)Flush(){
	this.top = -1
}

func (this *ArrayStack)Print(){
	if this.IsEmpty(){
		fmt.Println("stack is empty")
		return
	}
	for i := this.top; i >= 0; i--{
		fmt.Println(this.data[i])
	}
}

