package tree

import (
	"testing"
)

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (checkSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func checkSame(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}

	return checkSame(A.Left, B.Left) && checkSame(A.Right, B.Right)
}

func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	queue := []*TreeNode{A}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == B.Val && checkSame(node, B) {
			return true
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}

func TestIsSubStructure(t *testing.T) {
	res := isSubStructure(SortedArrayToBST([]int{1, 2, 3, 4}), sortedArrayToBST([]int{4}))
	t.Log(res)

	res = isSubStructure1(SortedArrayToBST([]int{1, 2, 3, 4}), sortedArrayToBST([]int{4}))
	t.Log(res)
}
