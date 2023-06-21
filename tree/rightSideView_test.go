package tree

import (
	"testing"
)

func TestRightSideView(t *testing.T) {
	var root = &TreeNode{ // 初始化一棵树
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	var bfs func(root *TreeNode, level int) []int // 广度搜索
	bfs = func(root *TreeNode, level int) []int {
		var values []int // 存放每层最右端节点的值
		if root == nil { // 注意root为空的情况
			return values
		}
		queue := []*TreeNode{root}                // 初始化队列
		for level := 0; len(queue) > 0; level++ { // 注意queue是否为空
			q := queue
			queue = nil
			values = append(values, q[len(q)-1].Val) // 取出队列最右侧的元素即为右节点的值
			for _, node := range q {                 // 先放左节点，再放右节点，注意顺序，方便后续取值
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}
		return values
	}
	t.Log(bfs(root, 0))
}
