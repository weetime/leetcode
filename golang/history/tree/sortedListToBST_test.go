package tree

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSortedListToBST(t *testing.T) {
	var head *ListNode
	for i := 9; i >= 0; i-- {
		curr := &ListNode{
			Val: i,
		}
		curr.Next = head
		head = curr
	}

	// 题解：
	// 有序链表 和 有序数组一样，构建BST的时候，采用分治的方式
	// 每次构建树，需要找到中间件的位置，才能高度平衡
	var getMid func(left, right *ListNode) *ListNode    // 获取中间节点
	var buildTree func(left, right *ListNode) *TreeNode // 构建树
	getMid = func(left, right *ListNode) *ListNode {    // 采用快慢指针查找链表的中间件元素
		fast, slow := left, left
		for fast != right && fast.Next != right { // 这里注意查找的结束边界为右指针
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
	buildTree = func(left, right *ListNode) *TreeNode { // 分治
		if left == right {
			return nil
		}
		mid := getMid(left, right)
		tree := &TreeNode{Val: mid.Val}
		tree.Left = buildTree(left, mid)
		tree.Right = buildTree(mid.Next, right) // 这里注意是mid.Next
		return tree
	}
	tree := buildTree(head, nil)
	t.Log(tree)
}

func TestSortedListToBST2(t *testing.T) {
	var head *ListNode
	for i := 9; i >= 0; i-- {
		curr := &ListNode{
			Val: i,
		}
		curr.Next = head
		head = curr
	}
	// 题解:
	// 上面耗时 主要在寻找链表的中间位置，BST中序遍历其实就是一个升序链表
	// 可采用分治 + LDR的方式 构建树
	var getLength func(head *ListNode) int        // 获取链表的长度
	var buildTree func(left, right int) *TreeNode // 构建树
	getLength = func(head *ListNode) int {
		n := 0
		for ; head != nil; head = head.Next {
			n++
		}
		return n
	}
	buildTree = func(left, right int) *TreeNode {
		if left > right { // 注意不是>=，如果=的时候退出，mid就不会执行到了
			return nil
		}
		mid := (left + right) >> 1 // 奇数的时候刚好是中的，偶数的时候是中间位置的前面一个
		tree := &TreeNode{}
		tree.Left = buildTree(left, mid-1) // mid-1
		tree.Val = head.Val
		head = head.Next
		tree.Right = buildTree(mid+1, right) // mid+1
		return tree
	}
	tree := buildTree(0, getLength(head)-1)
	t.Log(tree)
}
