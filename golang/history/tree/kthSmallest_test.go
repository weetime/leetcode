package tree

import "testing"

func TestKthSmallest(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode { // 构建一颗二叉树
		if left > right {
			return nil
		}

		mid := (left + right) >> 1
		return &TreeNode{
			array[mid],
			buildTree(left, mid-1),
			buildTree(mid+1, right),
		}
	}
	root := buildTree(0, len(array)-1)
	var list []int
	var k = 3
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root != nil {
			dfs(root.Left)
			list = append(list, root.Val)
			if len(list) == k {
				return root.Val
			}
			dfs(root.Right)
		}
		return -1
	}

	t.Log(dfs(root))
}

func TestKthSmallest2(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode { // 构建一颗二叉树
		if left > right {
			return nil
		}

		mid := (left + right) >> 1
		return &TreeNode{
			array[mid],
			buildTree(left, mid-1),
			buildTree(mid+1, right),
		}
	}
	root := buildTree(0, len(array)-1)
	var kthSmallest func(root *TreeNode, k int) int
	kthSmallest = func(root *TreeNode, k int) int {
		stack := []*TreeNode{}
		for {
			for root != nil {
				stack = append(stack, root)
				root = root.Left
			}
			stack, root = stack[:len(stack)-1], stack[len(stack)-1]
			k--
			if k == 0 {
				return root.Val
			}
			root = root.Right
		}
	}
	res := kthSmallest(root, 3)
	t.Log(res)
}
