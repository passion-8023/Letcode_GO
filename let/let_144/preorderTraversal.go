package let_144

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var handleFun func(node *TreeNode)
	handleFun = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		handleFun(node.Left)
		handleFun(node.Right)
	}
	handleFun(root)
	return res
}

func preorderTraversal_1(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
