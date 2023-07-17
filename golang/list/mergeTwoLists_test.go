package list

import "testing"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
			// tail.Next, l2 = l2, l2.Next 这一行不能替代上面这一行，但是新的golang 1.20版本没问题
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}
	return dummy.Next
}

func TestMergeTwoLists(t *testing.T) {
	res := mergeTwoLists(buildFromArray([]int{1, 3, 5}), buildFromArray([]int{2, 4, 6}))
	t.Log(res)
}
