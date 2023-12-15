package let_lcr_123

import (
	"github.com/passion-8023/letcodego/data"
)

func reverseBookList(head *data.ListNode) []int {
	var result []int

	curr := head

	for curr != nil {
		result = append([]int{curr.Val}, result...)
		curr = curr.Next
	}

	return result
}

// 先翻转链表 然后再循环写入切片
func reverseBookList2(head *data.ListNode) []int {
	curr := head

	var newNode *data.ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = newNode
		newNode = curr
		curr = next
	}

	var resutl []int
	for newNode != nil {
		resutl = append(resutl, newNode.Val)
		newNode = newNode.Next
	}

	return resutl
}

// 先写入切片，然后双指针翻转切片
func reverseBookList3(head *data.ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	// 从两端往中间
	//for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
	//	result[left], result[right] = result[right], result[left]
	//}

	// 从中间往两端
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}
