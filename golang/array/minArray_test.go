package array

import (
	"testing"
)

func minArray(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	maxLen := n
	for numbers[0] >= numbers[n-1] && maxLen > 0 {
		temp := numbers[n-1]
		numbers = append([]int{temp}, numbers[:n-1]...)
		maxLen--
	}

	return numbers[0]
}

func minArray1(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := left + (right-left)/2
		if numbers[mid] == numbers[right] {
			right--
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return numbers[left]
}

func TestMinArray(t *testing.T) {
	res := minArray([]int{3, 4, 5, 2, 1})
	t.Log(res)

	res = minArray1([]int{3, 1, 3})
	t.Log(res)
}
