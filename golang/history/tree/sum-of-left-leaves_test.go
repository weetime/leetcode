package tree

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	var dfs func(root *TreeNode, isLeft bool) int
	dfs = func(root *TreeNode, isLeft bool) int {
		if root == nil {
			return 0
		}
		if isLeft && root.Left == nil && root.Right == nil {
			return root.Val
		}
		return dfs(root.Left, true) + dfs(root.Right, false)
	}
	t.Log(dfs(root, false))
}
