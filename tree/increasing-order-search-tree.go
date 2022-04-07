package tree

func increasingBST(root *TreeNode) *TreeNode {
	var curr, sentinel *TreeNode
	var inorder func(*TreeNode)
	sentinel = new(TreeNode)
	curr = sentinel
	inorder = func(tn *TreeNode) {
		if tn != nil {
			inorder(tn.Left)
			tn.Left = nil
			curr.Right = tn
			curr = curr.Right
			inorder(tn.Right)
		}
	}
	inorder(root)
	return sentinel.Right
}
