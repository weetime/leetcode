package chapter02

import (
	"math"
	"strconv"
	"testing"
)

func TestSuperEggDrop(t *testing.T) {
	var solution func(K, N int) int
	var key func(K, N int) string
	var dp = make(map[string]int)
	var min func(x, y int) int
	// var max func(x, y int) int
	key = func(K, N int) string {
		return strconv.Itoa(K) + "_" + strconv.Itoa(N)
	}

	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	// max = func(x, y int) int {
	// 	if x > y {
	// 		return x
	// 	}
	// 	return y
	// }

	solution = func(K, N int) int {
		if K == 1 {
			return N
		}
		if N == 0 {
			return 0
		}
		if v, ok := dp[key(K, N)]; ok {
			return v
		}
		res := math.MaxInt
		// for i := 1; i <= N; i++ {
		// 	res = min(res, max(solution(K, N-i), solution(K-1, i-1))+1)
		// }
		lo, hi := 1, N
		for lo <= hi {
			mid := (lo + hi) / 2
			broker := solution(K-1, mid-1)
			not_broker := solution(K, N-mid)
			if broker > not_broker {
				hi = mid - 1
				res = min(res, broker+1)
			} else {
				lo = mid + 1
				res = min(res, not_broker+1)
			}
		}
		dp[key(K, N)] = res
		return res
	}
	t.Log(solution(4, 5000))
}
