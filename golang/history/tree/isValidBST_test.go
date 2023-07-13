package tree

import "testing"

func TestIsValidBST(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var isValidBST func(root *TreeNode) bool
	var valid func(root, min, max *TreeNode) bool
	isValidBST = func(root *TreeNode) bool {
		return valid(root, nil, nil)
	}

	valid = func(root, min, max *TreeNode) bool {
		if root == nil {
			return true
		}

		// Left 如果 >= 父节点的值 则是false
		if min != nil && root.Val >= min.Val {
			return false
		}

		// Right 如果 <= 父节点的值 则返回false
		if max != nil && root.Val <= max.Val {
			return false
		}

		// 交给递归方法
		return valid(root.Left, root, max) && valid(root.Right, min, root)
	}
	t.Log(isValidBST(tree))
}

func TestAB(t *testing.T) {
	// signFlag := false
	// if dividend < 0 {
	// 	signFlag = !signFlag
	// 	dividend = -dividend
	// }
	// if divisor < 0 {
	// 	signFlag = !signFlag
	// 	divisor = -divisor
	// }
	// if dividend < divisor {
	// 	return 0
	// }
	// tmp := 1
	// tmpValue := divisor
	// value := 0
	// res := 0
	// resMap := map[int]int{1: divisor}
	// for {
	// 	if value+resMap[tmp]*2 < dividend {
	// 		if _, ok := resMap[tmp*2]; !ok {
	// 			resMap[tmp*2] = resMap[tmp] * 2
	// 		}
	// 		tmp *= 2
	// 		tmpValue = resMap[tmp]
	// 		continue
	// 	}
	// 	value += tmpValue
	// 	res += tmp
	// 	if value == dividend {
	// 		break
	// 	}
	// 	tmp = 1
	// 	tmpValue = resMap[tmp]
	// 	if value+resMap[tmp] > dividend {
	// 		break
	// 	}
	// }
	// if signFlag {
	// 	if res > 2147483648 {
	// 		return 214748483647
	// 	}
	// 	return -res
	// }
	// if res > 2147483647 {
	// 	return 2147483647
	// }
	// return res
}

func TestA2B(t *testing.T) {
	// a, b := 6, 2
	// var res = make(map[int]int)
	// tmp = 1
	// res[tmp] = 1
	// var value = 0
	// for {
	// 	if b+res[tmp]*2 < a {
	// 		if _, ok := res[tmp*2]; !ok {
	// 			res[tmp*2] = res[tmp] * 2
	// 		}
	// 		tmp *= 2
	// 		tmpValue = resMap[tmp]
	// 		continue
	// 	}

}
