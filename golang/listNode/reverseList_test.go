package listnode

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	var solution func(head *ListNode) *ListNode
	solution = func(head *ListNode) *ListNode {
		new := &ListNode{} // 重新初始化新链表
		for head != nil {
			current := &ListNode{ // 新增一个节点
				Val:  head.Val,
				Next: nil,
			}
			// 0->1->2
			// 0->2 => 1->0
			current.Next = new.Next
			new.Next = current
			head = head.Next
		}
		return new.Next
	}
	t.Log(solution(GetListNode()))
}

func TestReverseList2(t *testing.T) {
	var solution func(head *ListNode) *ListNode
	solution = func(head *ListNode) *ListNode {
		var new *ListNode
		current := head
		for current != nil {
			next := current.Next
			current.Next, new = new, current
			current = next
		}
		return new
	}
	t.Log(solution(GetListNode()))
}

func TestReverse(t *testing.T) {
	var solution func(*ListNode) *ListNode
	solution = func(ln *ListNode) *ListNode {
		var new *ListNode
		current := ln
		for current != nil {
			next := current.Next
			// 0->1;0->nil;new=0->nil
			// 1->2;1->0;new=1->0->nil
			// 2->3;2->1;new=2->1->0->nil
			current.Next, new = new, current
			current = next
		}

		return new
	}
	t.Log(solution(GetListNode()))
}

func TestReverse2(t *testing.T) {
	head := GetListNode()
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		last := dfs(head.Next)
		head.Next.Next = head
		head.Next = nil
		return last
	}
	l := dfs(head)
	t.Log(l)
}

func TestMergeList(t *testing.T) {
	l1 := GetListNode()
	l2 := GetListNode()
	var curr = &ListNode{}
	var new = curr
	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			break
		} else if l2 == nil {
			curr.Next = l1
		}
		if l1 != nil && l2 != nil && l1.Val >= l2.Val {
			// println("This is v2", l2.Val)
			curr.Next = &ListNode{
				Val: l2.Val,
			}
			curr = curr.Next
			l2 = l2.Next
			// curr.Next = nil
		} else {
			// println("This is v1", l1.Val)
			curr.Next = &ListNode{
				Val: l1.Val,
			}
			curr = curr.Next
			l1 = l1.Next
		}
	}

	println("This is list node", new.Next)
}

func TestMergeList2(t *testing.T) {
	l1, l2 := GetListNode(), GetListNode()
	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(ln1, ln2 *ListNode) *ListNode {
		if ln1 == nil {
			return ln2
		}
		if ln2 == nil {
			return ln1
		}
		if ln1.Val < ln2.Val {
			ln1.Next = merge(ln1.Next, ln2)
			return ln1
		} else {
			ln2.Next = merge(ln1, ln2.Next)
			return ln2
		}
	}
	res := merge(l1, l2)
	println(res)
}

func TestDeleteList(t *testing.T) {
	l1 := GetListNode()
	curr := l1
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}

	}
	println(l1)
}

func TestHasCycle(t *testing.T) {
	head := GetListNode()
	var hasCycle func(head *ListNode) bool
	hasCycle = func(head *ListNode) bool {
		fast, slow := head, head
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		}
		return false
	}
	res := hasCycle(head)
	println(res)
}

func TestGetIntersectionNode3(t *testing.T) {
	h1, h2 := GetListNode(), GetListNode()
	var getIntersectionNode func(headA, headB *ListNode) *ListNode
	getIntersectionNode = func(headA, headB *ListNode) *ListNode {
		if headA == nil || headB == nil {
			return nil
		}
		n1 := headA
		n2 := headB
		for n1 != nil || n2 != nil {
			if n1 == n2 {
				return n1
			}
			if n1 == nil {
				n1 = headB
			} else {
				n1 = n1.Next
			}
			if n2 == nil {
				n2 = headA
			} else {
				n2 = n2.Next
			}
		}
		return nil
	}
	res := getIntersectionNode(h1, h2)
	println(res)
}

func TestRemove_elements(t *testing.T) {
	var remove_elements func(head *ListNode, val int) *ListNode
	head := GetListNode()
	remove_elements = func(head *ListNode, val int) *ListNode {
		if head == nil {
			return head
		}

		sentinel := &ListNode{Next: head}
		curr, prev := head, sentinel
		for curr != nil {
			if curr.Val == val {
				prev.Next = curr.Next
			} else {
				prev = curr
			}
			curr = curr.Next
		}
		return sentinel.Next
	}
	res := remove_elements(head, 1)
	println("{:?}", res)

	// sentnal := &ListNode{}
	// sentnal.Next = head
	// prev, curr := sentnal, head
	// for curr != nil {
	// 		if curr.Val == val {
	// 				prev.Next = curr.Next
	// 		} else {
	// 				prev = curr
	// 		}
	// 		curr = curr.Next
	// }
	// return sentnal.Next
}

