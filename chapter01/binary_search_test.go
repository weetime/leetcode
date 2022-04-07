package chapter01

import "testing"

func TestBinarySearch(t *testing.T) {
	var solution func(num []int, target int) int
	solution = func(num []int, target int) int {
		left, right := 0, len(num)-1
		for left <= right {
			mid := left + (right-left)/2
			if num[mid] == target {
				return mid
			} else if num[mid] > target {
				right = mid - 1
			} else if num[mid] < target {
				left = mid + 1
			}
		}

		return -1
	}

	res := solution([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	t.Log(res)
}

func TestLeftBound(t *testing.T) {
	var solution func(num []int, target int) int
	solution = func(num []int, target int) int {
		left, right := 0, len(num)-1
		for left <= right {
			mid := left + (right-left)/2
			if num[mid] < target {
				left = mid + 1
			} else if num[mid] > target {
				right = mid - 1
			} else if num[mid] == target {
				right = mid - 1 // 如果找到了 就缩小右边界，锁定左边界
			}

		}

		if left >= len(num) || num[left] != target { // 如果越界了 就返回-1
			return -1
		}
		return left
	}

	res := solution([]int{1, 2, 2, 3, 4, 5, 6, 7}, 2)
	t.Log(res)
}

func TestRightBound(t *testing.T) {
	var solution func(num []int, target int) int
	solution = func(num []int, target int) int {
		left, right := 0, len(num)-1
		for left <= right {
			mid := left + (right-left)/2
			if num[mid] < target {
				left = mid + 1
			} else if num[mid] > target {
				right = mid - 1
			} else if num[mid] == target { // 如果找到了就缩小左边界，锁定右边界
				left = left + 1
			}

		}

		if right < 0 || num[right] != target { //  如果越界了 或者最后一个也找不到 就返回-1
			return -1
		}
		return right
	}

	res := solution([]int{1, 2, 2, 3, 4, 5, 6, 7}, 2)
	t.Log(res)
}

func TestBinarySearch2(t *testing.T) {
	var solution func(nums []int, target int) int
	solution = func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] > target {
				right = mid - 1
			} else if nums[target] < target {
				left = mid + 1
			} else if nums[mid] == target {
				return mid
			}
		}
		return -1
	}
	t.Log(solution([]int{1, 2, 3, 4, 5, 6, 7}, 2))
}
