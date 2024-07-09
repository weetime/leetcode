package list

import (
	"testing"
)

func a(nums []int) int {
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var len = len(nums)
	if len == 1 {
		return nums[0]
	}
	var dp = make([]int, len)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len-1]
}

func b(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var rober = func(nums []int) int {
		prev, curr := 0, 0
		for _, num := range nums {
			prev, curr = curr, max(curr, prev+num)
		}
		return curr
	}

	return max(rober(nums[1:]), rober(nums[:len(nums)-1]))
}

func TestA(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	t.Log(a(nums))
}

func TestB(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	t.Log(b(nums))
}

func deleteAndEarn(nums []int) int {
	var maxVal = 0
	for _, n := range nums {
		maxVal = max(maxVal, n)
	}

	var dp = make([]int, maxVal+1)

	for _, n := range nums {
		dp[n] += n
	}

	var rober = func(nums []int) int {
		var prev, curr int
		for _, num := range nums {
			prev, curr = curr, max(curr, prev+num)
		}
		return curr
	}

	return rober(dp)
}

func TestDeleteAndEarn(t *testing.T) {
	t.Log(deleteAndEarn2([]int{2, 2, 3, 3, 3, 4}))
	t.Log(deleteAndEarn2([]int{2, 4, 3}))
}

func deleteAndEarn2(nums []int) int {
	var set = make(map[int]int)
	var maxVal = 0
	for _, v := range nums {
		set[v] += v
		maxVal = max(maxVal, v)
	}

	var prev, curr = 0, 0
	for i := 0; i <= maxVal; i++ {
		prev, curr = curr, max(curr, prev+set[i])
	}
	return curr
}
