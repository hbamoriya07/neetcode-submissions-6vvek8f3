/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := height(root.Left)
		right := height(root.Right)

		if left > right {
			return left + 1
		}

		return right + 1
	}

	return height(root)
}

