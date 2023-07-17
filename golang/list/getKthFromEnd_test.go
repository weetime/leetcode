package list

import "testing"

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	fast, slow := head, head
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

func TestGetKthFromEnd(t *testing.T) {
	res := getKthFromEnd(buildFromArray([]int{1, 2, 3, 4}), 5)
	t.Log(res)
}
