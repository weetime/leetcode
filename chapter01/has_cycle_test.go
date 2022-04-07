package chapter01

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func cycleNode() *ListNode {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	node.Next.Next.Next.Next = node.Next

	return node
}

func TestHasCycle(t *testing.T) {
	var solution func(*ListNode) bool
	solution = func(ln *ListNode) bool {
		faster, slow := ln, ln
		for faster != nil && faster.Next != nil {
			faster = faster.Next.Next
			slow = slow.Next
			if slow == faster {
				return true
			}
		}
		return false
	}
	res := solution(cycleNode())
	t.Log(res)
}

func TestCycleHeader(t *testing.T) {
	var solution func(*ListNode) *ListNode
	solution = func(ln *ListNode) *ListNode {
		fast, slow := ln, ln
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if slow == fast {
				break
			}
		}

		slow = ln
		for slow != fast {
			fast = fast.Next
			slow = slow.Next
		}

		return slow
	}
	t.Log(solution(cycleNode()))
}

func TestMiddle(t *testing.T) {
	var solution func(*ListNode) *ListNode
	solution = func(ln *ListNode) *ListNode {
		slow, fast := ln, ln
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		return slow
	}

	t.Log(solution(cycleNode()))
}
