package tree

import "container/list"

func preorder(root *Node) []int {
	var res []int
	var dfs func(*Node)

	if root == nil {
		return nil
	}

	dfs = func(n *Node) {
		res = append(res, n.Val) // 前序遍历
		for i := 0; i < len(n.Children); i++ {
			dfs(n.Children[i])
		}
	}
	dfs(root)
	return res
}

// has bug
func preorder2(root *Node) []int {
	var res []int
	queue := list.New()
	queue.PushBack(root)

	for e := queue.Back(); e != nil; e = e.Next() {
		n := e.Value.(*Node)
		res = append(res, n.Val)
		for _, ne := range n.Children {
			queue.PushBack(ne)
		}
	}
	// length := len(res)
	// for i := 0; i < length/2; i++ {
	// 	res[i], res[length-1-i] = res[length-1-i], res[i]
	// }
	return res
}
