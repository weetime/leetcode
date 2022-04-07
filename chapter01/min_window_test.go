package chapter01

import (
	"math"
	"testing"
)

func TestMinWindow(t *testing.T) {
	var solution func(s, t string) string // 解决方案
	solution = func(s, t string) string {
		var window = map[string]int{} // 滑动窗口
		var need = map[string]int{}   // 目标态
		var valid, left, right, start int
		var step = math.MaxInt // 初始化最大步长
		for _, v := range t {
			need[string(v)]++ // 初始化目标态
		}
		length := len(s)
		for right < length { // 开始求解
			c := string(s[right]) // 取出要判断的字符
			right++               // 右移窗口

			if _, ok := need[c]; ok { // 命中目标态
				window[c]++ // 放入窗口
				if window[c] == need[c] {
					valid++ // 该字符已达到目标态
				}
			}

			for valid == len(need) { // 全部完成目标态
				if right-left < step { // 缩小窗口
					start = left        // 最终返回字符的左边界
					step = right - left // 窗口大小
				}

				d := string(s[left])
				left++ // 缩小窗口
				if _, ok := need[d]; ok {
					if window[d] == need[d] {
						valid-- // 不满足目标态了
					}
					window[d]-- // 从窗口中减去数据
				}

			}
		}
		if step < math.MaxInt { //找到结果
			return s[start : start+step] // Golang语言特性 [左区间, 左区间+步长]
		}

		return "" // 未找到
	}

	res := solution("ADOBECODEBANC", "ABC")
	t.Log(res)
}

func TestCheckInClustion(t *testing.T) {
	var need = map[string]int{}         // 目标态
	var window = map[string]int{}       // 窗口
	var left, right, valid int          // 左边界、右边界、命中
	var solution func(S, T string) bool // 解决方案
	solution = func(S, T string) bool {
		for _, v := range T {
			need[string(v)]++ // 初始化目标态
		}
		for right < len(S) { // 遍历长串
			c := string(S[right])
			right++ // 右移边界

			if _, ok := need[c]; ok { // 命中目标态
				window[c]++
				if window[c] == need[c] {
					valid++ // 满足条件
				}
			}

			for right-left == len(T) { // 窗口达到子串长度
				if valid == len(need) {
					return true
				}
				d := string(S[left])
				left++ // 收缩边界

				if _, ok := need[d]; ok {
					if window[d] == need[d] {
						valid--
					}
					window[d]--
				}
			}
		}

		return false
	}
	t.Log(solution("ADESFE", "DEF"))
}

func TestFindArgus(t *testing.T) {
	var need = map[string]int{}
	var window = map[string]int{}
	var left, right, valid int
	var res []int
	var solution func(s, t string)
	solution = func(s, t string) {
		for _, v := range t {
			need[string(v)]++
		}
		for right < len(s) {
			c := string(s[right])
			right++
			if _, ok := need[c]; ok {
				window[c]++
				if window[c] == need[c] {
					valid++
				}
			}

			for right-left == len(t) {
				if valid == len(need) {
					res = append(res, left)
				}
				d := string(s[left])
				left++
				if _, ok := need[d]; ok {
					if window[d] == need[d] {
						valid--
					}
					window[d]--
				}
			}
		}
	}

	solution("cbaebabacd", "abc")
	t.Log(res)
}

func TestMaxLenSubstr(t *testing.T) {
	var left, right int // 左右
	var window = map[string]int{}
	var solution func(s string) int
	solution = func(s string) int {
		var res float64
		for right < len(s) {
			c := string(s[right])
			right++
			window[c]++
			for window[c] > 1 {
				d := string(s[left])
				left++
				window[d]--
			}
			res = math.Max(float64(res), float64(right-left))
		}
		return int(res)
	}

	t.Log(solution("pwwkew"))
}
