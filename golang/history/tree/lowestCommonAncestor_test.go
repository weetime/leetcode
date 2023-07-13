package tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	var root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	// 题解：
	// 这里需要采用后序遍历 LRD, 这样第一次出现公共节点一定是最近的公共节点
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil { // 递归终止条件
			return nil
		}
		if root == p || root == q { // 只要找到了一个就得返回
			return root
		}
		l := dfs(root.Left, p, q)
		r := dfs(root.Right, p, q)
		if l != nil && r != nil { // p,q一定是分散在root的左右的
			return root
		}
		if l == nil { // 右侧找到了
			return r
		}
		return l // 左侧找到了
	}
	res := dfs(root, root.Left.Left, root.Right)
	t.Log(res)
}
