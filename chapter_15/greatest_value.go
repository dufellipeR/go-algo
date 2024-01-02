package chapter_15

func GreatestValue(node *TreeNode) any {
	if node.RightChild == nil {
		return node.Value
	}

	return GreatestValue(node.RightChild)
}
