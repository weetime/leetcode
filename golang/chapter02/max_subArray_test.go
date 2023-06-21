package chapter02

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var solution func(nums []int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(nums []int) int {
		var length = len(nums)
		if length == 0 {
			return 0
		}
		var dp = make([]int, length) // 带赋零值
		dp[0] = nums[0]
		res := nums[0]
		for i := 1; i < length; i++ {
			dp[i] = max(nums[i], dp[i-1]+nums[i])
			res = max(res, dp[i])
		}
		return res
	}
	t.Log(solution([]int{-3, 1, 3, -1, 2, -4, 2}))
}

func TestMaxSubArray2(t *testing.T) {
	var solution func(nums []int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(nums []int) int {
		var prev, current, res int       // 上一状态，当前状态，最大值
		prev, res = nums[0], nums[0]     // 初始状态为数组第一个元素
		for i := 1; i < len(nums); i++ { // 状态转移方程，res = max(上一个状态+当前值,当前值)
			current = max(nums[i], prev+nums[i])
			prev = current
			res = max(res, current)
		}
		return res
	}
	t.Log(solution([]int{1}))
}
