package chapter03

import "testing"

// 用于存储数据
type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

// 双向链表
type DoubleListNode struct {
	head *Node
	tail *Node
	size int
}

func (d *DoubleListNode) add(node *Node) {
	node.next = d.tail
	node.prev = d.tail.prev
	d.tail.prev.next = node
	d.tail.prev = node
	d.size++
}

func (d *DoubleListNode) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	d.size--
}

func (d *DoubleListNode) removeFirst() *Node {
	first := d.head.next
	if first == d.tail {
		return nil
	}
	d.remove(first)
	return first
}

func (d *DoubleListNode) Size() int {
	return d.size
}

type LRUCache struct {
	hashMap map[int]*Node
	cache   *DoubleListNode
	cap     int
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:     capacity,
		hashMap: make(map[int]*Node),
		cache: &DoubleListNode{
			head: head,
			tail: tail,
			size: 0,
		},
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.hashMap[key]; ok {
		l.makeRecently(key) // 将元素添加到哦队列尾部
		return node.val
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.hashMap[key]; ok {
		l.deleteKey(key) // key 对应的val 可能有变化，得把老的删了,1=>2 变成了1=>3
		l.addRecently(key, value)
		return
	}
	if l.cap == l.cache.Size() {
		l.removeLeastRecently()
	}
	l.addRecently(key, value)
}

// 将某个key提升为最近使用
func (l *LRUCache) makeRecently(key int) {
	node := l.hashMap[key]
	l.cache.remove(node)
	l.cache.add(node)
}

// 添加一个key
func (l *LRUCache) addRecently(key, val int) {
	node := &Node{
		key: key,
		val: val,
	}
	l.hashMap[key] = node
	l.cache.add(node)
}

// 删除一个key
func (l *LRUCache) deleteKey(key int) {
	node := l.hashMap[key]
	l.cache.remove(node)
	delete(l.hashMap, key)
}

// 删除最久未使用key
func (l *LRUCache) removeLeastRecently() {
	node := l.cache.removeFirst()
	if node != nil {
		delete(l.hashMap, node.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRU(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	t.Log(cache.Get(1))
	cache.Put(3, 3)
	t.Log(cache.Get(2))
	t.Log(cache.Get(3))
}
