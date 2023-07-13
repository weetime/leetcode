package tree

import (
	"testing"
)

func TestDiameterOfBinaryTree2(t *testing.T) {
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
		},
	}
	var ans = 1
	var solution func(root *TreeNode) int
	var max func(x, y int) int
	var dfs func(root *TreeNode) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := dfs(root.Left)
		r := dfs(root.Right)
		ans = max(ans, l+r+1)
		dep := max(l, r) + 1
		return dep
	}
	solution = func(root *TreeNode) int {
		dfs(root)
		return ans - 1
	}
	t.Log(solution(root))
}

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{ // 初始化一棵树
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 6,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var ans int                           // 最大直径 默认是0
	var solution func(root *TreeNode) int // 解决方案呢
	var max func(x, y int) int            // 取最大值
	var dfs func(root *TreeNode) int      // 深度遍历
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dfs = func(root *TreeNode) int {
		if root == nil { // 递归终止条件
			return 0
		}
		l := dfs(root.Left)   // 左侧高度
		r := dfs(root.Right)  // 右侧高度
		dept := max(l, r) + 1 // 数的最大高度
		ans = max(ans, l+r)   // 数的最大直径 = 左侧最大高度 + 右侧的最大高度

		return dept
	}
	solution = func(root *TreeNode) int {
		dfs(root)
		return ans
	}
	t.Log(solution(root))
}
