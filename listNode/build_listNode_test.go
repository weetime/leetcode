package listnode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTail(t *testing.T) {
	l1 := &ListNode{}
	r := l1
	for i := 0; i < 2; i++ {
		current := &ListNode{
			Val: i,
		}
		r.Next = current
		r = current
	}
	t.Log(l1.Next)
}

func GetListNode() *ListNode {
	node := &ListNode{}
	tail := node
	for i := 0; i < 10; i++ {
		current := &ListNode{
			Val: i,
		}
		tail.Next = current
		tail = current
	}

	return node.Next
}

func TestAddHead(t *testing.T) {
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

func TestMe(t *testing.T) {
	l1 := &ListNode{}
	first := l1

	current := &ListNode{
		Val: 1,
	}
	first.Next = current
	first = first.Next

	current = &ListNode{
		Val: 2,
	}
	first.Next = current
	first = first.Next
	t.Log(l1)
}

func TestBuildTail(t *testing.T) {
	var head = &ListNode{}
	curr := head
	for i := 0; i < 10; i++ {
		new := &ListNode{
			Val: i,
		}

		curr.Next, curr = new, new
	}

	t.Log(head.Next)
}

func TestBuildHead(t *testing.T) {
	var head *ListNode
	for i := 0; i < 10; i++ {
		new := &ListNode{
			Val: i,
		}
		new.Next = head
		head = new
	}

	curr := head
	var res *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next, res = res, curr
		curr = next
	}
	t.Log(head)
}
