package tree

import "testing"

/**
* 深度搜索，哪边树高就取哪一边的树
 */
func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode, height int) int
	dfs = func(root *TreeNode, height int) int {
		if root == nil {
			return height
		}
		l := dfs(root.Left, height+1)
		r := dfs(root.Right, height+1)
		if l > r {
			return l
		}
		return r
	}

	return dfs(root, 0)
}

/**
* 精简版深度搜索，比上面一个方法简单一点
 */
func maxDepth1(root *TreeNode) int {
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(dfs(root.Left), dfs(root.Right)) + 1
	}

	return dfs(root)
}

/*
* 广度搜索也可以解决这个问题，深度其实就是层级
 */
func maxDepth2(root *TreeNode) int {
	var bfs func(root *TreeNode) int
	bfs = func(root *TreeNode) int {
		var ans int
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				node := queue[0]
				queue = queue[1:]
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			ans++
		}
		return ans
	}

	return bfs(root)
}
func TestMaxDepth(t *testing.T) {
	res := maxDepth(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(res)

	res = maxDepth1(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(res)

	res = maxDepth2(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(res)
}
