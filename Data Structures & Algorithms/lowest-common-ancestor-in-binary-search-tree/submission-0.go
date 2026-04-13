/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
curr := root

    for curr != nil {
        // Both nodes in left subtree
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
        // Both nodes in right subtree
        } else if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
        } else {
            // Split happens here → LCA
            return curr
        }
    }

    return nil
}
