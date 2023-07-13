package chapter01

import (
	"testing"
)

func TestFib(t *testing.T) {
	t.Log(Fib1(10))
	t.Log(Fib2(10))
	t.Log(Fib3(10))
	t.Log(Fib4(10))
}

func Fib1(n int) int {
	if n <= 2 {
		return 1
	}

	return Fib1(n-2) + Fib1(n-1)
}

func Fib2(n int) int {
	if n <= 2 {
		return 1
	}
	sum, pre, current := 0, 1, 1

	for i := 3; i <= n; i++ {
		sum = pre + current
		pre, current = current, sum
	}

	return sum
}

var fibMemo = map[int]int{}

func Fib3(n int) int {
	if n <= 2 {
		return 1
	}
	if res, ok := fibMemo[n]; ok {
		return res
	} else {
		fibMemo[n] = Fib3(n-2) + Fib3(n-1)
	}
	return fibMemo[n]
}

var fibTmp = map[int]int{
	1: 1,
	2: 1,
}

func Fib4(n int) int {

	for i := 3; i <= n; i++ {
		fibTmp[i] = fibTmp[i-2] + fibTmp[i-1]
	}

	return fibTmp[n]
}
