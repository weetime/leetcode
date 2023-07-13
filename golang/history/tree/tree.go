package tree

import (
	"container/list"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

/**
      0
     / \
   -3   9
   /   /
 -10  5


LDR: -10 -3 0 5 9

DLR: 0 -3 -10 9 5

LRD: -10 -3 5 9 0
**/

// 中序遍历LDR
func ldrBST(t *TreeNode) {
	if t.Left != nil {
		ldrBST(t.Left)
	}
	log.Println(t.Val)
	if t.Right != nil {
		ldrBST(t.Right)
	}
}

// 前序遍历 DLR
func dlrBST(t *TreeNode) {
	log.Println(t.Val)
	if t.Left != nil {
		dlrBST(t.Left)
	}
	if t.Right != nil {
		dlrBST(t.Right)
	}
}

// 后序遍历 LRD
func lrdBST(t *TreeNode) {
	if t.Left != nil {
		lrdBST(t.Left)
	}
	if t.Right != nil {
		lrdBST(t.Right)
	}
	log.Println(t.Val)
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	return dfs_stack(root, L, R)
}

func dfs_stack(root *TreeNode, L int, R int) int {
	res := 0
	levelQueue := list.New()
	levelQueue.PushBack(root)

	for e := levelQueue.Back(); e != nil; e = e.Next() {
		node := e.Value.(*TreeNode)
		if node != nil {
			if node.Val >= L && node.Val <= R {
				res += node.Val
			}
			if L > node.Val {
				levelQueue.PushBack(node.Left)
			}
			if node.Val < R {
				levelQueue.PushBack(node.Right)
			}
		}
	}

	return res
}