func TestDeleteDuplicates(t *testing.T) {
	var deleteDuplicates func(head *ListNode) *ListNode
	head := GetListNode()

	deleteDuplicates = func(head *ListNode) *ListNode {
		var duplicates = head.Val
		sentinel := &ListNode{Val: 0, Next: head}
		fast, slow := head, sentinel

		for fast.Next != nil {
			if duplicates != fast.Next.Val {
				duplicates = fast.Next.Val
				slow = slow.Next
				fast = fast.Next
			} else {
				println("this is in", duplicates)
				fast = fast.Next
				slow.Next = slow.Next.Next
			}

		}
		return sentinel.Next
	}

	res := deleteDuplicates(head)
	t.Log(res)
}

func TestDeleteDuplicates2(t *testing.T) {
	var deleteDuplicates func(head *ListNode) *ListNode
	head := GetListNode()

	deleteDuplicates = func(head *ListNode) *ListNode {
		var tmp = make(map[int]struct{})
		sentinel := &ListNode{Val: 0, Next: head}
		curr, prev := head, sentinel

		for curr != nil {
			if _, ok := tmp[curr.Val]; ok || (curr.Next != nil && curr.Val == curr.Next.Val) {
				prev.Next = curr.Next
				tmp[curr.Val] = struct{}{}
			} else {
				prev = curr
			}
			curr = curr.Next
		}
		return sentinel.Next
	}

	res := deleteDuplicates(head)
	t.Log(res)
}

func TestDeleteDuplicates3(t *testing.T) {
	var deleteDuplicates func(head *ListNode) *ListNode
	head := GetListNode()

	// deleteDuplicates = func(head *ListNode) *ListNode {
	// 	if head == nil {
	// 		return head
	// 	}
	// 	sentinel := &ListNode{Val: 0, Next: head}
	// 	curr, prev := head, sentinel
	// 	for curr != nil && curr.Next != nil {
	// 		if curr.Val == curr.Next.Val {
	// 			val := curr.Val
	// 			for curr != nil && curr.Val == val {
	// 				prev.Next = curr.Next
	// 				curr = curr.Next
	// 			}
	// 		} else {
	// 			prev = curr
	// 			curr = curr.Next
	// 		}
	// 	}
	// 	return sentinel.Next
	// }

	deleteDuplicates = func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}
		sentinel := &ListNode{Val: 0, Next: head}
		curr := sentinel
		for curr.Next != nil && curr.Next.Next != nil {
			if curr.Next.Val == curr.Next.Next.Val {
				val := curr.Next.Val
				for curr.Next != nil && curr.Next.Val == val {
					curr.Next = curr.Next.Next
				}
			} else {
				curr = curr.Next
			}
		}
		return sentinel.Next
	}

	res := deleteDuplicates(head)
	t.Log(res)
}

// 快慢指针
func TestDeleteDuplicates4(t *testing.T) {
	var deleteDuplicates func(head *ListNode) *ListNode
	head := GetListNode()

	deleteDuplicates = func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}
		sentinel := &ListNode{Val: 0, Next: head}
		fast, slow := head, sentinel          // 快慢指针，相当于有个prev，先慢一拍
		for fast != nil && fast.Next != nil { // 可以到fast is nill 或者 fast.Next is nil, 但都不影响slow，因为到了最后一个节点，没人和它相同了
			if fast.Val == fast.Next.Val { // 快指针每次和后面的一个值对比
				val := fast.Val
				for fast.Next != nil && fast.Next.Val == val { // 开启新的循环，只要后面还有节点，并且值也是和上一个值相等的就继续往前遍历
					fast = fast.Next
				}
				slow.Next = fast.Next // 慢指针下一个节点 等于值不重复的节点
			} else {
				slow = slow.Next // 正常往前移动
			}
			fast = fast.Next // 不管怎样 快指针都有往前挪
		}
		return sentinel.Next
	}

	res := deleteDuplicates(head)
	t.Log(res)
}

