package math

import "testing"

/**
*	随机抽5张牌判断是否是顺子
* 解题思路，大小王为0可以代表任意值
* 除大小王以外不能有重复的，所以弄一个map判断
* 最大值和最小值之差小于5(因为可能存在0代表任意数，所以差值小于5即可)
* */
func isStraight(nums []int) bool {
	exist := make(map[int]struct{})
	min, max := 14, 0
	for _, n := range nums {
		if n == 0 {
			continue
		}
		if _, ok := exist[n]; ok {
			return false
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		exist[n] = struct{}{}
	}

	return max-min < 5
}

func TestIsStraight(t *testing.T) {
	res := isStraight([]int{0, 0, 2, 3, 5})
	t.Log(res)
}
