package let_interview_0402

import "github.com/passion-8023/letcodego/data"

func sortedArrayToBST(nums []int) *data.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &data.TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	// 确定树的高度
	high := len(nums) / 2

	// 确定树的根节点
	tree := &data.TreeNode{
		Val:   nums[high],
		Left:  sortedArrayToBST(nums[:high]),
		Right: sortedArrayToBST(nums[high+1:]),
	}
	return tree
}
