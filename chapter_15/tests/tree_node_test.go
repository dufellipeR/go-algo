package tests_test

import (
	. "algorithms/chapter_15"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("TreeNode", func() {
	var node *TreeNode

	BeforeEach(func() {
		nodeVIII := TreeNode{Value: "Alice in Wonderland"}
		nodeVI := TreeNode{Value: "Lord of the flies"}
		nodeIV := TreeNode{Value: "The Odyssey"}
		nodeV := TreeNode{Value: "Great Expectations", LeftChild: &nodeVIII, RightChild: &nodeVI}
		nodeIII := TreeNode{Value: "Pride and Prejudice"}
		nodeII := TreeNode{Value: "Robinson Crusoe", LeftChild: &nodeIII, RightChild: &nodeIV}
		nodeI := TreeNode{Value: "Moby Dick", RightChild: &nodeII, LeftChild: &nodeV}

		node = &nodeI
	})

	It("should print in traversal preorder", func() {

		TraverseAndPrintPreOrder(node)
	})

	It("should print in traversal post order", func() {

		TraverseAndPrintPostOrder(node)
	})
})
