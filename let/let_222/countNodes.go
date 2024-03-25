package let_222

import (
	"fmt"
	"github.com/passion-8023/letcodego/data"
	"math"
)

type TreeNode = data.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes_2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 获取层
	level := 0
	for curr := root; curr.Left != nil; curr = curr.Left {
		level++
	}

	min := math.Pow(2, float64(level))
	max := math.Pow(2, float64(level+1)) - 1

	for min < max {
		mid := math.Ceil((min + max) / 2)
		if checkExists(root, mid) {
			min = mid
		} else {
			max = mid - 1
		}
	}
	return int(min)
}

func checkExists(root *TreeNode, k float64) bool {
	curr := root
	bits := fmt.Sprintf("%b", int(k))

	for i := 1; i < len(bits); i++ {
		if curr == nil {
			break
		}
		if bits[i] == '0' {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	if curr != nil {
		return true
	}
	return false
}
