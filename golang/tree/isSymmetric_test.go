package tree

import "testing"

/**
*	是否对称，解题思路：
* 判断根节点的左树==右树
* 很容易想到的就是深度搜索dfs
* 和它相关的一道题是：mirrorTree
* 时间复杂度O(n),空间复杂度O(n)
* */
func isSymmetric(root *TreeNode) bool {
	var dfs func(A *TreeNode, B *TreeNode) bool
	dfs = func(A *TreeNode, B *TreeNode) bool {
		if A == nil && B == nil {
			return true
		}
		if A == nil || B == nil || A.Val != B.Val {
			return false
		}

		return dfs(A.Left, B.Right) && dfs(A.Right, B.Left)
	}
	return dfs(root, root)
}

/**
* 另外一个解决方案，采用广度搜索
*	每次取两个节点，判断是否相同
* 加入队列的时候，需要注意一下，A的左节点，B的右节点，A的右节点，B的左节点
* 时间复杂度O(n),空间复杂度O(n)
* */
func isSymmetric1(root *TreeNode) bool {
	var bfs func(A *TreeNode, B *TreeNode) bool
	bfs = func(A *TreeNode, B *TreeNode) bool {
		queue := []*TreeNode{A, B}
		for len(queue) > 0 {
			l, r := queue[0], queue[1]
			queue = queue[2:]
			if l == nil && r == nil {
				continue
			}
			if l == nil || r == nil || l.Val != r.Val {
				return false
			}
			queue = append(queue, l.Left, r.Right, l.Right, r.Left)
		}
		return true
	}
	return bfs(root, root)
}

func TestIsSymmetric(t *testing.T) {
	res := isSymmetric(SortedArrayToBST([]int{1, 2, 1}))
	t.Log(res)

	res = isSymmetric1(SortedArrayToBST([]int{1, 2, 1}))
	t.Log(res)
}