// 双指针 优雅版本
func TestDeleteDuplicates5(t *testing.T) {
	var deleteDuplicates func(head *ListNode) *ListNode
	head := GetListNode()

	deleteDuplicates = func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}
		sentinel := &ListNode{Val: 0, Next: head}
		fast, slow := head, sentinel // 快慢指针，相当于有个prev，先慢一拍
		for fast != nil {
			for fast.Next != nil && fast.Val == fast.Next.Val {
				fast = fast.Next
			}
			if slow.Next == fast {
				slow = slow.Next
			} else {
				slow.Next = fast.Next
			}
			fast = fast.Next
		}
		return sentinel.Next
	}

	res := deleteDuplicates(head)
	t.Log(res)
}

func TestSorList(t *testing.T) {

	var sortList func(head *ListNode) *ListNode
	sortList = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		var mid *ListNode
		slow, fast := head, head.Next
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		mid, slow.Next = slow.Next, nil

		left, right := sortList(mid), sortList(head) // 这一行不能拆成两行，顺序也不能变，先mid 然后 head

		sentinel := new(ListNode)
		h := sentinel
		for left != nil && right != nil {
			if left.Val < right.Val {
				h.Next, left = left, left.Next
			} else {
				h.Next, right = right, right.Next
			}
			h = h.Next
		}
		if left != nil {
			h.Next = left
		} else {
			h.Next = right
		}

		return sentinel.Next
	}

	res := sortList(GetListNode())
	t.Log(res)
}

// 自顶向下归并 空间复杂度O(logn),时间复杂度O(nlogn)
func TestSortList2(t *testing.T) {
	var sort func(head *ListNode) *ListNode
	sort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		slow, fast := head, head.Next
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}
		var mid *ListNode
		mid, slow.Next = slow.Next, nil
		left, right := sort(head), sort(mid)
		dummy := new(ListNode)
		new := dummy
		for left != nil && right != nil {
			if left.Val <= right.Val {
				new.Next, left = left, left.Next
			} else {
				new.Next, right = right, right.Next
			}
			new = new.Next
		}
		if left != nil {
			new.Next = left
		} else {
			new.Next = right
		}

		return dummy.Next
	}
	res := sort(GetListNode())
	t.Log(res)
}

func TestSort2(t *testing.T) {
	var sort func(head *ListNode) *ListNode
	var merge func(head_a, head_b *ListNode) *ListNode
	merge = func(head_a, head_b *ListNode) *ListNode {
		var dummy = &ListNode{}
		new := dummy
		for head_a != nil && head_b != nil {
			if head_a.Val <= head_b.Val {
				new.Next, head_a = head_a, head_a.Next
			} else {
				new.Next, head_b = head_b, head_b.Next
			}
			new = new.Next
		}
		if head_a != nil {
			new.Next = head_a
		} else {
			new.Next = head_b
		}

		return dummy.Next
	}
	sort = func(head *ListNode) *ListNode {
		var length = 0
		var head_p = head
		for head_p != nil {
			head_p = head_p.Next
			length++
		}
		var size = 1
		var res = &ListNode{Next: head}
		for size < length {
			p, curr := res, res.Next
			for curr != nil {
				h1 := curr
				for i := 1; i < size && curr.Next != nil; i++ {
					curr = curr.Next
				}
				h2 := curr.Next
				curr.Next = nil
				curr = h2
				for i := 1; i < size && curr != nil && curr.Next != nil; i++ {
					curr = curr.Next
				}
				var next *ListNode
				if curr != nil {
					next = curr.Next
					curr.Next = nil
				}

				p.Next = merge(h1, h2)
				for p.Next != nil {
					p = p.Next
				}

				curr = next
			}
			size *= 2
		}

		return res.Next
	}

	node := sort(GetListNode())
	t.Log(node)
}

