package tree

import "testing"

/**
* 可以采用深度优先搜索，前序遍历整棵树，当到达叶子节点的时候，计算路径和是否等于target
* 注意在每次搜索结束的时候，需要把做出选择的路径放回去
 */
func pathSum(root *TreeNode, target int) [][]int {
	paths := [][]int{}
	if root == nil {
		return paths
	}
	path := []int{}
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		defer func() { // 需要还原path
			path = path[:len(path)-1]
		}()
		sum += root.Val
		if root.Left == nil && root.Right == nil && sum == target {
			paths = append(paths, append([]int{}, path...))
			return
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
	dfs(root, 0)
	return paths
}

func TestPathSum(t *testing.T) {
	res := pathSum(sortedArrayToBST([]int{1, 2, 3}), 5)
	t.Log(res)
}
