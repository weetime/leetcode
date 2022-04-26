package chapter02

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	var solution func(nums []int, target int) int
	solution = func(nums []int, target int) int {
		var sum int
		for _, v := range nums {
			sum += v
		}
		diff := sum - target
		if diff < 0 || diff%2 == 1 {
			return 0
		}

		sum = (sum - target) / 2
		var n = len(nums)
		var dp = make([][]int, n+1)
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, sum+1)
			dp[i][0] = 1
		}
		for i := 1; i <= n; i++ {
			for j := 0; j <= sum; j++ {
				if j >= nums[i-1] {
					dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}

		return dp[n][sum]
	}
	t.Log(solution([]int{1, 1, 1, 1, 1}, 3))
}
