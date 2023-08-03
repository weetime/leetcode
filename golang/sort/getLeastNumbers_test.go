package sort

import (
	"sort"
	"testing"
)

/*
* 取第k小的数，解题思路1，快排，取0:k的下标
 */
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func TestSort(t *testing.T) {
	arr := []int{3, 1, 2, 5, 4}
	// sort.Ints(a)
	// t.Log(a)
	res := getLeastNumbers(arr, 2)
	t.Log(res)
}
