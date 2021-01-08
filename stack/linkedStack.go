package stack

import "fmt"

type node struct {
	next   *node
	val    interface{}
}

// 涉及到链表的都用带头结点链表
type LinkedStack struct {
	// 栈顶节点
	topNode   *node
}

func NewLinkedStack()*LinkedStack{
	return &LinkedStack{nil}
}

func (this *LinkedStack)IsEmpty()bool{
	if this.topNode == nil{
		return true
	}
	return false
}

func (this *LinkedStack)Push(v interface{}){
	newNode := &node{nil,v}
	newNode.next = this.topNode
	this.topNode = newNode
}

func (this *LinkedStack)Pop()interface{}{
	if this.IsEmpty(){
		return nil
	}
	// 改变栈顶的头结点
	v := this.topNode.val
	this.topNode = this.topNode.next
	return v
}

func (this *LinkedStack)Top()interface{}{
	if this.IsEmpty(){
		return nil
	}
	return this.topNode.val
}

func (this *LinkedStack)Flush(){
	this.topNode = nil
}

func (this *LinkedStack)Print(){
	if this.IsEmpty(){
		fmt.Println("stack is empty")
		return
	}
	cur := this.topNode
	for cur != nil{
		// 打印数据
		fmt.Println(cur.val)
		cur = cur.next
	}
}


