package tree

import "testing"

/*
*	平衡二叉树，左节点和右节点相差小于1，深度递归
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(dfs(root.Left)-dfs(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(dfs(root.Left), dfs(root.Right)) + 1
}

func abs(x int) int {
	if x > 0 {
		return -1 * x
	}
	return x
}

func TestIsBalanced(t *testing.T) {
	res := isBalanced(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(res)
}
