package listnode

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	// map 临时存储 时间复杂度O(m+n) 空间复杂度O(m)
	var getIntersectionNode func(headA, headB *ListNode) *ListNode
	getIntersectionNode = func(headA, headB *ListNode) *ListNode {
		var tmp = make(map[*ListNode]struct{})
		c1, c2 := headA, headB
		for c1 != nil {
			tmp[c1] = struct{}{}
			c1 = c1.Next
		}
		for c2 != nil {
			if _, ok := tmp[c2]; ok {
				return c2
			}
			c2 = c2.Next
		}
		return nil
	}
	t.Log(getIntersectionNode(GetListNode(), GetListNode()))
}

func TestGetIntersectionNode2(t *testing.T) {
	// 双指针 时间复杂度O(m+n) 空间复杂度O(1)
	var getIntersectionNode func(headA, headB *ListNode) *ListNode
	getIntersectionNode = func(headA, headB *ListNode) *ListNode {
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
	t.Log(getIntersectionNode(GetListNode(), GetListNode()))
}
