package chapter01

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueen(t *testing.T) {
	n := 8
	var res = []string{}
	var fn func([][]int, int)
	var isValid func(board [][]int, row int, col int) bool

	board := initBoard(n)
	fn = func(board [][]int, row int) {
		if row == n {
			str := ""
			for i := 0; i < n; i++ {
				str += "|"
				for j := 0; j < n; j++ {
					str = str + strconv.Itoa(board[i][j])
				}
			}
			res = append(res, str)
			// fmt.Println(res)
			return
		}
		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = 1
			fn(board, row+1)
			board[row][col] = 0
		}
	}

	isValid = func(board [][]int, row, col int) bool {
		for i := row - 1; i >= 0; i-- {
			if board[i][col] == 1 {
				return false
			}
		}

		j := col + 1
		for i := row - 1; i >= 0; i-- {
			if j >= n {
				break
			}

			if board[i][j] == 1 {
				return false
			}
			j++
		}

		j = col - 1
		for i := row - 1; i >= 0; i-- {
			if j < 0 {
				break
			}

			if board[i][j] == 1 {
				return false
			}
			j--
		}
		return true
	}

	fn(board, 0)

	for _, data := range res {
		fmt.Println(data)
	}
}

func initBoard(n int) [][]int {
	var board = [][]int{}
	for i := 0; i < n; i++ {
		var tmp []int
		for j := 0; j < n; j++ {
			tmp = append(tmp, 0)
		}
		board = append(board, tmp)
	}

	return board
}

func TestQueue2(t *testing.T) {
	var number int                                     // 棋盘大小
	var board [][]int                                  // 棋盘
	var res [][]string                                 // 返回结果
	var initBoard func(int)                            // 初始化棋盘
	var isValid func(board [][]int, row, col int) bool // 判断是否可放置皇后
	var solution func(board [][]int, row int)          // 解决方案

	// 初始化棋盘
	initBoard = func(i int) {
		for i := 0; i < number; i++ {
			tmp := []int{}
			for j := 0; j < number; j++ {
				tmp = append(tmp, 0)
			}
			board = append(board, tmp)
		}
	}

	// 判断是否能放置皇后
	isValid = func(board [][]int, row, col int) bool {
		for i := row - 1; i >= 0; i-- {
			if board[i][col] == 1 {
				return false
			}
		}

		j := col + 1
		for i := row - 1; i >= 0; i-- {
			if j >= number {
				break
			}
			if board[i][j] == 1 {
				return false
			}
			j++
		}

		j = col - 1
		for i := row - 1; i >= 0; i-- {
			if j < 0 {
				break
			}
			if board[i][j] == 1 {
				return false
			}
			j--
		}

		return true
	}

	// 寻找解决方案
	solution = func(board [][]int, row int) {
		if row == number { // 触发结束条件
			tmp := []string{}
			for i := 0; i < number; i++ {
				str := ""
				for j := 0; j < number; j++ {
					if board[i][j] == 0 {
						str += "."
					} else {
						str += "Q"
					}
				}
				tmp = append(tmp, str)
			}
			res = append(res, tmp)
			return
		}

		// 寻找解决方案
		for col := 0; col < number; col++ {
			// 排除不合法选项
			if !isValid(board, row, col) {
				continue
			}

			board[row][col] = 1    // 做选择
			solution(board, row+1) // 执行下个决策树
			board[row][col] = 0    // 撤销选择
		}
	}

	number = 4
	initBoard(number)
	solution(board, 0)
	t.Log(res)
}
