package chapter01

import (
	"testing"
)

// 全排列
func TestCombo(t *testing.T) {
	var fn func([]int)
	array := []int{1, 2, 3}
	res := [][]int{}
	track := []int{}

	fn = func(track []int) {
		if len(track) == len(array) {
			res = append(res, track)
			return
		}

		for _, v := range array {
			if isSet(track, v) {
				continue
			}
			track = append(track, v)
			fn(track)
			track = pushBack(track)
		}
	}
	fn(track)
	t.Log(res)
}

func isSet(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func pushBack(arr []int) []int {
	return arr[0 : len(arr)-1]
}
