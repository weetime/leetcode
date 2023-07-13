package chapter01

import (
	"leetcode/golang/history/tree"
	"testing"
)

func TestMinTreeDepth(t *testing.T) {
	var queue []*tree.TreeNode
	var minDepth func(*tree.TreeNode) int

	push := func(node *tree.TreeNode) {
		queue = append(queue, node)
	}

	pop := func() *tree.TreeNode {
		tmp := queue[0]
		queue = queue[1:]
		return tmp
	}

	minDepth = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}

		queue = append(queue, root)
		depth := 1
		for len(queue) > 0 {
			sz := len(queue)
			for i := 0; i < sz; i++ {
				node := pop()
				if node.Left == nil && node.Right == nil { // 当左节点和右节点都是null的时候返回
					return depth
				}
				if node.Left != nil {
					push(node.Left)
				}
				if node.Right != nil {
					push(node.Right)
				}
			}
			depth++ // 每层结束会+1
		}

		return depth
	}

	treeNode := tree.SortedArrayToBST([]int{3, 1})

	height := minDepth(treeNode)
	t.Log(height)
}
