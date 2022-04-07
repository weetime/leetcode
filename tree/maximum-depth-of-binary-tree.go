package tree

func maxDepth(root *TreeNode) int {
	ret := 0
	if root == nil {
		return 0
	}

	ret++ // 永远是1

	lHight := maxDepth(root.Left)
	rHight := maxDepth(root.Right)

	if lHight > rHight {
		return ret + lHight
	}
	return ret + rHight
}

// LRD 后序遍历 最最大值
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lHight := maxDepth2(root.Left)
	rHight := maxDepth2(root.Right)
	depth := max(lHight, rHight)
	depth++
	return depth
}
