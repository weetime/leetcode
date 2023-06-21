package tree

import "testing"

func TestConstruct2DArray(t *testing.T) {
	var construct2DArray func(original []int, m int, n int) [][]int
	construct2DArray = func(original []int, m, n int) [][]int {
		length := len(original)
		if length != m*n {
			return nil
		}
		var res = make([][]int, m)
		for i := 0; i < m; i++ {
			res[i] = make([]int, n)
		}
		for i := 1; i <= m; i++ {
			index := 0
			for j := (i - 1) * n; j <= i*n-1; j++ {
				res[i-1][index] = original[j]
				index++
			}
		}
		return res
	}

	t.Log(construct2DArray([]int{1, 2, 3, 4}, 2, 2))
}

func TestConstruct2DArray2(t *testing.T) {
	var construct2DArray func(original []int, m int, n int) [][]int
	construct2DArray = func(original []int, m, n int) [][]int {
		length := len(original)
		if length != m*n {
			return nil
		}
		// var res = make([][]int, 0, m)
		// for i := 0; i < length; i += n {
		// 	res = append(res, original[i:i+n])
		// }
		var res = make([][]int, m)
		for i := 0; i < m; i++ {
			res[i] = original[i*n : (i+1)*n]
		}
		return res
	}

	t.Log(construct2DArray([]int{1, 2, 3, 4}, 2, 2))
}
