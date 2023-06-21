package tree

import (
	"math"
	"testing"
)

func TestLargestValues(t *testing.T) {
	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var buildTree func(Left, right int) *TreeNode
	buildTree = func(Left, right int) *TreeNode {
		if Left > right {
			return nil
		}
		mid := (Left + right) >> 1
		root := &TreeNode{array[mid], nil, nil}
		root.Left = buildTree(Left, mid-1)
		root.Right = buildTree(mid+1, right)
		return root
	}
	root := buildTree(0, len(array)-1)
	var bfs func(root *TreeNode) []int
	bfs = func(root *TreeNode) []int {
		var res []int
		if root == nil {
			return res
		}
		var queue = []*TreeNode{root}
		var vals []int
		for level := 0; len(queue) > 0; level++ {
			var ans = math.MinInt
			q := queue
			queue = nil
			for _, node := range q {
				if ans < node.Val {
					ans = node.Val
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			vals = append(vals, ans)
		}
		return vals
	}
	t.Log(bfs(root))
}
