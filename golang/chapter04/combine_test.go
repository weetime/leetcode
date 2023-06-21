package chapter04

import "testing"

func TestCombine(t *testing.T) {
	var res [][]int
	var trackBack func(n, k, start int, trace []int)
	trackBack = func(n, k, start int, trace []int) {
		if len(trace) == k {
			res = append(res, append([]int{}, trace...))
			return
		}
		for i := start; i <= n; i++ {
			trace = append(trace, i)
			trackBack(n, k, i+1, trace)
			trace = trace[:len(trace)-1]
		}
	}
	trackBack(4, 2, 1, []int{})
	t.Log(res)
}
