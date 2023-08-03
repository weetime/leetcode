package huisu

import "testing"

/**
* 查看字符串是否存在矩阵中
* 解题思路，采用回溯算法
* 找开头，矩阵中，任意一个字符都可以是启动，不能漏
* 检测合法性，每个字符需要和word依次对比，在检测的时候，需要往四个不同的方向扩散
* 回溯最重要的，就是什么时机把做出的选择回溯到上个状态，这里是每次check结束的时候，将当前的访问状态置成false
* 回溯需要注意边界
* 需要不合适的，直接返回false
* 当匹配到最后一个字符的时候，返回true
* 只要任意有一条路径就返回true
**/
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0]) // 需要一个数组标记是否已经被访问过了
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 设计四个方向
	var check func(row, col, index int) bool                // check 函数，如果有true 就直接返回
	check = func(row, col, index int) bool {
		if board[row][col] != word[index] { // 不匹配直接剪枝
			return false
		}
		if index == len(word)-1 { // 匹配到结束位置，返回
			return true
		}
		visited[row][col] = true // 标记被访问过了

		defer func() {
			visited[row][col] = false // 回溯的时候，还原到上一个节点
		}()

		for _, direction := range directions { // 按四个方向往下查找
			newRow, newCol := row+direction[0], col+direction[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow][newCol] { // 注意边界
				if check(newRow, newCol, index+1) {
					return true
				}
			}
		}

		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if check(i, j, 0) { // 每个位置都要检测
				return true
			}
		}
	}

	return false
}
func TestExist(t *testing.T) {
	res := exist([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}, {'B', 'C', 'D'}}, "BEH")
	t.Log(res)
}

/*
复杂度分析

时间复杂度：一个非常宽松的上界为 O(MN⋅3L)O(MN \cdot 3^L)O(MN⋅3
L
 )，其中 M,NM, NM,N 为网格的长度与宽度，LLL 为字符串 word\textit{word}word 的长度。在每次调用函数 check\text{check}check 时，除了第一次可以进入 444 个分支以外，其余时间我们最多会进入 333 个分支（因为每个位置只能使用一次，所以走过来的分支没法走回去）。由于单词长为 LLL，故 check(i,j,0)\text{check}(i, j, 0)check(i,j,0) 的时间复杂度为 O(3L)O(3^L)O(3
L
 )，而我们要执行 O(MN)O(MN)O(MN) 次检查。然而，由于剪枝的存在，我们在遇到不匹配或已访问的字符时会提前退出，终止递归流程。因此，实际的时间复杂度会远远小于 Θ(MN⋅3L)\Theta(MN \cdot 3^L)Θ(MN⋅3
L
 )。

空间复杂度：O(MN)O(MN)O(MN)。我们额外开辟了 O(MN)O(MN)O(MN) 的 visited\textit{visited}visited 数组，同时栈的深度最大为 O(min⁡(L,MN))O(\min(L, MN))O(min(L,MN))。
*/
