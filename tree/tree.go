package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func main() {
//	bt := &TreeNode{
//		Val: 1,
//		Left: &TreeNode{
//			Val: 2,
//			Left: &TreeNode{
//				Val:   5,
//				Left:  nil,
//				Right: nil,
//			},
//			Right: &TreeNode{
//				Val: 4,
//				Left: &TreeNode{
//					Val:   6,
//					Left:  nil,
//					Right: nil,
//				},
//				Right: &TreeNode{
//					Val:   7,
//					Left:  nil,
//					Right: nil,
//				},
//			},
//		},
//		Right: &TreeNode{
//			Val: 3,
//			Left: &TreeNode{
//				Val:   8,
//				Left:  nil,
//				Right: nil,
//			},
//			Right: &TreeNode{
//				Val:   9,
//				Left:  nil,
//				Right: nil,
//			},
//		},
//	}
//	before(bt)
//	fmt.Println("")
//	center(bt)
//	fmt.Println("")
//	after(bt)
//	fmt.Println("")
//}

// 前序遍历
func before(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val)
	before(node.Left)
	before(node.Right)
}

// 中序遍历
func center(node *TreeNode) {
	if node == nil {
		return
	}
	center(node.Left)
	fmt.Print(node.Val)
	center(node.Right)
}

// 后序遍历
func after(node *TreeNode) {
	if node == nil {
		return
	}
	after(node.Left)
	after(node.Right)
	fmt.Print(node.Val)
}

// 前序+中序构建二叉树
func BuildBinaryTree(preorder, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}
	nodeVal := preorder[0]
	if len(preorder) == 1 {
		return &TreeNode{Val: nodeVal}
	}
	root := &TreeNode{Val: nodeVal}
	var index int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == nodeVal {
			index = i
		}
	}
	preLeft, preRight := preorder[1:index+1], preorder[index+1:]
	inLeft, inRight := inorder[:index], inorder[index+1:]

	root.Left = BuildBinaryTree(preLeft, inLeft)
	root.Right = BuildBinaryTree(preRight, inRight)
	return root
}

// 中序+后序构造二叉树
func BuildBinaryTreeII(inorder, postorder []int) *TreeNode {
	if len(postorder) < 1 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	if len(postorder) == 1 {
		return root
	}
	var index int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			index = i
		}
	}
	inLeft, inRight := inorder[:index], inorder[index+1:]
	postLeft, postRight := postorder[:len(inLeft)], postorder[len(inLeft):len(postorder)-1]
	root.Left = BuildBinaryTreeII(inLeft, postLeft)
	root.Right = BuildBinaryTreeII(inRight, postRight)
	return root
}
