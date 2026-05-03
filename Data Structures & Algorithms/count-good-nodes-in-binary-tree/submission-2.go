/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func goodNodes(root *TreeNode) int {
    return dfs(root, root.Val)
}

func dfs(root *TreeNode, maxSoFar int) int {
    if root == nil {
        return 0
    }

    count := 0 

    if root.Val >= maxSoFar {
        count = 1
        maxSoFar = root.Val
    }

    count += dfs(root.Left, maxSoFar)
    count += dfs(root.Right, maxSoFar)

    return count
}
