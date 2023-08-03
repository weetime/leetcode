package tree

import "testing"

/**
* 解题思考：第K大的元素，典型的中序遍历，RDL，这样就能得到一个逆序的数组，取第k个即可
 */
func kthLargest(root *TreeNode, k int) int {
	var res int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			res = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

func TestKthLargest(t *testing.T) {
	res := kthLargest(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), 3)
	t.Log(res)
}
