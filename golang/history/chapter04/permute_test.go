package chapter04

import "testing"

// 46. 全排列
func TestPermute(t *testing.T) {
	var res [][]int
	var permute func(nums []int) [][]int
	var backTrack func(nums []int, track []int)
	var isSet func(trace []int, target int) bool // 判断是否存在选择的结果集中
	isSet = func(trace []int, target int) bool {
		for _, val := range trace {
			if target == val {
				return true
			}
		}
		return false
	}
	backTrack = func(nums []int, track []int) { // 回溯
		if len(track) == len(nums) { // 添加路径
			res = append(res, append([]int{}, track...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if false == isSet(track, nums[i]) {
				track = append(track, nums[i]) // 做出选择
				backTrack(nums, track)         // 回溯
				track = track[:len(track)-1]   // 撤销选择
			}
		}
	}
	permute = func(nums []int) [][]int {
		backTrack(nums, []int{})
		return res
	}
	permute([]int{1, 2, 3})
	t.Log(res)
}
