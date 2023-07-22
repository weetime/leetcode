package queue

import "testing"

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	var res int = -1
	if len(this.stack2) == 0 {
		l1 := len(this.stack1) - 1
		for l1 >= 0 {
			this.stack2 = append(this.stack2, this.stack1[l1])
			this.stack1 = this.stack1[0:l1]
			l1--
		}
	}

	l2 := len(this.stack2) - 1
	if l2 >= 0 {
		res = this.stack2[l2]
		this.stack2 = this.stack2[0:l2]
	}
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func TestXxx(t *testing.T) {
	// a := []int{1}
	// t.Log(a[0:0])

	obj := Constructor()
	obj.AppendTail(3)
	obj.AppendTail(1)
	param_2 := obj.DeleteHead()
	t.Log(param_2)
	obj.AppendTail(2)
	param_2 = obj.DeleteHead()
	t.Log(param_2)
	param_2 = obj.DeleteHead()
	t.Log(param_2)
}
