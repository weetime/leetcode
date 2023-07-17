package list

import "testing"

func exchange(nums []int) []int {
	odd := []int{}
	even := []int{}
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i]%2 != 0 {
			odd = append(odd, nums[i])
		} else {
			even = append(even, nums[i])
		}
	}
	return append(odd, even...)
}

func exchange2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	for _, val := range nums {
		if val%2 != 0 {
			res[left] = val
			left++
		} else {
			res[right] = val
			right--
		}
	}
	return res
}

func TestExchange(t *testing.T) {
	res := exchange([]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1})
	t.Log(res)

	res = exchange2([]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1})
	t.Log(res)
}
