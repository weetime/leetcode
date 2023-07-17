package list

import "testing"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

func TestGetIntersectionNode(t *testing.T) {
	res := getIntersectionNode(buildFromArray([]int{1, 2, 3}), buildFromArray([]int{3, 4, 5}))
	t.Log(res)
}
