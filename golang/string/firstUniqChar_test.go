package string

import (
	"math"
	"testing"
)

/**
* 解题的关键：
* 首次，只出现一次
* 传统的解决方案，二次遍历，存哈希表，第二次遍历的时候，如果出现次数为1就返回
* 另外一种解决方案，遍历一次字符串，存哈希表，value为字符串首次出现的索引，第二次出现就修改value=-1
 */
func firstUniqChar(s string) byte {
	table := make(map[rune]int)
	for i, v := range s {
		if _, ok := table[v]; ok {
			table[v] = -1
		} else {
			table[v] = i
		}
	}
	min := math.MaxInt
	for _, v := range table {
		if v >= 0 && min > v {
			min = v
		}
	}
	if min != math.MaxInt {
		return s[min]
	}

	return ' '
}

func TestFirstUniqChar(t *testing.T) {
	res := firstUniqChar("abaccde")
	t.Log(res)
}
