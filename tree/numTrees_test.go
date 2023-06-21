package tree

import "testing"

func TestNumTrees(t *testing.T) {
	var n = 3
	// 题解：
	// 二叉树节点从[1...n]中选一个，不重复的二叉搜索树，根节点从1..n中取一个数i
	// 那左侧节点选取1...i，右侧节点选取i+1...n
	// 左侧有x种取法，右侧有y种取法 G[i] 就有 x·y中组合
	// 可以采用动态规划的方式
	// 状态转移方程: dp[i] += dp[j-1] * dp[i-1]
	var dp = make([]int, n+1)
	// base case
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ { // 从2到n
		for j := 1; j <= i; j++ { // 左侧取j，右侧就是i-j
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	t.Log(dp[n])
}

func numTrees(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}
