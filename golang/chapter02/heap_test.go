package chapter02

import (
	"fmt"
	"testing"
)

// func TestHeap(t *testing.T) {
// 	list := map[string]int{ // map 统计全文单次频次
// 		"This":   1,
// 		"is":     20,
// 		"me":     5,
// 		"Thanks": 8,
// 		"good":   2,
// 		"luck":   8,
// 		"he":     9,
// 		"done":   6,
// 		"like":   7,
// 		"can":    100,
// 		"test":   22,
// 		"etc":    34,
// 	}
// 	arr := make(Words, len(list))
// 	for n, v := range list {
// 		arr = append(arr, Word{
// 			Count: v,
// 			Name:  n,
// 		})
// 	}

// 	// 堆排序
// 	heapSort(arr, 0, arr.Len())
// 	// 取top 10
// 	fmt.Println(arr[:10])
// }

// type Words []Word
// type Word struct {
// 	Count int
// 	Name  string
// }

// // sort 接口
// func (a Words) Len() int      { return len(a) }
// func (a Words) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// func (a Words) Less(i, j int) bool {
// 	return a[i].Count > a[j].Count // 保障稳定性 只有大于才交换 不是>=
// }

// // 构建堆
// func heapSort(data Words, a, b int) {
// 	first := a // 起始指针
// 	lo := 0    // 最小指针
// 	hi := b    // 最大指针

// 	for i := hi - 1; i >= 0; i-- {
// 		data.Swap(first, first+i)
// 		siftDown(data, lo, i, first)
// 	}
// }

// func siftDown(data Words, lo, hi, first int) {
// 	root := lo
// 	for {
// 		child := 2*root + 1 // 从中值之后构建
// 		if child >= hi {
// 			break
// 		}
// 		if child+1 < hi && data.Less(first+child, first+child+1) {
// 			child++
// 		}

// 		if !data.Less(first+root, first+child) {
// 			return
// 		}

// 		data.Swap(first+root, first+child)
// 		root = child
// 	}
// }

func TestHeap(t *testing.T) {
	words := map[string]int{ // map 统计全文单次频次
		"This":   1,
		"is":     20,
		"me":     5,
		"Thanks": 8,
		"good":   2,
		"luck":   8,
		"he":     9,
		"done":   6,
		"like":   7,
		"can":    100,
		"test":   22,
		"etc":    34,
	}

	// 构建最大堆
	h := NewMinHeap(10)
	for name, count := range words {
		h.Push(Word{
			Count: count,
			Name:  name,
		})
	}

	// 将堆元素移除
	for range h.Array {
		t.Log(h.Pop())
	}

	// 打印排序后的值
	fmt.Println(h.Array)
}

type Word struct {
	Count int
	Name  string
}

// 一个最大堆，一棵完全二叉树
// 最大堆要求节点元素都不小于其左右孩子
type Heap struct {
	// 堆的大小
	Size int
	// 堆容量
	Capacity int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []Word
}

// 初始化一个堆
func NewMinHeap(capacity int) *Heap {
	h := new(Heap)
	h.Capacity = capacity
	h.Array = make([]Word, capacity)
	return h
}

// 最大堆插入元素
func (h *Heap) Push(x Word) {
	h.Size = len(h.Array)
	// 堆没有元素时，使元素成为顶点后退出
	if h.Size == 0 {
		h.Array[h.Size] = x
		h.Size++
		return
	}

	if h.Size == h.Capacity {
		if x.Count < h.Array[0].Count {
			return
		}
		h.Pop()
	}

	// i 是要插入节点的下标
	i := h.Size

	// 如果下标存在
	// 将小的值 x 一直上浮
	for i > 0 {
		// parent为该元素父亲节点的下标
		parent := (i - 1) / 2

		// 如果插入的值小于等于父亲节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x.Count >= h.Array[parent].Count {
			break
		}

		// 否则将父亲节点与该节点互换，然后向上翻转，将最大的元素一直往上推
		h.Array[i] = h.Array[parent]
		i = parent
	}

	// 将该值 x 放在不会再翻转的位置
	h.Array[i] = x

}

// 最大堆移除根节点元素，也就是最大的元素
func (h *Heap) Pop() Word {
	// 没有元素，返回-1
	if h.Size == 0 {
		return Word{}
	}

	// 取出根节点
	ret := h.Array[0]

	// 因为根节点要被删除了，将最后一个节点放到根节点的位置上
	h.Size--
	x := h.Array[h.Size]  // 将最后一个元素的值先拿出来
	h.Array[h.Size] = ret // 将移除的元素放在最后一个元素的位置上

	// 对根节点进行向下翻转，小的值 x 一直下沉，维持最大堆的特征
	i := 0
	for {
		// a，b为下标 i 左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 左儿子下标超出了，表示没有左子树，那么右子树也没有，直接返回
		if a >= h.Size {
			break
		}

		// 有右子树，拿到两个子节点中较大节点的下标
		if b < h.Size && h.Array[b].Count < h.Array[a].Count {
			a = b
		}

		// 父亲节点的值都大于或等于两个儿子较大的那个，不需要向下继续翻转了，返回
		if x.Count <= h.Array[a].Count {
			break
		}

		// 将较大的儿子与父亲交换，维持这个最大堆的特征
		h.Array[i] = h.Array[a]

		// 继续往下操作
		i = a
	}

	// 将最后一个元素的值 x 放在不会再翻转的位置
	h.Array[i] = x
	return ret
}
