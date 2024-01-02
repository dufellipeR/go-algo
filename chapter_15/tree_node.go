package chapter_15

import "fmt"

type TreeNode struct {
	Value      string
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func Search(searchValue string, node *TreeNode) any {
	if node == nil {
		return nil
	}

	if node.Value == searchValue {
		return node
	} else if searchValue < node.Value {
		return Search(searchValue, node.LeftChild)
	} else if searchValue > node.Value {
		return Search(searchValue, node.RightChild)
	}
	return 0
}

func Insert(value string, node *TreeNode) {
	if value < node.Value {
		if node.LeftChild == nil {
			node.LeftChild = &TreeNode{Value: value}
		} else {
			Insert(value, node.LeftChild)
		}
	} else if value > node.Value {
		if node.RightChild == nil {
			node.RightChild = &TreeNode{Value: value}
		} else {
			Insert(value, node.RightChild)
		}
	}
}

func Delete(value string, node *TreeNode) *TreeNode {
	// This base is when we've hit the bottom of the tree
	// and the parent node has no children
	if node == nil {
		return nil
	} else if value < node.Value {
		node.LeftChild = Delete(value, node.LeftChild)
		return node
	} else if value > node.Value {
		node.RightChild = Delete(value, node.RightChild)
		return node
	} else if value == node.Value {
		if node.LeftChild == nil {
			return node.RightChild
		} else if node.RightChild == nil {
			return node.LeftChild
		} else {
			node.RightChild = lift(node.RightChild, node)
			return node
		}
	}
	return nil
}

func lift(node *TreeNode, nodeToDelete *TreeNode) *TreeNode {
	if node.LeftChild != nil {
		node.LeftChild = lift(node.LeftChild, nodeToDelete)
		return node
	}
	return nil
}

func TraverseAndPrintInOrder(node *TreeNode) {
	if node == nil {
		return
	}

	TraverseAndPrintInOrder(node.LeftChild)
	fmt.Println(node.Value)
	TraverseAndPrintInOrder(node.RightChild)
}

func TraverseAndPrintPreOrder(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)

	TraverseAndPrintInOrder(node.LeftChild)
	TraverseAndPrintInOrder(node.RightChild)
}

func TraverseAndPrintPostOrder(node *TreeNode) {
	if node == nil {
		return
	}

	TraverseAndPrintInOrder(node.LeftChild)
	TraverseAndPrintInOrder(node.RightChild)

	fmt.Println(node.Value)
}
