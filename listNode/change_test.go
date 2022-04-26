package listnode

import "testing"

func TestChange(t *testing.T) {
	var change func(amount int, coins []int) int // 函数签名
	change = func(amount int, coins []int) int {
		n := len(coins)             // 总硬币数
		var dp = make([][]int, n+1) // dp[0][j]=0 表示0个硬币 达到j金额的 组合数 当然是0
		for i := 0; i <= n; i++ {
			dp[i] = make([]int, amount+1) // i=0 表示0个硬币 j=0 表示目标金额为0
			// base case
			dp[i][0] = 1 // i种硬币 总金额为0 当然有1种选择(不取硬币即可)
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= amount; j++ {
				if j-coins[i-1] >= 0 {
					// dp[i][j] i种硬币达到目标为j的选择
					// 不放i 就保持上个状态，取i 就看放了i后，剩余金额是否能放满
					dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
				} else {
					// j为总金额 当减去第coins[i-1]小于0，当然当前不放硬币
					dp[i][j] = dp[i-1][j]
				}
			}
		}

		return dp[n][amount]
	}
	t.Log(change(5, []int{1, 2, 5}))
}
