package chapter02

import (
	"testing"
)

func TestLengthOFLIS(t *testing.T) {
	var solution func([]int) int // 解决方案
	var max func(x, y int) int   // 取最大值
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	solution = func(num []int) int {
		length := len(num)
		var dpTable = make([]int, length) // 动态规划 dptable
		var res int                       // 最长子串长度
		for i := 0; i < len(num); i++ {
			_max := 0
			for j := 0; j < i; j++ { // dp[5] 为dp[1~4]中最小的统计值+1
				if num[i] > num[j] {
					_max = max(_max, dpTable[j])
				}
			}
			dpTable[i] = _max + 1
			res = max(res, dpTable[i])
		}

		return res
	}

	t.Log(solution([]int{1, 4, 3, 4, 2, 3}))
}

func TestBinarySearchLIS(t *testing.T) {
	var solution func([]int) int
solution = func(nums []int) int {
		var piles int
		length := len(nums)
		top := make([]int, length)
		for i := 0; i < length; i++ {
			top = append(top, 0)
		}
		for i := 0; i < length; i++ {
			poker := nums[i]
			left, right := 0, piles
			for left < right {
				mid := left + (right-left)/2
				if top[mid] > poker {
					right = mid
				} else if top[mid] < poker {
					left = mid + 1
				} else if top[mid] == poker {
					right = mid
				}
			}
			if left == piles {
				piles++
			}
			top[left] = poker
		}

		return piles
	}
	t.Log(solution([]int{1, 4, 3, 4, 2, 3}))
}
