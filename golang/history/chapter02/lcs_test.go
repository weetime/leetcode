package chapter02

import (
	"testing"
)

func TestLCS(t *testing.T) {
	var solution func(str1, str2 string) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(str1, str2 string) int {
		m, n := len(str1), len(str2)
		var dp = make([][]int, m+1)
		for i := 0; i <= m; i++ {
			dp[i] = make([]int, n+1)
		}
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if str1[i-1] == str2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i][j-1], dp[i-1][j])
				}
			}
		}

		return dp[m][n]
	}

	t.Log(solution("abcba", "abcbcba"))
}
