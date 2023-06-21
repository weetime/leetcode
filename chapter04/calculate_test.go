package chapter04

import (
	"testing"
	"unicode"
)

func TestCalculate(t *testing.T) {
	var s = "34-(2*2)+1"
	var helper func() int
	var sum func(stack []int) int
	sum = func(stack []int) int {
		ans := 0
		for len(stack) > 0 {
			ans += stack[0]
			stack = stack[1:]
		}
		return ans
	}
	helper = func() int {
		var stack []int
		var num int
		var sign = '+'
		for len(s) > 0 { // 这里的s是全局的，处理一个字符少一个
			c := s[0]                     // 每次弹出一个字符
			s = s[1:]                     // 弹出字符
			if unicode.IsDigit(rune(c)) { // 如果是数字 继续读取
				num = num*10 + int(c-'0') // 转成数字
			}
			if c == '(' { // 遇到括号继续递归
				num = helper() // 开始递归
			}
			// 遇到符号 或者 循环到结尾 处理上一个符号
			if (!unicode.IsDigit(rune(c)) && c != ' ') || len(s) == 0 {
				switch sign { // sign 为上一个 符号 初始化为+
				case '+':
					stack = append(stack, num) // +num
				case '-':
					stack = append(stack, -num) // -num
				case '*':
					pre := stack[len(stack)-1]     // 栈顶
					stack = stack[:len(stack)-1]   // 弹出栈
					stack = append(stack, pre*num) // 计算'*'
				case '/':
					pre := stack[len(stack)-1]     // 栈顶
					stack = stack[:len(stack)-1]   // 弹出栈
					stack = append(stack, pre/num) // 计算'/'
				}
				sign, num = rune(c), 0 // 归位 符号 和 num
			}
			if c == ')' { // 括号递归终止条件
				break
			}
		}
		return sum(stack) // 返回结果集
	}

	t.Log(helper())
}
