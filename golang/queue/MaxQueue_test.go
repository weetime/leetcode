package queue

import "testing"

type MaxQueue struct {
	q1  []int
	max []int
}

func ConstructorMaxQueue() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q1 = append(this.q1, value)
	// 维持单调递减队列，一个个比对，注意重复值的情况
	for i := 0; i < len(this.max); i++ {
		if value > this.max[i] {
			this.max = this.max[:i]
			break
		}
	}
	// for len(this.max) != 0 && value > this.max[len(this.max)-1] {
	// 	this.max = this.max[:len(this.max)-1]
	// }
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q1) != 0 {
		n = this.q1[0]
		this.q1 = this.q1[1:]
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}
	return n
}

func TestMaxQueue(t *testing.T) {
	obj := ConstructorMaxQueue()
	param_1 := obj.Max_value()
	obj.Push_back(3)
	obj.Push_back(1)
	obj.Push_back(10)
	obj.Push_back(5)
	obj.Push_back(10)
	obj.Push_back(9)
	obj.Push_back(6)
	// max => [10,10,9,6]
	param_2 := obj.Max_value()
	obj.Pop_front()
	obj.Pop_front()
	obj.Pop_front()
	param_3 := obj.Pop_front()
	param_4 := obj.Max_value()

	t.Log(param_1, param_2, param_3, param_4)
}
