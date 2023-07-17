package list

import "testing"

/**
* 解题思路，比较简单，主要是要注意一下，可能需要删除第一个节点，或者最后一个节点
* 所以需要引入dummy，为了保证dummy指针不变，需要再引入一个指针操作dummy
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	tail := dummy
	for tail.Next != nil {
		if tail.Next.Val == val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}

	return dummy.Next
}

func TestDeleteNode(t *testing.T) {
	res := deleteNode(buildFromArray([]int{1, 2, 3, 4, 5}), 5)
	t.Log(res)
}
