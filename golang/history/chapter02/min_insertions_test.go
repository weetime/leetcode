package chapter02

import "testing"

func TestMinInsertions(t *testing.T) {
	var solution func(str string) int // 解决方案
	var min func(x, y int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	solution = func(str string) int {
		n := len(str)
		var dp = make([][]int, n) // 构建tpTable
		for i := 0; i < n; i++ {
			dp[i] = make([]int, n) // 初始化都是0
		}

		for i := n - 2; i >= 0; i-- { // 从下往上，从左往右遍历
			for j := i + 1; j < n; j++ {
				if str[i] == str[j] {
					dp[i][j] = dp[i+1][j-1] // 如果已经是回文了 不做任何操作
				} else {
					dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1 // 如果不是回文 要么在右边插入，要么在左边插入
				}
			}
		}

		return dp[0][n-1]
	}

	t.Log(solution("abcd"))
}
