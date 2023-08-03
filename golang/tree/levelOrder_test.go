package tree

import "testing"

/**
* 如果按层级输出，只放到一个队列中，更简单了，只用广度搜索就可以了
 */
func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		res = append(res, node.Val)
	}
	return res
}

/**
*	解题思路：层级输出，需要用一个队列存储每层的节点，那如何知道每一层有多少节点呢。
* 方式1：遍历一层的时候，插入一个空节点，这样相当于埋入了一个特殊标记
* 每次遍历到这个特殊标记的时候，就知道换了一层了，最后输出一个二维数组即可
* */
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue, root, nil)
	level := 0
	temp := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
			res = append(res, temp)
			temp = []int{}
			level++
			continue
		}
		temp = append(temp, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
	return res
}

/**
*	方式2：其实可以不引入这个特殊标记的，A.队列遍历之前是知道当前是第几层的，全部拿出来即可。
* B.还可以每次拿去队列之前，放到另外一个队列里面
* */
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		res = append(res, temp)

	}
	return res
}

/**
*	如果奇数的时候反转呢，那就加一个反转就完事了
* */
func levelOrder3(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		temp := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		if level%2 == 1 {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
		res = append(res, temp)
	}
	return res
}

func TestLevelOrder(t *testing.T) {
	r := levelOrder(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(r)

	r1 := levelOrder1(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(r1)

	r2 := levelOrder2(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(r2)

	r3 := levelOrder3(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6}))
	t.Log(r3)
}
