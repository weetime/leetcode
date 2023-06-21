package chapter04

import (
	"sort"
	"testing"
)

// 时间复杂度 O(N²) 空间复杂度O(logN)用于排序
func TestThreeSum(t *testing.T) {
	var nums = []int{-1, 0, 1, 2, -1, -4} // 目标数组
	var ans [][]int                       // 结果集
	sort.Ints(nums)                       // 先要排序，防止重复结果 排序用的Golang的快排
	n := len(nums)                        // 数组长度
	for first := 0; first < n; first++ {  // 第一重循环
		if first > 0 && nums[first] == nums[first-1] { // 排序后数据 如果存在重复的 右移指针
			continue
		}
		target := -1 * nums[first]                      // 和为0，所以取相反数
		third := n - 1                                  // 第三个指针
		for second := first + 1; second < n; second++ { // 第二重循环 采用双指针 第二个指针右移，第三个指针左移
			if second > first+1 && nums[second] == nums[second-1] { // 排序后的数据，如果存在重复，右移指针
				continue
			}
			for second < third && nums[second]+nums[third] > target { // 由于是排好序的 所以如果第三个数过大 就左移指针
				third--
			}
			if second == third { // 如果指针重合 说明没有找到跳出当前循环
				break
			}
			if nums[second]+nums[third] == target { // 找到了结果放到结果集
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	t.Log(ans)
}
