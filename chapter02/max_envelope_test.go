package chapter02

import (
	"sort"
	"testing"
)

type Envelope [][]int // 排序算法

func (e Envelope) Len() int { return len(e) }
func (e Envelope) Less(i, j int) bool { // 先按宽度正序排序，宽度相等的按高度逆序排序
	if e[i][0] == e[j][0] {
		return e[i][1] > e[j][1]
	}
	return e[i][0] < e[j][0]
}
func (e Envelope) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

func TestMaxEnvelopes(t *testing.T) {
	height := []int{}                  // 套娃的高度
	solution := func(nums []int) int { // 二分搜索寻找最长增长子序列
		var piles int
		length := len(nums)
		top := make([]int, length)
		for i := 0; i < length; i++ {
			top = append(top, 0)
		}
		for i := 0; i < length; i++ {
			poker := nums[i]
			left, right := 0, piles
			for left < right {
				mid := left + (right-left)/2
				if top[mid] > poker {
					right = mid
				} else if top[mid] < poker {
					left = mid + 1
				} else if top[mid] == poker {
					right = mid
				}
			}
			if left == piles {
				piles++
			}
			top[left] = poker
		}

		return piles
	}

	a := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
		{5, 2},
		{1, 8},
	}

	e := Envelope(a) // 转换成可排序的对象

	sort.Sort(e)
	for i := 0; i < len(e); i++ {
		height = append(height, e[i][1])
	}

	t.Log(e)
	t.Log(solution(height)) // 取高度的LIS
}
