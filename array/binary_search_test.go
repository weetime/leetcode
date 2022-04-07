package array

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	intArray := []int{1, 2, 3, 4, 5, 6, 7, 100, 102, 203, 205, 280, 289, 290, 303}
	// intArray := []int{4}
	res := search(intArray, 303, 0, len(intArray)-1)
	fmt.Println(res)
}

func search(input []int, pick int, min, max int) int {
	if max < min {
		return -1
	}

	mid := ((max - min) >> 1) + min

	if input[mid] == pick {
		return mid
	} else if input[mid] > pick {
		return search(input, pick, 0, mid-1)
	} else {
		return search(input, pick, mid+1, max)
	}
}
