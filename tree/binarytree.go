package tree

import (
	"fmt"
)

type TreeNode struct{
	val    interface{}
	left   *TreeNode
	right  *TreeNode
}

type BinaryTree struct{
	root *TreeNode
}

func NewTreeNode(v interface{}) *TreeNode{
	return &TreeNode{v,nil,nil}
}

func NewBinaryTree(v interface{})*BinaryTree{
	return &BinaryTree{NewTreeNode(v)}
}

func (this *BinaryTree)PreOrderTraverse(node *TreeNode){
	if node == nil{
		return
	}
	fmt.Println(node.val)
	this.PreOrderTraverse(node.left)
	this.PreOrderTraverse(node.right)
}

func (this *BinaryTree)InOrderTraverse(node *TreeNode){
	if node == nil{
		return
	}
	this.InOrderTraverse(node.left)
	fmt.Println(node.val)
	this.InOrderTraverse(node.right)
}

func (this *BinaryTree)PostOrderTraverse(node *TreeNode){
	if node == nil {
		return
	}
	this.PostOrderTraverse(node.left)
	this.PostOrderTraverse(node.right)
	fmt.Println(node.val)
}

// 用迭代的方式进行中序遍历
func (this *BinaryTree)InOrderTraverseRec(){
	node := this.root
	// 打印地址
	//fmt.Println(unsafe.Pointer(node),unsafe.Pointer(this.root))

	stack := NewArrayStack()
	for !stack.IsEmpty() || node != nil{
		for node != nil{
			stack.Push(node)
			node = node.left
		}
		if !stack.IsEmpty(){
			node = stack.Top().(*TreeNode)
			stack.Pop()
			fmt.Println(node.val)
			node = node.right
		}
	}
}

func (this *BinaryTree)PreOrderTraverseRec(){
	node := this.root
	stack := NewArrayStack()
	for !stack.IsEmpty() || node != nil{
		for node != nil{
			fmt.Println(node.val)
			stack.Push(node)
			node = node.left
		}
		if !stack.IsEmpty(){
			node := stack.Top().(*TreeNode)
			stack.Pop()
			node = node.right
		}
	}
}

func (this *BinaryTree)PostOrderTraverseRec(){
	node := this.root
	if node == nil{
		return
	}
	stack := NewArrayStack()
	var pNodeVisited *TreeNode
	// 把所有的左子树放到stack
	for node != nil{
		stack.Push(node)
		node = node.left
	}
	for !stack.IsEmpty(){
		// 先取出左子树
		node = stack.Top().(*TreeNode)
		stack.Pop()
		//一个节点被访问的条件是右子树为空，或者右子树已被访问
		if node.right == nil || node.right == pNodeVisited{
			fmt.Println(node.val)
			pNodeVisited = node // 修改访问的点
		}else{
			// 根节点再次入栈
			stack.Push(node)
			// 访问右节点,然后右节点相当于根节点
			node = node.right
			for node != nil{
				stack.Push(node)
				node = node.left
			}
		}
	}

}

// 二叉搜索树