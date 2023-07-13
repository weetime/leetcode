package chapter01

import (
	"testing"
)

func TestOpenLock(t *testing.T) {
	var list []string                                   // 队列
	var pop func() string                               // 出队
	var push func(str string)                           // 入队
	var open func(deadends []string, target string) int // 解锁函数
	var plusOne func(lock string, position int) string  // 往上拨
	var minusOne func(lock string, position int) string // 往下拨

	// 从队列中头部弹出一个元素
	pop = func() string {
		first := list[0]
		list = list[1:]
		return first
	}

	// 添加元素到队列尾部
	push = func(str string) {
		list = append(list, str)
	}

	// 加1
	plusOne = func(lock string, position int) string {
		bytes := []byte(lock)
		if bytes[position] == '9' {
			bytes[position] = '0'
		} else {
			bytes[position]++
		}

		return string(bytes)
	}

	// 减1
	minusOne = func(lock string, position int) string {
		bytes := []byte(lock)
		if bytes[position] == '0' {
			bytes[position] = '9'
		} else {
			bytes[position]--
		}
		return string(bytes)
	}

	open = func(deadends []string, target string) int {
		var visited = map[string]struct{}{}     // 标记已访问过的
		var deadendsMap = map[string]struct{}{} // 死亡密码本
		var step int

		// 初始化死亡密码本
		for _, item := range deadends {
			deadendsMap[item] = struct{}{}
		}

		// 开始检索
		push("0000")
		for len(list) > 0 {
			sz := len(list)

			for i := 0; i < sz; i++ {
				current := pop()
				if _, ok := deadendsMap[current]; ok {
					continue // 跳过死亡密码
				}

				if current == target {
					return step // 终止条件
				}

				// 拨动密码锁
				for j := 0; j < 4; j++ {
					p := plusOne(current, j) // 加1
					if _, ok := visited[p]; !ok {
						visited[p] = struct{}{}
						push(p)
					}
					q := minusOne(current, j) // 减1
					if _, ok := visited[q]; !ok {
						visited[q] = struct{}{}
						push(q)
					}
				}
			}
			step++
		}
		return -1
	}
	res := open([]string{"1111", "2222"}, "9999")
	t.Log(res)
}

