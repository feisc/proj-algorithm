package linkedlist

import "fmt"

type LinkNode struct{
	next   *LinkNode
	value      interface{}
}

type LinkedList struct {
	head   *LinkNode
	length  uint
}

func (this *LinkNode)GetValue() interface{}{
	return this.value
}

func (this *LinkNode)GetNext() *LinkNode{
	return this.next
}

func NewLinkNode(v interface{}) *LinkNode{
	return &LinkNode{nil,v}
}

// 斗节点不存数据
func NewLinkedList() *LinkedList{
	return &LinkedList{NewLinkNode(0),0}
}

// p 为当前的指针
func (this *LinkedList)InsertAfter(p *LinkNode, v interface{}) bool{
	// 空指针不能插入
	if p == nil{
		return false
	}

	NewNode := NewLinkNode(v)
	OldNode := p.next // p 的next指针
	NewNode.next = OldNode
	p.next = NewNode
	this.length++
	return true
}

// 头结点
func (this *LinkedList)InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head,v)
}

func (this *LinkedList)InsertToTail(v interface{})bool{
	cur := this.head
	// 从头结点开始向后遍历,尾节点的特征next 为nil
	for nil != cur.next{
		cur = cur.next
	}
	return this.InsertAfter(cur,v)
}

func (this *LinkedList)InsertBefore(p *LinkNode, v interface{}) bool{
	// 边界条件,头结点不能插，插入的节点不能为nil
	if this.head == p || p == nil{
		return false
	}

	// 当前的从第一个节点开始，有头结点的话
	cur := this.head.next
	pre := this.head

	for nil != cur {
		if cur == p{
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	// 找到赋值
	NewNode := NewLinkNode(v)
	NewNode.next = cur
	pre.next = NewNode
	this.length ++
	return true

	//// 首先要找到 p 的前节点，从头开始遍历
	//cur := this.head
	//for cur.next != p && cur != nil{
	//	cur = cur.next
	//}
	//
	//// 找到最后没有找到
	//if cur == nil{
	//	return false
	//}
	//
	//// 找到p的前驱节点cur
	//NewNode := NewLinkNode(v)
	//NewNode.next = p
	//cur.next = NewNode
	//this.length++
	//return true
}

func (this *LinkedList)Reverse(){
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}
	pre := this.head.next
	cur := pre.next
	for cur != nil{
		pre.next = cur.next
		cur.next = this.head.next
		this.head.next = cur
		cur = pre.next
	}
}

func (this *LinkedList)FindMiddleNode() *LinkNode{
	// 通过长度来判断，只需要通过长度来判断
	length := this.length

	if 0 == length{
		return nil
	}

	mid := length/2
	if mid % 2 != 0{
		mid = length/2 + 1 // 奇数加1
	}

	cur := this.head
	for i := uint(0); i < mid; i++ {
		cur = cur.next
	}
	return cur
}

// 删除倒数第N个节点
func (this *LinkedList)DeleteRecNode(n int) bool {
	length := this.length
	if uint(n) > length{
		return false
	}

	// 倒数第几个，从开始循环
	count := length - uint(n)
	pre := this.head
	cur := this.head.next

	for i := uint(0); i < count; i++{
		pre = cur
		cur = cur.next
	}

	// 找到当前的节点
	pre.next = cur.next
	cur = nil
	this.length --
	return true
}

func MergeLinked(L1 *LinkedList, L2 *LinkedList) *LinkedList{
	// 还是可以通过长度来进行判断
	if L1.length == 0{
		return L2
	}
	if L2.length == 0 {
		return L1
	}

	// 申请一个新的链表,合并
	linkedList := NewLinkedList()
	cur := linkedList.head
	cur1 := L1.head.next
	cur2 := L2.head.next

	for nil != cur1 && nil != cur2{
		if cur1.value.(int) < cur2.value.(int){
			cur.next = cur1
			cur1 = cur1.next
		}else{
			cur.next = cur2
			cur2 = cur2.next
		}
	}

	// 如果其中的一个链表还没有遍历完，后面的直接赋值到linkedList
	for nil != cur1{
		cur.next = cur1
		cur1 = cur1.next
	}

	for nil != cur2{
		cur.next = cur2
		cur2 = cur2.next
	}
	linkedList.length = L1.length + L2.length
	return linkedList
}

func (this *LinkedList)DeleteNode(n int) bool {
	length := this.length
	if uint(n) > length || uint(n) == 0{
		return false
	}
	pre := this.head
	cur := this.head.next
	for i := uint(0); i < uint(n) - 1; i++{
		pre = cur
		cur = cur.next
	}
	pre.next = cur.next
	cur = nil
	this.length--
	return true
}

func (this *LinkedList)Print(){
	 cur := this.head.next
	 format := ""
	 for cur != nil{
		 format += fmt.Sprintf("%+v", cur.value)
		 cur = cur.next
		 if nil != cur {
			 format += "->"
		 }
	 }
	 fmt.Println(format)
}




