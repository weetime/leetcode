package tree

func sortedArrayToBST(nums []int) *TreeNode {
	return toBST(nums, 0, len(nums))
}

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums)
}

func toBST(arr []int, s int, e int) *TreeNode {
	if s == e {
		return nil
	}
	m := (s + e) >> 1
	t := new(TreeNode)
	t.Val = arr[m]
	t.Left = toBST(arr, s, m)
	t.Right = toBST(arr, m+1, e)

	return t
}
