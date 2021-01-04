package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T){
	binaryTree := NewBinaryTree(5)
	binaryTree.root.left = NewTreeNode(6)
	binaryTree.root.right = NewTreeNode(7)
	binaryTree.root.left.left = NewTreeNode(8)
	binaryTree.root.left.right = NewTreeNode(9)
	binaryTree.root.right.left = NewTreeNode(10)

	//binaryTree.PreOrderTraverse(binaryTree.root)
	//binaryTree.InOrderTraverse(binaryTree.root)
	//binaryTree.PostOrderTraverse(binaryTree.root)
	//binaryTree.InOrderTraverseRec()
	//fmt.Println("----------")
	//binaryTree.InOrderTraverse(binaryTree.root)

	//binaryTree.InOrderTraverse(binaryTree.root)
	//fmt.Println("-----------")
	//binaryTree.InOrderTraverseRec()

	binaryTree.PostOrderTraverse(binaryTree.root)
	fmt.Println("----------")
	binaryTree.PostOrderTraverseRec()
}

