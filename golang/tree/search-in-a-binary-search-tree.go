package tree

import (
	"container/list"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val || root == nil {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

func searchBST2(root *TreeNode, val int) *TreeNode {
	queue := list.New()
	queue.PushBack(root)

	for e := queue.Back(); e != nil; e = e.Next() {
		n := e.Value.(*TreeNode)
		if n == nil || n.Val == val {
			return n
		}
		if n.Val > val {
			queue.PushBack(n.Left)
		} else if n.Val < val {
			queue.PushBack(n.Right)
		}
	}

	return nil
}
