package chapter02

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func TestOddEvenList(t *testing.T) {
	l1 := &ListNode{}
	r := l1
	for i := 0; i < 10; i++ {
		current := &ListNode{
			Val: i,
		}
		r.Next = current
		r = current
	}
	t.Log(l1.Next)
}

func TestOddEvenList2(t *testing.T) {
	l1 := &ListNode{}
	for i := 0; i < 10; i++ {
		current := &ListNode{
			Val: i,
		}
		current.Next = l1.Next
		l1.Next = current
	}
	t.Log(l1.Next)
}
