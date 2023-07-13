package chapter02

import (
	"testing"
)

func TestCanPartition(t *testing.T) {
	var solution func(nums []int) bool
	solution = func(nums []int) bool {
		n, sum := len(nums), 0
		for _, v := range nums {
			sum += v
		}
		if sum%2 != 0 {
			return false
		}
		sum = sum / 2
		var dp = make([][]bool, n+1)
		// base case
		// dp[i][j] 表示前i个数组合是否和为j
		for i := 0; i <= n; i++ {
			dp[i] = make([]bool, sum+1)
			dp[i][0] = true // 背包为空的时候，相等于装满了
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= sum; j++ {
				if j-nums[i-1] < 0 { // 数过大，保持上个状态不变
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
				}
			}
		}
		return dp[n][sum]
	}
	t.Log(solution([]int{1, 5, 11, 5}))
}
