package tree

import "testing"

func verifyPostorder(postorder []int) bool {
	var res []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		res = append(res, root.Val)
	}
	if len(res) != len(postorder) {
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] != postorder[i] {
			return false
		}
	}

	return true
}

func TestMain(t *testing.T) {
	res := verifyPostorder([]int{1, 3, 2, 6, 5})
	t.Log(res)
}
