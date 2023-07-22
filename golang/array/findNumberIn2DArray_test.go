package array

import (
	"sort"
	"testing"
)

/**
*	思路很重要，解题思路1，由于每行是递增的，可以采用二分查找
 */
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

/**
* 解决方案2：
* 从右上角开始比较，>target左移，<target 下移动，找到了就返回true 超过边界就返回false
**/
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows, cols := len(matrix), len(matrix[0])
	x, y := 0, cols-1
	for x < rows && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func TestFindNumberIn2DArray(t *testing.T) {
	res := findNumberIn2DArray([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 10}}, 9)
	t.Log(res)

	res = findNumberIn2DArray2([][]int{{1}}, 1)
	t.Log(res)
}
