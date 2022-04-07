package tree

// sumRootToLeaf2 N杈树的路径求和
func sumRootToLeaf2(root *Node) int {
	road, total := 0, 0
	if root == nil {
		return 0
	}
	var dfs func(*Node, int)
	dfs = func(tn *Node, road int) { // 每次在上一次的基础上增加当前的值，所有road要一直透传下去
		if tn != nil {
			road = road<<1 + tn.Val
			if tn.Children == nil {
				total += road
			}
			for i := 0; i < len(tn.Children); i++ {
				dfs(tn.Children[i], road)
			}
		}
	}
	dfs(root, road)

	return total
}
