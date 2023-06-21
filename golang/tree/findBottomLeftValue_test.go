package tree

import (
	"math"
	"testing"
)

func TestFindBottomLeftValue(t *testing.T) {
	var root = &TreeNode{Val: 1}
	var maxDept = math.MinInt
	var res int
	var findBottomLeftValue func(root *TreeNode) int
	var dfs func(root, prev *TreeNode, depth int)
	dfs = func(root, prev *TreeNode, depth int) {
		if root == nil {
			return
		}
		if maxDept < depth{
			maxDept = depth
			res = root.Val
		}
		dfs(root.Left, root, depth+1)
		dfs(root.Right, root, depth+1)
	}
	findBottomLeftValue = func(root *TreeNode) int {
		res = root.Val
		dfs(root, root, 0)
		return res
	}
	t.Log(findBottomLeftValue(root))
}
