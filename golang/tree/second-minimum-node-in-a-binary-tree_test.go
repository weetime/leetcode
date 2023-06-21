package tree

import (
	"testing"
)

func TestFindSecondMinimumValue(t *testing.T) {
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
			Val: 2,
		},
	}
	var ans = -1
	var res = root.Val
	var solution func(root *TreeNode) int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || (ans != -1 && root.Val >= ans) {
			return
		}
		if res < root.Val {
			if root.Left != nil && root.Right != nil {
				if root.Left.Val > root.Right.Val {
					ans = root.Right.Val
				} else {
					ans = root.Left.Val
				}
			} else if root.Left != nil {
				ans = root.Left.Val
			} else if root.Right != nil {
				ans = root.Right.Val
			}
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	solution = func(root *TreeNode) int {
		dfs(root)
		return ans
	}

	t.Log(solution(root))
}

func TestFindSecondMinimumValue2(t *testing.T) {
	// 根节点是最小的，后面再取一个次小的即可
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
			Val: 2,
		},
	}
	var ans = -1
	var solution func(root *TreeNode) int
	var dfs func(root *TreeNode, res int)
	dfs = func(root *TreeNode, res int) {
		// 如果遍历结束，或者ans已经取到第二小的值了，并且其他节点大于小于ans 就返回ans
		if root == nil || (ans != -1 && root.Val >= ans) {
			return
		}
		// 取第二小
		if res < root.Val {
			ans = root.Val
		}
		dfs(root.Left, res)
		dfs(root.Right, res)
	}
	solution = func(root *TreeNode) int {
		dfs(root, root.Val)
		return ans
	}

	t.Log(solution(root))
}
