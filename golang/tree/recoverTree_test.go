package tree

import "testing"

func TestRecoverTree(t *testing.T) {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var pre, x, y *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != nil && root.Val < pre.Val {
			// [1,3,2,4,5,6,7] 3,2 交换 只有一处 a[i] > a[i+1]
			// [1,6,3,4,2,5,7] 6,2 交换  有两处 a[i] > a[i+1] 和 a[j] > a[j+1]
			y = root // 可能是相邻 也可能不相邻
			if x == nil {
				x = pre
			}
		}
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	x.Val, y.Val = y.Val, x.Val // 交换位置
	t.Log(root)
}
