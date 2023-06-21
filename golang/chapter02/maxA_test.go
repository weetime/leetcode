package chapter02

import "testing"

func TestMaxA(t *testing.T) {
	var solution func(n int) int
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(n int) int {
		var dp = make([]int, n+1)
		for i := 1; i <= n; i++ {
			dp[i] = i
			for j := 2; j < i; j++ {
				dp[i] = max(dp[i], dp[j-2]*(i-j+1))
			}
		}
		return dp[n]
	}
	t.Log(solution(9))
}

func TestTwoSum(t *testing.T) {
	var solution func(nums []int, target int) []int
	solution = func(nums []int, target int) []int {
		var tmp = make(map[int]int)
		for k, v := range nums {
			if index, ok := tmp[v]; ok {
				return []int{index, k}
			}
			tmp[target-v] = k
		}
		return []int{}
	}

	t.Log(solution([]int{2, 5, 7, 8, 9}, 9))
}
