package tree

import "testing"

func TestFindMode(t *testing.T) {
	// [1,nil,2,2]
	root := &TreeNode{ // 初始化BFS
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	// 题解：BST LDR中序遍历 得到的是一个升序的数组，众数出现为连续的一组相同的数
	// count 用于统计数字的频次，maxCount 用于统计最多出现次数 base为连续数据的值
	// 设计一个update函数，如果入参val=base 累计count,如果val!=base 更新base为val，同时重置count为0
	// 如果count == maxCount 将新的base存入到结果集中
	// 如果count > maxCount，则更新maxCount 重置结果集为base

	var solution func(root *TreeNode) []int // 解决方案呢
	var dfs func(root *TreeNode)            // LDR
	var update func(val int)                // 更新结果集
	var res []int                           // 结果集
	var base, count, maxCount int           // 基数，频次，最高频次

	update = func(val int) {
		if val == base { // 值相等，累加计数器
			count++
		} else { // 值不相等 重新开始计算
			base, count = val, 1
		}

		if count == maxCount { // 累加到最大频次，更新结果集
			res = append(res, base)
		} else if count > maxCount { // 新的最高频次出现
			maxCount = count
			res = []int{base}
		}
	}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		update(root.Val)
		dfs(root.Right)
	}
	solution = func(root *TreeNode) []int {
		dfs(root)
		return res
	}
	t.Log(solution(root))
}
