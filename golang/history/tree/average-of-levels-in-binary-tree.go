package tree

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	tmp := map[int][]int{}
	if root == nil {
		return res
	}
	var dfs func(*TreeNode, int)
	var avg func([]int) float64
	dfs = func(tn *TreeNode, level int) {
		if tn != nil {
			tmp[level] = append(tmp[level], tn.Val)
			dfs(tn.Left, level+1)
			dfs(tn.Right, level+1)
		}
	}
	avg = func(arr []int) float64 {
		length := len(arr)
		var sum float64
		if length == 0 {
			return 0
		}
		for _, val := range arr {
			sum += float64(val)
		}
		return sum / float64(length)
	}
	dfs(root, 0)

	for i := 0; i < len(tmp); i++ {
		res = append(res, avg(tmp[i]))
	}

	return res
}

func averageOfLevels2(root *TreeNode) []float64 {
	result := make([]float64, 0)
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		current := queue
		sum := 0.0
		queue = make([]*TreeNode, 0) // 重置一下queu
		for j := 0; j < len(current); j++ {
			sum += float64(current[j].Val)
			if current[j].Left != nil {
				queue = append(queue, current[j].Left)
			}
			if current[j].Right != nil {
				queue = append(queue, current[j].Right)
			}
		}
		result = append(result, sum/float64(len(current)))
	}

	return result
}
