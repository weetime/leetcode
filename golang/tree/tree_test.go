package tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	res := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	lrdBST(res)
}

func TestMaxDepth(t *testing.T) {
	t.Log(maxDepth(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
	t.Log(maxDepth2(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
}

func TestPostorder(t *testing.T) {
	var root, a, b, c = &Node{}, &Node{}, &Node{}, &Node{}
	a.Val = 1
	b.Val = 2
	c.Val = 3
	root.Val = 0
	b.Children = append(b.Children, a)
	c.Children = append(c.Children, b, a)
	root.Children = append(root.Children, c)
	t.Log(postorder(root))
	t.Log(postorder2(root))

	t.Log(preorder(root))
	// t.Log(preorder2(root)) // has bug
}

func TestSearchBST(t *testing.T) {
	root := sortedArrayToBST([]int{1, 3, 2, 4, 7})
	t.Log(searchBST(root, 2))
}

func TestIncreasingBST(t *testing.T) {
	var a, b, c, d, e, root = &TreeNode{Val: 1}, &TreeNode{Val: 2}, &TreeNode{Val: 3}, &TreeNode{Val: 4}, &TreeNode{Val: 6}, &TreeNode{Val: 5}
	b.Left = a
	c.Left = b
	c.Right = d
	root.Left = c
	root.Right = e
	t.Log(increasingBST(root))
}

func TestMaxDepth_n(t *testing.T) {
	var root, a, b, c = &Node{}, &Node{}, &Node{}, &Node{}
	a.Val = 1
	b.Val = 2
	c.Val = 3
	root.Val = 0
	b.Children = append(b.Children, a)
	c.Children = append(c.Children, b, a)
	root.Children = append(root.Children, c)
	t.Log(maxDepth_n(root))
}

func TestSumRootToLeaf(t *testing.T) {
	t.Log(sumRootToLeaf(sortedArrayToBST([]int{0, 0, 1, 1, 0, 1, 1})))
	// dlrBST(sortedArrayToBST([]int{0, 0, 1, 1, 0, 1, 1}))
}

// N杈树的路径求和
func TestSumRootToLeaf2(t *testing.T) {
	var root, a, b, c = &Node{}, &Node{}, &Node{}, &Node{}
	a.Val = 0
	b.Val = 1
	c.Val = 0
	root.Val = 1
	b.Children = append(b.Children, a)
	c.Children = append(c.Children, b, a)
	root.Children = append(root.Children, c)
	t.Log(sumRootToLeaf2(root))
}

func TestAverageOfLevels(t *testing.T) {
	t.Log(averageOfLevels(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
	t.Log(averageOfLevels2(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
}
