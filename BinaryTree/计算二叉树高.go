package main

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
