package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorder(node *TreeNode, target int) int {
	res := 0
	if node == nil {
		return 0
	}
	if node.Val == target {
		res++
	}
	res += preorder(node.Left, target-node.Val)
	res += preorder(node.Right, target-node.Val)
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	res = preorder(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res

}
