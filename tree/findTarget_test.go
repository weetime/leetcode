package tree

import "testing"

func TestFindTarget(t *testing.T) {
	root := &TreeNode{ // 初始化树，注意树中可能存在负数
		Val: 1,
	}
	// 题解:
	// 深度遍历整棵树，用一个map存储 targer-root.Val
	// 在遍历的时候，如果map中存在差值则返回true
	var hashMap = make(map[int]struct{}) // 差值map
	var target = 2                       // 目标和
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil { // 整棵树遍历结束，未找到
			return false
		}
		
		if _, ok := hashMap[root.Val]; ok {   // 判断差值是否存在
			return true
		}
		hashMap[target-root.Val] = struct{}{} // 将差值存入map中
		return dfs(root.Left) || dfs(root.Right)
	}

	t.Log(dfs(root))
}
