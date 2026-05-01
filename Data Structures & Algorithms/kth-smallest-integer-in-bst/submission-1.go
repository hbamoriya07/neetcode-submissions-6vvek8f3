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
   result := 0

	var inorder func(*TreeNode) 
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		inorder(node.Left)
		k--
		if k == 0 {
			result = node.Val
			return
		}
		inorder(node.Right)

	}
	inorder(root)

	return result
}
