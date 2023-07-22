package stack

import "testing"

type MinStack struct {
	stack1 []int
	stack2 []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack1 = append(this.stack1, x)
	if len(this.stack2) > 0 && this.stack2[len(this.stack2)-1] < x {
		this.stack2 = append(this.stack2, this.stack2[len(this.stack2)-1])
	} else {
		this.stack2 = append(this.stack2, x)
	}
}

func (this *MinStack) Pop() {
	this.stack1 = this.stack1[0 : len(this.stack1)-1]
	this.stack2 = this.stack2[0 : len(this.stack2)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack1) == 0 {
		return -1
	}
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack2) == 0 {
		return -1
	}
	return this.stack2[len(this.stack2)-1]
}

func TestConstructor(t *testing.T) {
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(2)
	// minStack.Push(-3)
	t.Log(minStack.stack1, minStack.stack2)
	t.Log(minStack.Top())
	t.Log(minStack.Min())
	minStack.Pop()
	t.Log(minStack.Min())
	t.Log(minStack.Top())
	t.Log(minStack.stack1, minStack.stack2)
}
