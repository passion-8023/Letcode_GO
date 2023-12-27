package let_2058

import "github.com/passion-8023/letcodego/data"

func nodesBetweenCriticalPoints(head *data.ListNode) []int {
	min, max := -1, -1

	i := 1 // 记录节点位置

	var tmpPoints []int

	//head.Val > head.Next.Val < head.Next.Next.Val // 局部极小值点
	//head.Val < head.Next.Val > head.Next.Next.Val // 局部极大值点

	curr := head
	for curr.Next != nil && curr.Next.Next != nil {
		i++
		// 判断Next节点是否是极值节点
		if curr.Next.Val < curr.Val && curr.Next.Val < curr.Next.Next.Val {
			// 极小
			tmpPoints = append(tmpPoints, i)
		} else if curr.Next.Val > curr.Val && curr.Next.Val > curr.Next.Next.Val {
			// 极大
			tmpPoints = append(tmpPoints, i)
		}

		curr = curr.Next
	}

	l := len(tmpPoints)
	if l > 1 {
		min = tmpPoints[1] - tmpPoints[0]
		for i := 1; i < l-1; i++ {
			val := tmpPoints[i+1] - tmpPoints[i]
			if min > val {
				min = val
			}
		}
		max = tmpPoints[l-1] - tmpPoints[0]
	}

	return []int{min, max}
}
