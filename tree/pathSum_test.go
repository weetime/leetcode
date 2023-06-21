package tree

import "testing"

func TestPathSum(t *testing.T) {
	var root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	var targetSum = 22
	var routes = [][]int{}
	var path = []int{}
	var dfs func(root *TreeNode, sum int) // DLR 遍历
	dfs = func(root *TreeNode, sum int) {
		if root == nil { // 递归终止条件
			return
		}
		path = append(path, root.Val) // 将值放入到路径中
		defer func() {                // 每次回溯结束的时候哦回到上个位置
			path = path[:len(path)-1]
		}()
		sum += root.Val                                                // 累加值
		if root.Left == nil && root.Right == nil && sum == targetSum { // 到达叶子节点的时候处理逻辑
			// routes = append(routes, path) 这种写法是错误的是，这种复制是按引用复制，path如果修改会影响route
			routes = append(routes, append([]int{}, path...)) // 按值复制 或者 copy 深度复制
		}
		dfs(root.Left, sum)
		dfs(root.Right, sum)
	}
	dfs(root, 0)
	t.Log(routes)
}
