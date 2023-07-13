package chapter02

import (
	"testing"
)

func TestPackageO_1(t *testing.T) {
	var wt = []int{2, 1, 3}
	var val = []int{4, 2, 3}
	var solution func(n, w int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(n, w int) int {
		var dp = make([][]int, n+1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, w+1)
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= w; j++ {
				if j-wt[i-1] < 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
				}
			}
		}

		return dp[n][w]
	}

	t.Log(solution(3, 4))
}

func TestRob(t *testing.T) {
	var rob func(nums []int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	rob = func(nums []int) int {
		n := len(nums)
		if n == 0 {
			return 0
		}
		var dp = make([]int, n+1)
		dp[1] = nums[0]           // base case  只有一家
		for i := 2; i <= n; i++ { // 两家以上
			dp[i] = max(dp[i-1], dp[i-2]+nums[i-1]) // 不偷第家 或者 偷第i家 + i-2家的最大值
		}
		return dp[n]
	}
	t.Log(rob([]int{2, 7, 9, 3, 1}))
}

func TestRob2(t *testing.T) {
	var rob func(nums []int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	rob = func(nums []int) int {
		n := len(nums)
		if n == 0 {
			return 0
		}
		if n == 1 {
			return nums[0]
		}
		first, second := nums[0], max(nums[0], nums[1])
		for i := 2; i < n; i++ {
			first, second = second, max(first+nums[i], second)
		}
		return second
	}
	t.Log(rob([]int{2, 7, 9, 3, 1}))
}

func TestRob3(t *testing.T) {
	var rob func(nums []int) int
	var max func(x, y int) int
	_rob := func(nums []int) int {
		first, second := nums[0], max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			first, second = second, max(first+nums[i], second)
		}
		return second
	}
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	rob = func(nums []int) int {
		n := len(nums)
		if n == 0 {
			return 0
		}
		if n == 1 {
			return nums[0]
		}
		if n == 2 {
			return max(nums[0], nums[1])
		}

		return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
	}

	t.Log(rob([]int{1, 2, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestRob4(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	var max func(x, y int) int       // 取最大值
	var dfs func(*TreeNode)          // 递归函数
	var rob func(root *TreeNode) int // 最大收益
	var f = make(map[*TreeNode]int)  // 选择o节点的最大值
	var g = make(map[*TreeNode]int)  // 不选择o节点的最大值

	max = func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	// 后序遍历
	dfs = func(node *TreeNode) {
		// base case
		if node == nil {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		// 做选择
		// 选择o节点，那么就不能选择子节点 f选择，g不选择
		f[node] = node.Val + g[node.Left] + g[node.Right]
		// 不选择o节点，那么子节点可选择或不选择，最两种选择的最大值
		g[node] = max(f[node.Left], g[node.Left]) + max(f[node.Right], g[node.Right])
	}

	rob = func(root *TreeNode) int {
		dfs(root)
		return max(f[root], g[root])
	}

	t.Log(rob(root))
}
