package sort

import "testing"

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	nums := []int{6, 5, 4, 3, 2, 1}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	t.Log(nums)
}

// 选择排序
func TestSelectSort(t *testing.T) {
	var nums = []int{6, 5, 4, 3, 1, 2}
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
	t.Log(nums)
}

// 插入排序（排序10000个整数，用时约30ms）
func TestInsertSort(t *testing.T) {
	nums := []int{6, 5, 4, 3, 2, 1}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = temp
		}
	}
	t.Log(nums)
}

func TestQuickSort(t *testing.T) {
	//快速排序（排序10000个随机整数，用时约0.9ms）
	nums := []int{1, 2, 3, 4, 5, 6}

	var recursionSort func(nums []int, left int, right int)
	var partition func(nums []int, left int, right int) int
	recursionSort = func(nums []int, left int, right int) {
		if left < right {
			pivot := partition(nums, left, right)
			recursionSort(nums, left, pivot-1)
			recursionSort(nums, pivot+1, right)
		}
	}

	partition = func(nums []int, left int, right int) int {
		for left < right {
			for left < right && nums[left] <= nums[right] {
				right--
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}

			for left < right && nums[left] <= nums[right] {
				left++
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
				right--
			}
		}
		return left
	}
	recursionSort(nums, 0, len(nums)-1)
	t.Log(nums)
}

type MyInt int

func (a MyInt) min(b MyInt) MyInt {
	if a < b {
		return a
	}
	return b
}

func TestMin(t *testing.T) {
	x := MyInt(20)
	y := x.min(10)
	println(y)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestReverse(t *testing.T) {
	var reverseList func(head *ListNode) *ListNode
	reverseList = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		var new *ListNode
		curr := head
		for curr != nil {
			next := curr.Next
			curr.Next, new = new, curr
			curr = next
		}
		return new
	}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	new := reverseList(head)
	t.Log(new)
}
