package tree

import "testing"

func TestInsertIntoBST(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var insertIntoBST func(root *TreeNode, val int) *TreeNode
	insertIntoBST = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return &TreeNode{
				Val: val,
			}
		}
		// 先递推 再 回归
		// 4>2 => 2.Right = insertIntoBST(3,4)
		// 4>3 => 3.Right = insertIntoBST(nil,4)
		// 再回归 2.Right = 3带上了4
		if val > root.Val {
			root.Right = insertIntoBST(root.Right, val)
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
		return root
	}

	t.Log(insertIntoBST(root, 4))
}
