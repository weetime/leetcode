package chapter04

import (
	"strconv"
	"strings"
	"testing"
)

func TestSlidingPuzzle(t *testing.T) {
	var board = [][]int{ // 初始化board
		{4, 1, 2},
		{5, 0, 3},
	}

	start := func() string { // 将初始化board转换成一维字符串
		start := ""
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				start += strconv.Itoa(board[i][j])
			}
		}
		return start
	}()

	var bfs func() int
	bfs = func() int { // 广度搜索
		var step int                            // 返回需要的步骤
		var target = "123450"                   // 目标状态
		var queue []string                      // 搜索的字符串
		var visited = make(map[string]struct{}) // 剪枝标记已经访问过的数据
		var neighbor = [][]int{                 // 相邻的索引
			{1, 3}, // 0 相邻的索引 1,3
			{0, 2, 4},
			{1, 5},
			{0, 4},
			{1, 3, 5},
			{2, 4},
		}

		queue = append(queue, start) // 加入到queue中
		for len(queue) > 0 {
			q := queue // 用一个新队列接收，方便重新计数
			queue = nil
			for _, current := range q { // 遍历队列
				if current == target { // 找到目标状态返回
					return step
				}
				idx := 0
				for current[idx] != '0' { // 找到0的索引
					idx++
				}
				for _, index := range neighbor[idx] {
					tmp := strings.Split(current, "")
					tmp[idx], tmp[index] = tmp[index], tmp[idx] // 交互数据索引
					newStr := strings.Join(tmp, "")             // 转换成新的字符
					if _, ok := visited[newStr]; !ok {          // 未访问过的放入到queue中
						queue = append(queue, newStr)
						visited[newStr] = struct{}{}
					}
				}
			}

			step++ // 增加步骤
		}

		return -1
	}

	t.Log(bfs())
}
