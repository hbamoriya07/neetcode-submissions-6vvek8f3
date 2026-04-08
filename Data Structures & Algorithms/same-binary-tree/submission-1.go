/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var checkSame func(p *TreeNode, q *TreeNode) bool
	checkSame = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil || p.Val != q.Val {
			return false
		}

		return checkSame(p.Left, q.Left) && checkSame(p.Right, q.Right)
	}

	return checkSame(p, q)
}
