package tree

import (
	"testing"
)

func TestIsCousins(t *testing.T) {
	root := &TreeNode{ // 初始化树
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	// 题解:x,y如果要是堂兄弟节点，则必须深度相等且父节点不相等，在遍历整棵树的时候
	// 检测到val=x||val=y的时候，记录当前的深度 和 父节点信息
	// 最后判断 是否找到了x或者y,并检查x,y的深度及各自的父节点
	var x, y = 3, 5                // 目标值
	var xDepth, yDepth int         // x,y 深度
	var xFount, yFound bool        // x,y深度
	var xParent, yParent *TreeNode // x,y 父节点
	var dfs func(root *TreeNode, parent *TreeNode, level int)
	dfs = func(root *TreeNode, parent *TreeNode, level int) {
		if root == nil { // 遍历完成
			return
		}
		if root.Val == x { // 记录x的相关信息
			xFount, xParent, xDepth = true, parent, level
		} else if root.Val == y { // 记录y的相关信息
			yFound, yParent, yDepth = true, parent, level
		}

		if xFount && yFound { // 如果找了提前退出
			return
		}

		dfs(root.Left, root, level+1)
		if xFount && yFound { // 如果找到了提前结束
			return
		}
		dfs(root.Right, root, level+1)
	}

	dfs(root, nil, 0)
	res := xFount && yFound && (xDepth == yDepth) && (xParent != yParent)
	t.Log(res)
}
