package chapter01

import (
	"math"
	"testing"
)

// dp(10) = dp(5)+dp(5)+dp(0)
// dp(10) = 2
func TestCoins(t *testing.T) {
	coins := []int{1, 2, 5}
	var dp func(int) float64
	dp = func(n int) float64 {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxFloat64
		for _, coin := range coins {
			subProblem := dp(n - coin) // subProblem 0->1->2
			if subProblem == -1 {
				continue
			}
			res = math.Min(res, subProblem+1)
		}

		if res == math.MaxFloat64 {
			return -1
		}

		return res
	}

	t.Log(dp(10))
}

func TestCoins2(t *testing.T) {
	memo := map[int]float64{}
	coins := []int{9, 5}
	var dp func(int) float64
	dp = func(n int) float64 {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxFloat64
		for _, coin := range coins {
			if _, ok := memo[n-coin]; !ok {
				memo[n-coin] = dp(n - coin)
			} else {
				// 这里会剪枝
				// fmt.Println(n - coin)
			}
			if memo[n-coin] == -1 {
				continue
			}
			res = math.Min(res, memo[n-coin]+1)
		}

		if res == math.MaxFloat64 {
			return -1
		}

		return res
	}

	t.Log(dp(19))
}

func TestCoins3(t *testing.T) {
	memo := map[int]float64{}
	coins := []int{9, 5}
	var dp func(int) float64
	dp = func(n int) float64 {
		if _, ok := memo[n]; ok {
			return memo[n]
		}
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		res := math.MaxFloat64
		for _, coin := range coins {
			subProblem := dp(n - coin)
			if subProblem == -1 {
				continue
			}
			res = math.Min(res, subProblem+1)
		}

		if res == math.MaxFloat64 {
			memo[n] = -1
		} else {
			memo[n] = res
		}
		// t.Log(memo)
		return memo[n]
	}

	t.Log(dp(20))
}
