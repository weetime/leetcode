package chapter02

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {
	var solution func(str string) int // 解决方案
	var max func(x, y int) int        // 取最大值
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(str string) int {
		length := len(str)            // 字符串长度
		dp := make([][]int, length)   // dpTable
		for i := 0; i < length; i++ { // 准备基础数据 对角线都是1
			dp[i] = make([]int, length)
			dp[i][i] = 1
		}
		for i := length - 2; i >= 0; i-- { // 倒序遍历(从下到上，从左到右) 从最下方 倒数第二行 最后一列开始填充
			for j := i + 1; j < length; j++ {
				if str[i] == str[j] { // 如果相等 则回文数+2
					dp[i][j] = dp[i+1][j-1] + 2
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1]) // 如果不相等则取前一个元素或后一个元素的最大回文数
				}
			}
		}

		return dp[0][length-1] // 终止条件就是整个字符的最大回文数
	}
	t.Log(solution("cbxab"))
}
