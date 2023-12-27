package let_160

import "github.com/passion-8023/letcodego/data"

func getIntersectionNode(headA, headB *data.ListNode) *data.ListNode {
	pointMap := make(map[*data.ListNode]struct{})

	currA := headA

	for currA != nil {
		pointMap[currA] = struct{}{}
		currA = currA.Next
	}

	currB := headB
	for currB != nil {
		if _, ok := pointMap[currB]; ok {
			return currB
		}
		currB = currB.Next
	}

	return nil
}
