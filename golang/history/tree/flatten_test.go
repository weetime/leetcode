package tree

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	// 题解：这题太经典了 时间复杂度O(n)，空间复杂度O(1)
	// 整个过程就是寻找前驱节点的过程
	// curr,next,pre
	// 每次需要找到next的Right 的最后一个不为空的Right节点 记为pre
	// pre.Right = curr.Right
	// curr.Left = nil
	// curr.Right = next
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
	t.Log(root)
}

func TestFlatten2(t *testing.T) {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	// 题解：时间复杂度O(n)，空间复杂度O(1)
	var list []*TreeNode         // 存储前序遍历的所有节点
	var dfs func(root *TreeNode) // DLR
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	for i := 1; i < len(list); i++ { // 交换前面的
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
	t.Log(root)
}