// [3,1,2,5,4]
// 自底向上归并
// step1 size=1切分[3],[1] [2],[5] [4] merge后[1,3] [2,5] [4] 得到的新的链表[1,3,2,5,4] //前2个已经排序完成
// step2 size=2切分[1,3],[2,5] [4] merge后[1,2,3,5] [4] 得到新的链表[1,2,3,5,4] // 前4个已经排序完成
// step3 size=4切分[1,2,3,5] [4] merge后[1,2,3,4,5] 排序完成
func TestSortList3(t *testing.T) {
	head := GetListNode()
	var cut func(head *ListNode, n int) (*ListNode, *ListNode)
	// 切分后 返回切分的节点 和 剩余的链表
	cut = func(head *ListNode, n int) (*ListNode, *ListNode) {
		var new *ListNode
		dummy := &ListNode{Next: head}
		p := dummy
		for p.Next != nil && n > 0 {
			p = p.Next
			n--
		}
		new, p.Next = p.Next, nil
		return dummy.Next, new
	}

	// 合并有序链表
	var merge func(head_a, head_b *ListNode) *ListNode
	merge = func(head_a, head_b *ListNode) *ListNode {
		var dummy = &ListNode{}
		new := dummy
		for head_a != nil && head_b != nil {
			if head_a.Val <= head_b.Val {
				new.Next, head_a = head_a, head_a.Next
			} else {
				new.Next, head_b = head_b, head_b.Next
			}
			new = new.Next
		}
		if head_a != nil {
			new.Next = head_a
		} else {
			new.Next = head_b
		}

		return dummy.Next
	}

	var sort func(head *ListNode) *ListNode

	sort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		var length = 0
		var head_p = head
		for head_p != nil {
			head_p = head_p.Next
			length++ // 统计链表长度
		}
		var size = 1

		for size < length {
			var res = &ListNode{}
			p, curr := res, head
			for curr != nil {
				var left, right *ListNode
				left, curr = cut(curr, size)
				right, curr = cut(curr, size)
				p.Next = merge(left, right)
				for p.Next != nil { // 这里需要注意每次需要把p移动到队列结尾
					p = p.Next
				}
			}
			head = res.Next
			size *= 2
		}
		return head
	}

	node := sort(head)
	t.Log(node)
}

func TestInsertionSortList(t *testing.T) {
	sort := func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}
		dummy := &ListNode{Next: head}
		lastSorted, curr := head, head.Next
		for curr != nil {
			if lastSorted.Val <= curr.Val {
				lastSorted = lastSorted.Next
			} else {
				prev := dummy
				for prev.Next.Val <= curr.Val {
					prev = prev.Next
				}
				lastSorted.Next = curr.Next
				curr.Next = prev.Next
				prev.Next = curr
			}

			curr = lastSorted.Next
		}
		return dummy.Next
	}

	res := sort(GetListNode())
	t.Log(res)
}

// 这个算法没有上面的好，如果是排序好的最后一个，可以先cache住，lastSorted，能提升速度
// 如果有已经排好序的就不用排了，这个算法还需要遍历一遍
func TestInsertionSortList2(t *testing.T) {
	sort := func(head *ListNode) *ListNode {
		if head == nil {
			return head
		}
		dummy, curr := &ListNode{}, head
		for curr != nil {
			var next, prev *ListNode
			next, curr.Next, prev = curr.Next, nil, dummy
			for prev.Next != nil && prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			curr.Next, prev.Next, curr = prev.Next, curr, next
		}
		return dummy.Next
	}

	res := sort(GetListNode())
	t.Log(res)
}

func TestOddEvenList(t *testing.T) {
	head := GetListNode()
	curr := head
	p1, p2 := &ListNode{}, &ListNode{}
	h1, h2 := p1, p2
	var n = 0
	for curr != nil {
		next := curr.Next
		curr.Next = nil
		if n%2 == 0 {
			h1.Next = curr
			h1 = h1.Next
		} else {
			h2.Next = curr
			h2 = h2.Next
		}
		n++
		curr = next
	}
	h1.Next = p2.Next
	t.Log(p1.Next)
}

func TestOddEvenList2(t *testing.T) {
	head := GetListNode()
	odd := head
	evenNode := head.Next
	even := evenNode
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenNode
	t.Log(head)
}

func TestMergeInBetween(t *testing.T) {
	mergeInBetween := func(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
		head := list1
		var n = 1
		for list1 != nil {
			if n >= a && n <= b {
				prev := list1
				for n >= a && n <= b {
					list1 = list1.Next
					n++
				}
				p2 := list2
				prev.Next = list2
				for p2.Next != nil {
					p2 = p2.Next
				}
				p2.Next = list1.Next
				// break
			} else {
				list1 = list1.Next
				n++
			}
		}
		return head
	}

	res := mergeInBetween(GetListNode(), 3, 3, GetListNode())
	t.Log(res)
}

func TestMergeInBetween2(t *testing.T) {
	mergeInBetween := func(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
		prevA := list1
		for i := 0; i < a-1; i++ {
			prevA = prevA.Next
		}

		prevB := prevA
		for i := 0; i < b-a+2; i++ {
			prevB = prevB.Next
		}
		prevA.Next = list2
		for list2.Next != nil {
			list2 = list2.Next
		}
		list2.Next = prevB
		return list1
	}

	res := mergeInBetween(GetListNode(), 3, 3, GetListNode())
	t.Log(res)
}

