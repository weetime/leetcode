package list

import "testing"

func reversePrint(head *ListNode) []int {
	var arr []int
	var dfs func(head *ListNode)
	dfs = func(head *ListNode) {
		if head == nil {
			return
		}
		dfs(head.Next)
		arr = append(arr, head.Val)
	}
	dfs(head)
	return arr
}

func TestReversePrint(t *testing.T) {
	res := reversePrint(buildFromArray([]int{1, 2, 3, 4, 5}))
	t.Log(res)
}
