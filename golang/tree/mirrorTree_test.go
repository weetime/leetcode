package tree

import (
	"testing"
)

/**
*	镜像二叉树，解题思路，深度搜索，左树=右树， 右树=左树
* dfs 需要注意两点，终止条件，root=nil
* 想象特例，如果root只有两个节点会怎么样，这样就比较容易了
* 那左右节点反转其实也一样了
* 和它相关的题目，判断一颗树是否对称isSymmetric
 */
func mirrorTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		root.Left, root.Right = dfs(root.Right), dfs(root.Left)
		return root
	}
	return dfs(root)
}

/**
* 另外一种解决方案：
* 广度搜索，遍历所有的节点，让根节点的左右节点互换
 */
func mirrorTree1(root *TreeNode) *TreeNode {
	var bfs func(root *TreeNode) *TreeNode
	bfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		return root
	}

	return bfs(root)
}

func TestMirrorTree(t *testing.T) {
	res := mirrorTree(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(res)

	res = mirrorTree1(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(res)
}
