package tests_test

import (
	. "algorithms/chapter_15"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GreatestValue", func() {
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

	It("should find greatest value in bst data structure", func() {
		result := GreatestValue(node)

		Expect(result).To(Equal("The Odyssey"))
	})
})
