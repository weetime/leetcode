package moni

import "testing"

/**
*	模拟其实了解了原理，其实很简单
* 先定义四个方向，按方向遍历，生成访问路径
* 定义一个矩阵，存储坐标是否被访问过了
* 达到边界或者被访问过了，就移动方位，直到全部遍历完为止
* 最后输出这个访问路径就可以了
* 时间复杂度O(mn)，空间复杂度O(mn)
 */
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	row, col := len(matrix), len(matrix[0])
	visitor := make([][]bool, row)
	for i := 0; i < row; i++ {
		visitor[i] = make([]bool, col)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directionIndex := 0
	total := row * col
	newRow, newCol := 0, 0
	res = make([]int, total)
	for i := 0; i < total; i++ {
		res[i] = matrix[newRow][newCol]
		visitor[newRow][newCol] = true
		nextRow, nextCol := newRow+directions[directionIndex][0], newCol+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= row || nextCol < 0 || nextCol >= col || visitor[nextRow][nextCol] {
			directionIndex = (1 + directionIndex) % 4
		}
		newRow = newRow + directions[directionIndex][0]
		newCol = newCol + directions[directionIndex][1]
	}

	return res
}

func TestSpiralOrder(t *testing.T) {
	res := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	t.Log(res)
}
