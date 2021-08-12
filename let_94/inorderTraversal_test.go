package let_94

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

//第一种方法：递归（使用全局变量）
var result []int
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result = []int{}
	handleProgram(root)
	return result
}

func handleProgram(node *TreeNode)  {
	if node == nil {
		return
	}
	handleProgram(node.Left)
	result = append(result, node.Val)
	handleProgram(node.Right)
}

//第二种方法：递归（函数变量 ）
func inorderTraversal_2(root *TreeNode) []int {
	var res []int
	var handleFun func(node *TreeNode)
	handleFun = func(node *TreeNode) {
		if node == nil {
			return
		}
		handleFun(node.Left)
		res = append(res, node.Val)
		handleFun(node.Right)
	}
	handleFun(root)
	return res
}

//方法三：迭代
func inorderTraversal_3(root *TreeNode) []int {
	//用链表定义一个栈
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
