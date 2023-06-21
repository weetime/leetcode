package tree

import "testing"

func TestLowestCommonAncestorBST(t *testing.T) {
	var root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}
		if p.Val < root.Val && q.Val < root.Val {
			return dfs(root.Left, p, q)
		} else if p.Val > root.Val && q.Val > root.Val {
			return dfs(root.Right, p, q)
		}
		return root
	}
	t.Log(dfs(root, root.Left, root.Right))
}
