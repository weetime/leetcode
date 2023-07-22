package array

import "testing"

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == mid {
			left = mid + 1
		} else if nums[mid] > mid {
			right = mid - 1
		}
	}

	return left
}

func TestMissingNumber(t *testing.T) {
	res := missingNumber([]int{0, 2, 3, 4, 5, 6})
	t.Log(res)
}