func TestDeleteMiddle(t *testing.T) {
	deleteMiddle := func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return nil
		}

		slow, fast := head, head
		var prev *ListNode
		for fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			prev = slow
			slow = slow.Next
		}
		prev.Next = slow.Next
		return head
	}
	res := deleteMiddle(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					// Next: &ListNode{
					// 	Val: 5,
					// },
				},
			},
		},
	})
	t.Log(res)
}

func TestAddTwoNumbers(t *testing.T) {
	var addTwoNumbers func(l1 *ListNode, l2 *ListNode) *ListNode
	addTwoNumbers = func(l1, l2 *ListNode) *ListNode {
		var carry int
		var newListNode = &ListNode{}
		new := newListNode
		for l1 != nil || l2 != nil || carry > 0 {
			n1, n2 := 0, 0
			if l1 != nil {
				n1 = l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				n2 = l2.Val
				l2 = l2.Next
			}
			res := n1 + n2 + carry
			res, carry = res%10, res/10
			new.Next = &ListNode{Val: res}
			new = new.Next
		}

		return newListNode.Next
	}

	l1, l2 := GetListNode(), &ListNode{Val: 9}
	res := addTwoNumbers(l1, l2)
	t.Log(res)
}

func TestMod(t *testing.T) {
	t.Log(9/10, 9%10)
}

// func TestReverse3(t *testing.T) {
// 	head := GetListNode()
// 	var dfs func(head *ListNode) *ListNode
// 	dfs = func(head *ListNode) *ListNode {
// 		if head == nil || head.Next == nil {
// 			return head
// 		}
// 		last := dfs(head.Next)
// 		head.Next.Next = head
// 		head.Next = nil
// 		return last
// 	}
// 	res := dfs(head)
// 	t.Log(res)
// }

func TestSum(t *testing.T) {
	a := []int{}
	a = append(a, 1)
}

// 头插法、尾插法构建链表
func TestBuild(t *testing.T) {
	arr := []int{1, 2, 3}
	var head *ListNode
	for _, val := range arr {
		curr := &ListNode{Val: val, Next: head}
		head = curr
	}

	t.Log(head)

	var dummy = &ListNode{}
	tail := dummy
	for _, val := range arr {
		curr := &ListNode{Val: val}
		tail.Next = curr
		tail = tail.Next
	}

	t.Log(dummy.Next)

	var ln, ptr *ListNode
	for _, val := range arr {
		if ln == nil {
			ln = &ListNode{Val: val}
			ptr = ln
		} else {
			ptr.Next = &ListNode{Val: val}
			ptr = ptr.Next
		}
	}

	t.Log(ln)
}

func TestSwapPairs(t *testing.T) {
	head := GetListNode()
	dummy := &ListNode{Val: 0, Next: head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	t.Log(dummy.Next)
}

func TestRotateRight(t *testing.T) {
	rotateRight := func(head *ListNode, k int) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		ptr := head
		var length int
		for ptr != nil {
			length++
			ptr = ptr.Next
		}

		k = k % length

		for i := 0; i < k; i++ {
			var prev *ListNode
			ptr := head
			for ptr != nil && ptr.Next != nil {
				prev = ptr
				ptr = ptr.Next
			}
			if prev != nil {
				prev.Next = nil
				ptr.Next = head
				head = ptr
			}
		}
		return head
	}

	res := rotateRight(GetListNode(), 1)
	t.Log(res)
}

func TestRotateRight1(t *testing.T) {
	// 闭合为环 比较经典
	rotateRight := func(head *ListNode, k int) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		ptr := head
		var length int = 1
		for ptr.Next != nil {
			length++
			ptr = ptr.Next
		}

		k = length - (k % length)
		ptr.Next = head
		for i := 0; i < k; i++ {
			ptr = ptr.Next
		}
		res := ptr.Next
		ptr.Next = nil
		return res
	}

	res := rotateRight(GetListNode(), 1)
	t.Log(res)
}

func TestRotateRight2(t *testing.T) {
	rotateRight := func(head *ListNode, k int) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		ptr := head
		var length int = 1
		for ptr.Next != nil {
			length++
			ptr = ptr.Next
		}

		k = k % length
		if k == 0 {
			return head
		}
		// 切割链表
		fast, slow := head, head
		for k >= 0 {
			fast = fast.Next
			k--
		}

		for fast != nil {
			fast = fast.Next
			slow = slow.Next
		}

		new_head := slow.Next
		slow.Next = nil
		tail := new_head
		for tail.Next != nil {
			tail = tail.Next
		}

		tail.Next = head
		return new_head
	}

	res := rotateRight(GetListNode(), 2)
	t.Log(res)
}
