/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var DFS func(node *TreeNode) int

	DFS = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := DFS(node.Left)
		if leftHeight == -1 {
			return -1
		}

		rightHeight := DFS(node.Right)
		if rightHeight == -1 {
			return -1
		}

		if math.Abs(float64(leftHeight-rightHeight)) > 1 {
			return -1
		}

		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	return DFS(root) != -1
}
