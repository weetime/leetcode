package huisu

import "testing"

/*
*	解法1:动态规划，二维矩阵visited，第visited[i][j] 是否可达，取决于i和j的数位和是否大于k
* 同时达到visited[i][j] 必须经过visited[i-i][j] 或者 visited[i][j-1] (向下或者向右移动才有意义)
* 如果前一个节点不可达，那么visited[i][j]也不可达
 */
func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	var get func(x int) int
	get = func(x int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x /= 10
		}
		return res
	}
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	ans := 1
	visited[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || get(i)+get(j) > k {
				continue
			}
			if i-1 >= 0 {
				visited[i][j] |= visited[i-1][j]
			}
			if j-1 >= 0 {
				visited[i][j] |= visited[i][j-1]
			}
			if visited[i][j] > 0 {
				ans++
			}
		}
	}
	return ans
}

/*
*
* 截图思路2：广度搜索
 */
func movingCount1(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	type pair struct {
		x int
		y int
	}

	var get func(x int) int
	get = func(x int) int {
		res := 0
		for x != 0 {
			res += x % 10
			x /= 10
		}
		return res
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	ans := 1
	visited[0][0] = true
	directions := []pair{{x: 1, y: 0}, {x: 0, y: 1}}
	queue := [][]int{{0, 0}}
	for len(queue) > 0 {
		target := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			i, j := target[0]+dir.x, target[1]+dir.y
			if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || get(i)+get(j) > k {
				continue
			}
			queue = append(queue, []int{i, j})
			visited[i][j] = true
			ans++
		}
	}
	return ans
}

func TestMovingCount(t *testing.T) {
	// t.Log(1 | 0) => 1
	// t.Log(123%10, 123/10, 12%10, 12/10, 1%10, 1/10)

	res := movingCount(2, 3, 1)
	t.Log(res)

	res = movingCount1(2, 3, 1)
	t.Log(res)
}
