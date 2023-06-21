package tree

import "testing"

func TestMumColor(t *testing.T) {
	root := &TreeNode{ // 初始化树
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	var hashMap = make(map[int]struct{})
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		hashMap[root.Val] = struct{}{}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	t.Log(len(hashMap))
}
