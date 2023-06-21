package listnode

import "testing"

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
