package array

import "testing"

func findRepeatNumber(nums []int) int {
	tmp := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := tmp[v]; ok {
			return v
		}
		tmp[v] = struct{}{}
	}
	return -1
}

func TestFindRepeatNumber(t *testing.T) {
	res := findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
	t.Log(res)
}
