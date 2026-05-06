/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var result []string

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			result = append(result, "null")
			return 
		}

		result = append(result, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    value := strings.Split(data, ",")
	index := 0 

	var build func() *TreeNode
	build = func() *TreeNode {
		if value[index] == "null" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(value[index])
		node := &TreeNode{Val: val}
		index++

		node.Left = build()
		node.Right = build()

		return node
	}

	return build()
}
