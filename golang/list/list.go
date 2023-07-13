package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildFromArray(arr []int) *ListNode {
	var head = &ListNode{}
	var tail = head
	for _, v := range arr {
		tail.Next = &ListNode{
			Val: v,
		}
		tail = tail.Next
	}
	return head.Next
}

func buildFromArrayByHead(arr []int) *ListNode {
	var head *ListNode
	for _, v := range arr {
		node := &ListNode{
			Val:  v,
			Next: head,
		}
		head = node
	}
	return head
}
