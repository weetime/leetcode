package tree

import (
	"math"
	"testing"
)

func TestMinDiffInBST(t *testing.T) {
	// [49,52,60,89,90]
	root := &TreeNode{ // 初始化一棵树
		Val: 90,
		Left: &TreeNode{
			Val: 60,
			Left: &TreeNode{
				Val: 49,
				Right: &TreeNode{
					Val: 52,
				},
			},
			Right: &TreeNode{
				Val: 89,
			},
		},
	}
	var minDiffInBST func(root *TreeNode) int
	// 题解 关键部分是 升序数组中最小差值一定是两个相邻数字的差值
	// 我们知道采用中序遍历可以得到BST的升序数组
	// 用一个pre记录前面一个节点的Val，如果和当前节点的Val差值小于ans，则更新ans
	var ans, pre = math.MaxInt, -1 // 初始化ans，pre
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil { // 当递归完成后退出
			return
		}

		// LDR 升序
		dfs(root.Left)
		// 处理逻辑 更新ans的值
		if pre != -1 && ans > root.Val-pre {
			ans = root.Val - pre
		}
		// 记录上一个节点的Val
		pre = root.Val
		dfs(root.Right)
	}
	minDiffInBST = func(root *TreeNode) int {
		dfs(root)
		if ans == math.MaxInt {
			return 0
		}
		return ans
	}
	t.Log(minDiffInBST(root))
}
