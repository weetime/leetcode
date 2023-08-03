package tree

import (
	"strconv"
	"strings"
	"testing"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	data := []string{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			data = append(data, "nil")
			return
		}
		data = append(data, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(data, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	dataArr := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		st := dataArr[0]
		dataArr = dataArr[1:]
		if st == "nil" {
			return nil
		}

		val, _ := strconv.Atoi(st)
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}

func TestConstructor(t *testing.T) {
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(SortedArrayToBST([]int{1, 2, 3, 4, 5}))
	ans := deser.deserialize(data)
	t.Log(ans)
}
