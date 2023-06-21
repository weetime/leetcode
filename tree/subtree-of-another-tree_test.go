package tree

import (
	"testing"
)

func TestIsSubtree(t *testing.T) {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	var subRoot = &TreeNode{
		Val: 1,
	}
	var isSubtree func(root *TreeNode, subRoot *TreeNode) bool
	var dfs func(root *TreeNode, subRoot *TreeNode, needCheck bool) bool
	isSubtree = func(root, subRoot *TreeNode) bool {
		if root == nil || subRoot == nil {
			return false
		}

		return dfs(root, subRoot, false)
	}

	dfs = func(root, subRoot *TreeNode, needCheck bool) bool {
		if needCheck {
			if root == nil && subRoot == nil {
				return true
			}
			if root == nil || subRoot == nil {
				return false
			}
			if root.Val != subRoot.Val {
				return false
			}
			return dfs(root.Left, subRoot.Left, true) && dfs(root.Right, subRoot.Right, true)
		} else {
			if root == nil {
				return false
			}
			if root.Val == subRoot.Val {
				return dfs(root.Left, subRoot.Left, true) && dfs(root.Right, subRoot.Right, true)
			}
			return dfs(root.Left, subRoot, false) || dfs(root.Right, subRoot, false)
		}
	}
	t.Log(isSubtree(root, subRoot))
}

func TestIsSubtree2(t *testing.T) {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	var subRoot = &TreeNode{
		Val: 1,
	}
	var dfs func(root *TreeNode, subRoot *TreeNode) bool
	var check func(root, subRoot *TreeNode) bool
	check = func(root, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		}
		if root == nil || subRoot == nil {
			return false
		}
		if root.Val == subRoot.Val {
			return check(root.Left, subRoot.Left) && check(root.Right, subRoot.Right)
		}
		return false
	}
	dfs = func(root, subRoot *TreeNode) bool {
		if root == nil {
			return false
		}
		return check(root, subRoot) || dfs(root.Left, subRoot) || dfs(root.Right, subRoot)
	}

	t.Log(dfs(root, subRoot))
}
