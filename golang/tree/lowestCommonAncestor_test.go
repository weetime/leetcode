package tree

import "testing"

/*
* 由于是二叉搜索数，所以能每次剪枝减一半
*	最近公共祖先节点，解题思路，递归
* root.Val 比两个节点值都大，则在root.Left 搜索
* root.Val 比两个节点值都小，则在root.Right 搜索
* 递归到nil 或者 值相等，或者p,q分别在root两侧，则返回root
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}
		if root.Val > p.Val && root.Val > q.Val {
			return dfs(root.Left, p, q)
		} else if root.Val < p.Val && root.Val < q.Val {
			return dfs(root.Right, p, q)
		}
		return root
	}
	return dfs(root, p, q)
}

/*
* 另外一种解法，遍历一次
 */
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	ans := root
	for {
		if ans.Val > p.Val && ans.Val > q.Val {
			ans = ans.Left
		} else if ans.Val < p.Val && ans.Val < q.Val {
			ans = ans.Right
		} else {
			return ans
		}
	}
}

/*
* 非二叉搜索树的解法
* p,q 要么在root.Left，要么在root.Right，要么就是一边一个
* 采用递归的方式，如果左侧为nil，那么就在右侧，同理就在左侧
* 如果各匹配了一个，那就是root本身
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}
		left := dfs(root.Left, p, q)
		right := dfs(root.Right, p, q)
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		return root
	}
	return dfs(root, p, q)
}

func TestLowestCommonAncestor(t *testing.T) {
	res := lowestCommonAncestor(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}), SortedArrayToBST([]int{1}), SortedArrayToBST([]int{6}))
	t.Log(res)

	res = lowestCommonAncestor1(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}), SortedArrayToBST([]int{1}), SortedArrayToBST([]int{6}))
	t.Log(res)
}
