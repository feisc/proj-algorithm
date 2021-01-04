package tree

type BinarySearchTree struct {
	*BinaryTree
	compareFunc func(v, nodeV interface{}) int
}

func NewBinarySearchTree(v interface{}, compareFunc func(v, nodeV interface{}) int) *BinarySearchTree{
	if compareFunc == nil{
		return nil
	}
	return &BinarySearchTree{BinaryTree:NewBinaryTree(v),compareFunc:compareFunc}
}

func (this *BinarySearchTree)Find(v interface{}) *TreeNode{
	node := this.root
	for node != nil{
		compare := this.compareFunc(v,node.val)
		if compare < 0{
			node = node.left
		}else if compare > 0{
			node = node.right
		}else{
			return node
		}
	}
	return nil
}

func (this *BinarySearchTree)Insert(v interface{}) bool{
	node := this.root
	for node != nil{
		compareResult := this.compareFunc(v,node.val)
		if compareResult == 0{
			return false
		}else if compareResult < 0{//插入到左子树
			// 如果左子树为空，则插入
			if node.left == nil{
				node.left = NewTreeNode(v)
				break
			}
			node = node.left
		}else{ // 右子树
			if node.right == nil{
				node.right = NewTreeNode(v)
				break
			}
			node = node.right
		}
	}
	return true
}

func (this *BinarySearchTree)Delete(v interface{}) bool {
	var parentNode *TreeNode = nil
	node := this.root
	for nil != node {
		compareResult := this.compareFunc(v,node.val)
		if compareResult < 0{
			parentNode = node
			node = node.left
		}else if compareResult > 0{
			parentNode = node
			node = node.right
		}else{
			break
		}
	}
	if nil == node{ // 节点不存在
		return false
	}else{
		// 只有一个右节点
		if nil == node.left{
			if node == this.root{
				this.root = node.right
			}else if node == parentNode.left{
				parentNode.left = node.right
				node = nil
			}else{
				parentNode.right = node.right
				node = nil
			}
		}else if nil == node.right{
			// 只有左节点
			if node == this.root{
				this.root = node.left
			}else if node == parentNode.right{
				parentNode.right = node.left
				node = nil
			}else{
				parentNode.left = node.left
				node = nil
			}
		}else{
			// 有左右节点，找到节点右子树的最左节点
			pNode := node.right
			for nil != pNode.left{
				pNode = pNode.left
			}
			// 然后最小的值赋予节点，然后将pNode的值赋值nil
			node.val = pNode.val
			pNode = nil
		}
	}
	return true
}

func (this *BinarySearchTree)Min() *TreeNode{
	node := this.root
	for nil != node.left{
		node = node.left // 遍历到左左边
	}
	return node
}

func (this *BinarySearchTree)Max() *TreeNode{
	node := this.root
	for nil != node.right{
		node = node.right
	}
	return  node
}
