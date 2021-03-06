package tree

// LRD εεΊιε
func maxDepth_n(root *Node) int {
	level := 0
	if root == nil {
		return level
	}
	for i := 0; i < len(root.Children); i++ {
		level = max(level, maxDepth_n(root.Children[i]))
	}

	level++
	return level
}
