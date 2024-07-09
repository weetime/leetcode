package dynamicprogramming

import (
	"testing"
)

func UniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	return UniquePaths(m-1, n) + UniquePaths(m, n-1)
}

func UniquePaths2(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func UniquePaths3(m int, n int) int {
	dp := make([]int, m)
	for i := range m {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[m-1]
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name   string
		m      int
		n      int
		result int
	}{
		{"test1", 1, 1, 1},
		{"test2", 1, 2, 1},
		{"test3", 2, 1, 1},
		{"test4", 2, 2, 2},
		{"test5", 3, 3, 6},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if UniquePaths(test.m, test.n) != test.result {
				t.Errorf("UniquePaths(%d, %d) = %d, want %d", test.m, test.n, UniquePaths(test.m, test.n), test.result)
			}
			if UniquePaths2(test.m, test.n) != test.result {
				t.Errorf("UniquePaths2(%d, %d) = %d, want %d", test.m, test.n, UniquePaths2(test.m, test.n), test.result)
			}
			if UniquePaths3(test.m, test.n) != test.result {
				t.Errorf("UniquePaths3(%d, %d) = %d, want %d", test.m, test.n, UniquePaths3(test.m, test.n), test.result)
			}
		})
	}
}
