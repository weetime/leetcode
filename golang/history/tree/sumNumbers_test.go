package tree

import (
	"math"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	var arrar = []int{1, 2, 3, 4}
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) >> 1
		tree := &TreeNode{Val: arrar[mid]}
		tree.Left = buildTree(left, mid-1)
		tree.Right = buildTree(mid+1, right)
		return tree
	}
	var dfs func(root *TreeNode)
	var path []int
	var ans int
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		defer func() {
			path = path[:len(path)-1]
		}()
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			for i, n := 0, len(path); i < n; i++ {
				ans += int(math.Pow10(n-1-i)) * path[i]
			}
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	root := buildTree(0, len(arrar)-1)
	dfs(root)
	t.Log(ans)
}

func TestSumNumbers2(t *testing.T) {
	var arrar = []int{1, 2, 3, 4}
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) >> 1
		tree := &TreeNode{Val: arrar[mid]}
		tree.Left = buildTree(left, mid-1)
		tree.Right = buildTree(mid+1, right)
		return tree
	}
	var dfs func(root *TreeNode, prevSum int) int
	dfs = func(root *TreeNode, prevSum int) int {
		if root == nil {
			return 0
		}
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	root := buildTree(0, len(arrar)-1)
	t.Log(dfs(root, 0))
}
