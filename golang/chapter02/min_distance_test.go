package chapter02

import (
	"strconv"
	"testing"
)

// 递归方式计算最小编辑距离
func TestMinDistance(t *testing.T) {
	var s1, s2 string          // 两个字符串
	var dp func(i, j int) int  // 递归方法
	var min func(x, y int) int // 取最小距离
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp = func(i, j int) int {
		if i == -1 { // 如果s1先遍历完 返回s2剩余的字符串长度
			return j + 1
		}
		if j == -1 { // 如果s2先遍历完，返回s1剩余的字符串长度
			return i + 1
		}

		if s1[i] == s2[j] { // 如果当前位置相同 则编辑距离不变呢
			return dp(i-1, j-1)
		} else { // 取插入、删除、替换 中最小的距离
			return min(dp(i-1, j)+1, min(dp(i, j-1)+1, dp(i-1, j-1)+1))
		}
	}

	s1 = "rad"
	s2 = "apple"
	t.Log(dp(len(s1)-1, len(s2)-1))
}

// 备忘录模式计算最小编辑距离
func TestMinDistanceMemo(t *testing.T) {
	var s1, s2 string             // 两个字符串
	var memo = map[string]int{}   // 备忘录
	var dp func(i, j int) int     // 递归函数
	var key func(i, j int) string // 备忘录key生成规则
	var min func(x, y int) int    // 比较编辑距离
	key = func(i, j int) string {
		return strconv.Itoa(i) + "." + strconv.Itoa(j)
	}
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp = func(i, j int) int {
		if res, ok := memo[key(i, j)]; ok { // 如果备忘录存在则直接返回
			return res
		}
		if i == -1 { // s1先遍历完，返回s2剩余字符长度
			return j + 1
		}
		if j == -1 { // s2先遍历完，返回s1剩余字符串长度
			return i + 1
		}

		if s1[i] == s2[j] { // 如果对应位置字符串相等，则编辑距离不变
			memo[key(i, j)] = dp(i-1, j-1)
		} else { // 取插入、删除、替换最小编辑距离
			memo[key(i, j)] = min(dp(i, j-1)+1, min(dp(i-1, j)+1, dp(i-1, j-1)+1))
		}

		return memo[key(i, j)] // 最后回归 返回 s1,s2 长度下的最短编辑距离
	}

	s1 = "rad"
	s2 = "apple"
	t.Log(dp(len(s1)-1, len(s2)-1))
}

// 动态规划
func TestMinDistanceDp(t *testing.T) {
	var min func(x, y int) int           // 最小值
	var solution func(s1, s2 string) int // 解决方案
	min = func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	solution = func(s1, s2 string) int {
		m, n := len(s1), len(s2)    // 两个字符串的长度
		var dp = make([][]int, m+1) // dpTable dp[0][i] 和 dp[j][0] 为基态数据
		for i := 0; i <= m; i++ {
			dp[i] = make([]int, n+1) // 初始化二维数据带零值
			dp[i][0] = i             // 初始化基态数据
		}
		for j := 0; j <= n; j++ {
			dp[0][j] = j
		}

		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if s1[i-1] == s2[j-1] { // 从字符串的第1个字符开始比较
					dp[i][j] = dp[i-1][j-1] // 如果相等 编辑距离不变
				} else {
					// 取插入，删除，替换 的最小距离+1
					dp[i][j] = min(dp[i][j-1]+1, min(dp[i-1][j]+1, dp[i-1][j-1]+1))
				}
			}
		}

		// 终态就是 两个字符串都分析完了
		return dp[m][n]
	}

	t.Log(solution("rad", "apple"))
}
