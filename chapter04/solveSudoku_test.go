package chapter04

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	var m = 9                                        // 9x9 的格子
	var board [][]byte                               // 数独
	var isValid func(row, col int, target byte) bool // 是否合法
	var traceback func(row, col int) bool            // 回溯

	isValid = func(row, col int, target byte) bool {
		for i := 0; i < m; i++ {
			if board[row][i] == target { // 同一行不能重复
				return false
			}
			if board[i][col] == target { // 同一列不能重复
				return false
			}
			if board[row/3*3+i/3][col/3*3+i%3] == target { //3x3区域不能重复
				return false
			}
		}
		return true
	}

	// 初始化棋盘
	func() {
		board = make([][]byte, m)
		for i := 0; i < m; i++ {
			board[i] = make([]byte, m)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				board[i][j] = '.'
			}
		}
	}()

	traceback = func(row, col int) bool {
		if col == m {
			return traceback(row+1, 0)

		}
		if row == m {
			return true
		}
		if board[row][col] != '.' {
			return traceback(row, col+1)
		}

		for i := '1'; i <= '9'; i++ {
			if !isValid(row, col, byte(i)) {
				continue
			}
			board[row][col] = byte(i)
			if traceback(row, col+1) {
				return true
			}
			board[row][col] = '.'
		}
		return false
	}
	traceback(0, 0)
	for i := 0; i < m; i++ {
		fmt.Println()
		for j := 0; j < m; j++ {
			fmt.Print(string(board[i][j]))
		}
	}
}
