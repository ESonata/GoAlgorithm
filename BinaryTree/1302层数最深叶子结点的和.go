package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var dfs func(node *TreeNode, height int)
	total := 0
	maxHeight := -1
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if height > maxHeight {
			maxHeight = height
			total = node.Val
		} else if height == maxHeight {
			total += node.Val
		}
		dfs(node.Left, height+1)
		dfs(node.Right, height+1)
	}
	dfs(root, 0)
	return total

}
func main() {

}
