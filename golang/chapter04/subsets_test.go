package chapter04

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	var res [][]int                                        // 接收的结果集
	var subsets func(nums []int) [][]int                   // 子集
	var backTrack func(nums []int, start int, track []int) // 回溯处理
	backTrack = func(nums []int, start int, track []int) {
		// var tmp = make([]int, len(track))
		// copy(tmp, track)
		// res = append(res, tmp)
		// 将路径写入到结果集中，注意数组的深度复制 copy
		res = append(res, append([]int{}, track...)) // 添加路径
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])  // 做选择
			backTrack(nums, i+1, track)     // 递归回溯
			track = track[0 : len(track)-1] // 撤销选择
		}
	}
	subsets = func(nums []int) [][]int {
		backTrack(nums, 0, []int{})
		return res
	}
	t.Log(subsets([]int{1, 2, 3}))
}
