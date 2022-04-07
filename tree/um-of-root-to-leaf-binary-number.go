package tree

func sumRootToLeaf(root *TreeNode) int {
	road, total := 0, 0
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, road int) { // 每次在上一次的基础上增加当前的值，所有road要一直透传下去
		if tn != nil {
			road = road<<1 + tn.Val
			if tn.Left == nil && tn.Right == nil {
				total += road
			}
			dfs(tn.Left, road)
			dfs(tn.Right, road)
		}
	}
	dfs(root, road)

	return total
}
