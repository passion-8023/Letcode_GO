package let_145

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

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var handleFun func(node *TreeNode)
	handleFun = func(node *TreeNode) {
		if node == nil {
			return
		}
		handleFun(node.Left)
		handleFun(node.Right)
		res = append(res, node.Val)
	}
	handleFun(root)
	return res
}

func postorderTraversal_1(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
