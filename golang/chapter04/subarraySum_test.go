package chapter04

import "testing"

func TestSubArraySum(t *testing.T) {
	nums := []int{1, 1, 1, 2}
	k := 2
	var ans int
	var preSum = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				ans++
			}
		}
	}
	t.Log(ans)
}

func TestSubArraySum2(t *testing.T) {
	nums := []int{1, 1, 1, 2}
	k := 2
	var ans int
	var pre int
	var preSum = make(map[int]int)
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := preSum[pre-k]; ok {
			ans += preSum[pre-k]
		}
		preSum[pre] += 1
	}
	t.Log(ans)
}

func TestXxx(t *testing.T) {
	subarraySum := func(nums []int, k int) int {
		count, pre := 0, 0
		m := map[int]int{}
		m[0] = 1
		for i := 0; i < len(nums); i++ {
			pre += nums[i]
			if _, ok := m[pre-k]; ok {
				count += m[pre-k]
			}
			m[pre] += 1
		}
		return count
	}
	t.Log(subarraySum([]int{1, 1, 1, 2}, 2))
}
