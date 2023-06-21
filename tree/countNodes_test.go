package tree

import (
	"testing"
)

func TestCountNodes(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := (left + right) >> 1
		return &TreeNode{
			array[mid],
			buildTree(left, mid-1),
			buildTree(mid+1, right),
		}
	}
	root := buildTree(0, len(array)-1)
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return dfs(root.Left) + dfs(root.Right) + 1
	}
	t.Log(dfs(root))
}
