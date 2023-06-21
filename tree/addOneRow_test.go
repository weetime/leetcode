package tree

import (
	"testing"
)

func TestAddOneRow(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	// 题解：

	var depth = 3                          // 插入的层级
	var val = 5                            // 新节点的val
	var dfs func(root *TreeNode, curr int) // 递归 DLR
	dfs = func(root *TreeNode, curr int) {
		if root == nil { // 递归终止条件
			return
		}
		if curr == depth-1 {
			l := root.Left         // 临时存储左节点
			r := root.Right        // 临时存储右节点
			root.Left = &TreeNode{ // 新节点
				Val:  val,
				Left: l,
			}
			root.Right = &TreeNode{ // 新节点
				Val:   val,
				Right: r,
			}
		}
		dfs(root.Left, curr+1)
		dfs(root.Right, curr+1)
	}

	solution := func(root *TreeNode) *TreeNode {
		if depth == 1 { // 特殊兼容depth=1的情况
			return &TreeNode{
				Val:  val,
				Left: root,
			}
		}
		dfs(root, 1)
		return root
	}

	t.Log(solution(root))
	t.Log("end")
}
