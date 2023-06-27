package listnode

import (
	"testing"
)

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
	// tail.Next = &ListNode{
	// 	Val: 2,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }

	return node.Next
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

func TestHead(t *testing.T) {
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

func TestBuildNode(t *testing.T) {
	var head *ListNode
	for i := 0; i < 10; i++ {
		new := &ListNode{i, nil}
		new.Next = head
		head = new
	}

	t.Log(head)

	var ln = &ListNode{}
	current := ln
	for i := 0; i < 10; i++ {
		new := &ListNode{i, nil}
		current.Next = new
		current = new
	}
	t.Log(ln)

	var reverseLn func(ln *ListNode) *ListNode
	reverseLn = func(ln *ListNode) *ListNode {
		if ln == nil || ln.Next == nil {
			return ln
		}
		// ln 8->9; last:9->nil => ln 8->nil; last:9->8->nil
		// ln: 7->8; last:9->8->nil => ln 7->nil;last:9->8->7->nil
		last := reverseLn(ln.Next)
		ln.Next.Next = ln
		ln.Next = nil
		return last
	}
	l := reverseLn(ln.Next)
	t.Log(l)
}

func TestAddHead(t *testing.T) {
	var head = &ListNode{}
	// for i := 0; i < 10; i++ {
	// 	tmp := &ListNode{i, nil}
	// 	tmp.Next = head.Next
	// 	head.Next = tmp
	// }
	for i := 0; i < 10; i++ {
		head.Next = &ListNode{i, head.Next}
	}
	t.Log(head.Next)
}

// 极限头插法 构建链表
func TestAddHead2(t *testing.T) {
	// 头插法构建链表 9->8->7->6->5->4->3->2->1->0
	var head *ListNode
	for i := 0; i < 10; i++ {
		head = &ListNode{i, head}
	}

	// 迭代法 反转
	var new *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		new, curr.Next = curr, new
		curr = next
	}

	t.Log(new)

	// 后序 递归  反转链表 0—>(nil<-1<-2<-3...<-9) head=>0,ln=>9
	var reverse func(head *ListNode) *ListNode
	reverse = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		defer func() {
			head.Next.Next, head.Next = head, nil
		}()
		return reverse(head.Next)
	}

	last := reverse(new)
	t.Log(last)
}
