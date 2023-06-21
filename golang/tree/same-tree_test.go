package tree

import "testing"

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var isSameTree func(p *TreeNode, q *TreeNode) bool
	isSameTree = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if q.Val != p.Val {
			return false
		}

		// 遍历左侧p.Left = 2,q.Left =2
		// p.Left=nil,q.Left=nil
		// p.Right=nil,q.Right=nil
		// p.Right = 3, q.Right=3 // 开始回溯到上个节点
		// p.Left = nil,q.Left=nil
		// p.Right=nil,q.Right=nil
		return isSameTree(p.Left, q.Left) && isSameTree(q.Right, p.Right)
	}
	t.Log(isSameTree(p, q))
}

func TestIsSameTree2(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var isSameTree func(p *TreeNode, q *TreeNode) bool
	isSameTree = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		queue := []*TreeNode{}
		queue = append(queue, p, q) // 先把p,q root放进去
		for len(queue) > 0 {
			node1, node2 := queue[0], queue[1] // 每次取两个
			queue = queue[2:]                  // 裁剪
			if node1 == nil && node2 == nil {  // 都为空 continue
				continue
			}
			if node1 == nil || node2 == nil || node1.Val != node2.Val { // 不相同
				return false
			}

			queue = append(queue, node1.Left, node2.Left, node1.Right, node2.Right) // 将左右节点放到queue中
		}
		return true
	}
	t.Log(isSameTree(p, q))
}
