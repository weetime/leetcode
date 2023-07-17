package list

import "testing"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
* 解题思路:需要深度复制，由于存在Random，可能这个节点还没有创建，所以在遇到的时候，采用递归去创建即可
* 由于Random可能已经创建过了，可以采用hashmap记录一下，如果遇到了直接返回
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	var table = make(map[*Node]*Node)
	var copy func(head *Node) *Node
	copy = func(head *Node) *Node {
		if head == nil {
			return head
		}
		if n, ok := table[head]; ok {
			return n
		}
		new_node := &Node{
			Val: head.Val,
		}
		table[head] = new_node
		new_node.Next = copy(head.Next)
		new_node.Random = copy(head.Random)
		return new_node
	}
	return copy(head)
}

func TestCopyRandomList(t *testing.T) {
	n1 := &Node{
		Val: 1,
	}
	n2 := &Node{
		Val: 2,
	}
	n1.Next = n2
	n1.Random = n2
	n2.Random = n2
	res := copyRandomList(n1)
	t.Log(res)
}
