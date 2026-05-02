/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var pathSum func(*TreeNode) int
	pathSum = func(root *TreeNode) int{
		if root == nil {
			return 0
		}

		left := max(0, pathSum(root.Left))
		right := max(0, pathSum(root.Right))

		maxSum = max(maxSum, left+right+root.Val)

		return max(left, right) + root.Val
	}

	pathSum(root)
	return maxSum
}
