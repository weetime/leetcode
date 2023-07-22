package moni

import "testing"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)

		for len(stack) > 0 {
			if stack[len(stack)-1] == popped[i] {
				stack = stack[:len(stack)-1]
				i++
			} else {
				break
			}
		}

		// for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
		// 	stack = stack[:len(stack)-1]
		// 	i++
		// }
	}

	return len(stack) == 0
}

func TestValidateStackSequences(t *testing.T) {
	res := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 2, 1})
	t.Log(res)
}
