/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// inorder
func kthSmallest(root *TreeNode, k int) int {
    ans := 0
	curr := 0
	var dfs func(*TreeNode) 
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		
		curr = curr + 1

		if curr == k {
			ans = root.Val
			return
		}
		
		
		dfs(root.Right)
		return
	}

	dfs(root)

	return ans
}
