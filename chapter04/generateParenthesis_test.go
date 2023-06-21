package chapter04

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	var n = 3                                         // 括号数
	var res []string                                  // 结果集
	var backtrack func(left, right int, track string) // 回溯
	backtrack = func(left, right int, track string) {
		if left < 0 || right < 0 { // 如果越过边界返回
			return
		}
		// ))(( 右括号数 先大于 左括号数 直接返回
		if right < left {
			return
		}
		if left == 0 && right == 0 { // 获取结果
			res = append(res, track)
			return
		}
		track += "("                    // 做选择
		backtrack(left-1, right, track) // 进入下个递归
		track = track[:len(track)-1]    // 撤销选择

		track += ")"                    // 做选择
		backtrack(left, right-1, track) // 进入下个递归
		track = track[:len(track)-1]    // 撤销选择
	}
	backtrack(n, n, "") // 开始穷举
	t.Log(res)
}
