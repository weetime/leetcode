package tree

import (
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) >> 1
		root := &TreeNode{array[mid], nil, nil}
		root.Left = buildTree(left, mid-1)
		root.Right = buildTree(mid+1, right)

		return root
	}
	root := buildTree(0, len(array)-1) // 初始化一颗二叉树

	var bfs func(root *TreeNode) [][]int // 广度搜索
	bfs = func(root *TreeNode) [][]int {
		var res = [][]int{}
		if root == nil {
			return res
		}
		queue := []*TreeNode{root}                // 存放root节点的 数组
		for level := 0; len(queue) > 0; level++ { // 层级遍历
			q := queue
			queue = nil // 释放queue 或者用数组的裁剪弹出元素
			vals := []int{}
			for _, r := range q { // 遍历
				vals = append(vals, r.Val)
				if r.Left != nil {
					queue = append(queue, r.Left)
				}
				if r.Right != nil {
					queue = append(queue, r.Right)
				}
			}
			if level%2 == 1 { // 奇数的时候，反转数组 双指针
				for i, n := 0, len(vals); i < n/2; i++ {
					vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
				}
			}
			res = append(res, vals)
		}

		return res
	}

	t.Log(bfs(root))
}
