package array

import (
	"sort"
	"testing"
)

func search(nums []int, target int) int {
	var searchIndex func(nums []int, target int) int
	searchIndex = func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (left + right) >> 1
			if nums[mid] == target {
				right = mid - 1
			} else if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			}
		}
		return left
	}

	leftmost := searchIndex(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}

	rightmost := searchIndex(nums, target+1) - 1
	return rightmost - leftmost + 1
}

func search2(nums []int, target int) int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return 0
	}

	rightmost := sort.SearchInts(nums, target+1) - 1
	return rightmost - leftmost + 1
}

func TestSearch(t *testing.T) {
	// t.Log(sort.SearchInts([]int{1, 2, 4, 5}, 6))

	res := search([]int{1, 2, 2, 5, 6, 7, 8, 9}, 2)
	t.Log(res)
}
