package tree

import "testing"

func TestPreorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var preorderTraversal func(root *TreeNode) []int
	var res []int
	preorderTraversal = func(root *TreeNode) []int {
		if root == nil {
			return res
		}

		res = append(res, root.Val)
		preorderTraversal(root.Left)
		preorderTraversal(root.Right)
		return res
	}

	t.Log(preorderTraversal(tree))
}
