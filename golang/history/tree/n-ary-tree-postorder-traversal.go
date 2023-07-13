package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return nil
	}

	order(root, &res)

	return res
}

func order(root *Node, res *[]int) {
	for i := 0; i < len(root.Children); i++ {
		order(root.Children[i], res)
	}

	*res = append(*res, root.Val)
}

func postorder2(root *Node) []int {
	var res []int
	var dfs func(*Node)
	dfs = func(n *Node) {
		for i := 0; i < len(n.Children); i++ {
			dfs(n.Children[i])
		}

		res = append(res, n.Val)
	}
	dfs(root)
	return res
}
