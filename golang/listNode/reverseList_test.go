package listnode

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var solution func(head *ListNode) *ListNode
	solution = func(head *ListNode) *ListNode {
		new := &ListNode{} // 重新初始化新链表
		for head != nil {
			current := &ListNode{ // 新增一个节点
				Val:  head.Val,
				Next: nil,
			}
			// 0->1->2
			// 0->2 => 1->0
			current.Next = new.Next
			new.Next = current
			head = head.Next
		}
		return new.Next
	}
	t.Log(solution(GetListNode()))
}

func TestReverseList2(t *testing.T) {
	var solution func(head *ListNode) *ListNode
	solution = func(head *ListNode) *ListNode {
		var new *ListNode
		current := head
		for current != nil {
			next := current.Next
			current.Next, new = new, current
			current = next
		}
		return new
	}
	t.Log(solution(GetListNode()))
}

func TestReverse(t *testing.T) {
	var solution func(*ListNode) *ListNode
	solution = func(ln *ListNode) *ListNode {
		var new *ListNode
		current := ln
		for current != nil {
			next := current.Next
			// 0->1;0->nil;new=0->nil
			// 1->2;1->0;new=1->0->nil
			// 2->3;2->1;new=2->1->0->nil
			current.Next, new = new, current
			current = next
		}

		return new
	}
	t.Log(solution(GetListNode()))
}

func TestReverse2(t *testing.T) {
	head := GetListNode()
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		last := dfs(head.Next)
		head.Next.Next = head
		head.Next = nil
		return last
	}
	l := dfs(head)
	t.Log(l)
}

func TestMergeList(t *testing.T) {
	l1 := GetListNode()
	l2 := GetListNode()
	var curr = &ListNode{}
	var new = curr
	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			break
		} else if l2 == nil {
			curr.Next = l1
		}
		if l1 != nil && l2 != nil && l1.Val >= l2.Val {
			// println("This is v2", l2.Val)
			curr.Next = &ListNode{
				Val: l2.Val,
			}
			curr = curr.Next
			l2 = l2.Next
			// curr.Next = nil
		} else {
			// println("This is v1", l1.Val)
			curr.Next = &ListNode{
				Val: l1.Val,
			}
			curr = curr.Next
			l1 = l1.Next
		}
	}

	println("This is list node", new.Next)
}

func TestMergeList2(t *testing.T) {
	l1, l2 := GetListNode(), GetListNode()
	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(ln1, ln2 *ListNode) *ListNode {
		if ln1 == nil {
			return ln2
		}
		if ln2 == nil {
			return ln1
		}
		if ln1.Val < ln2.Val {
			ln1.Next = merge(ln1.Next, ln2)
			return ln1
		} else {
			ln2.Next = merge(ln1, ln2.Next)
			return ln2
		}
	}
	res := merge(l1, l2)
	println(res)
}

func TestDeleteList(t *testing.T) {
	l1 := GetListNode()
	curr := l1
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}

	}
	println(l1)
}

func TestHasCycle(t *testing.T) {
	head := GetListNode()
	var hasCycle func(head *ListNode) bool
	hasCycle = func(head *ListNode) bool {
		fast, slow := head, head
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		}
		return false
	}
	res := hasCycle(head)
	println(res)
}

func TestGetIntersectionNode3(t *testing.T) {
	h1, h2 := GetListNode(), GetListNode()
	var getIntersectionNode func(headA, headB *ListNode) *ListNode
	getIntersectionNode = func(headA, headB *ListNode) *ListNode {
		if headA == nil || headB == nil {
			return nil
		}
		n1 := headA
		n2 := headB
		for n1 != nil || n2 != nil {
			if n1 == n2 {
				return n1
			}
			if n1 == nil {
				n1 = headB
			} else {
				n1 = n1.Next
			}
			if n2 == nil {
				n2 = headA
			} else {
				n2 = n2.Next
			}
		}
		return nil
	}
	res := getIntersectionNode(h1, h2)
	println(res)
}

func TestRemove_elements(t *testing.T) {
	var remove_elements func(head *ListNode, val int) *ListNode
	head := GetListNode()
	remove_elements = func(head *ListNode, val int) *ListNode {
		if head == nil {
			return head
		}

		sentinel := &ListNode{Next: head}
		curr, prev := head, sentinel
		for curr != nil {
			if curr.Val == val {
				prev.Next = curr.Next
			} else {
				prev = curr
			}
			curr = curr.Next
		}
		return sentinel.Next
	}
	res := remove_elements(head, 1)
	println("{:?}", res)

	// sentnal := &ListNode{}
	// sentnal.Next = head
	// prev, curr := sentnal, head
	// for curr != nil {
	// 		if curr.Val == val {
	// 				prev.Next = curr.Next
	// 		} else {
	// 				prev = curr
	// 		}
	// 		curr = curr.Next
	// }
	// return sentnal.Next
}
