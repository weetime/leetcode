package listnode

import "testing"

func TestIsPalindrome(t *testing.T) {
	var head *ListNode
	var array = []int{1, 2, 3, 2, 1}
	for _, item := range array {
		new := &ListNode{item, head}
		head = new
	}

	var reverse func(node *ListNode) bool
	reverse = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		res := reverse(node.Next)
		res = res && head.Val == node.Val
		head = head.Next
		return res
	}
	t.Log(reverse(head))
}
