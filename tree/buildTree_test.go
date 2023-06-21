package tree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	// 前序 + 中序 构建二叉树
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	var buildTree func(preorder, inorder []int) *TreeNode
	buildTree = func(preorder, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		tree := &TreeNode{preorder[0], nil, nil}
		var i = 0
		for ; i < len(inorder); i++ {
			if preorder[0] == inorder[i] {
				break
			}
		}
		tree.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
		tree.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
		return tree
	}
	root := buildTree(preorder, inorder)
	t.Log(root)
}

func TestBuildTree2(t *testing.T) {
	// 后序 + 中序 还原二叉树
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	var buildTree func(postorder, inorder []int) *TreeNode
	buildTree = func(postorder, inorder []int) *TreeNode {
		n := len(postorder)
		if n == 0 {
			return nil
		}
		tree := &TreeNode{postorder[n-1], nil, nil}
		var i = n - 1
		for ; i >= 0; i-- {
			if postorder[n-1] == inorder[i] {
				break
			}
		}
		tree.Left = buildTree(postorder[1:len(inorder[:i])+1], inorder[:i])
		tree.Right = buildTree(postorder[len(inorder[:i])+1:], inorder[i+1:])
		return tree
	}
	root := buildTree(postorder, inorder)
	t.Log(root)
}
