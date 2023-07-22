package array

import "testing"

// 如果是无序的，需要对比每一个
func twoSum(nums []int, target int) []int {
	res := []int{}
	n := len(nums)
	if n == 0 {
		return res
	}
	i, j := 0, n-1
	for i <= j {
		if j == 0 {
			return res
		}
		if i == j {
			i = 0
			j = j - 1
		}
		if nums[i]+nums[j] == target {
			res = append(res, nums[i], nums[j])
			return res
		}
		if nums[j] >= target {
			j--
			continue
		}
		if nums[j] < target {
			i++
		}
	}

	return res
}

// 如是是有序的
func twoSum2(nums []int, target int) []int {
	res := []int{}
	n := len(nums)
	if n == 0 {
		return res
	}
	i, j := 0, n-1
	for i <= j {
		temp := nums[i] + nums[j]
		if temp == target {
			res = append(res, nums[i], nums[j])
			return res
		}
		if temp >= target {
			j--
		} else {
			i++
		}
	}

	return res
}

func TestTwoSum(t *testing.T) {
	res := twoSum([]int{2, 3, 4, 5, 6}, 10)
	t.Log(res)
	res = twoSum2([]int{2, 3, 4, 5, 6}, 10)
	t.Log(res)
}
