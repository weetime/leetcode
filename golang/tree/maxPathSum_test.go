package tree

import (
	"math"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	var root = &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: -10,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// 题解：
	// ans 存储最大路径和
	// 最大路径和，前提是必现能连成一片
	// 左侧 + 右侧 + root.val 即为最大值
	// 最大贡献 = 左侧最大贡献 || 右侧最大贡献 + root.Val
	var ans = math.MinInt            // 初始最小值
	var dfs func(root *TreeNode) int // 采用LRD递归
	var max func(x, y int) int       // 取最大值
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
		l := max(0, dfs(root.Left))  // 左节点最大和
		r := max(0, dfs(root.Right)) // 右节点最大和
		ans = max(ans, l+r+root.Val) // 如果加上根的值小于上一次的ans，则保持不变
		return max(l, r) + root.Val  // 返回左节点 或 右节点 加 root 为了能连成一条线
	}
	a := dfs(root)
	t.Log(a, ans)
}
