package array

import (
	"fmt"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	res := strings.Split(s, " ")
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	trim_res := []string{}
	for i := 0; i < len(res); i++ {
		if res[i] != "" {
			fmt.Println(res[i])
			trim_res = append(trim_res, res[i])
		}
	}
	return strings.Join(trim_res, " ")
}

func TestReverseWords(t *testing.T) {
	res := reverseWords("This             is test.")
	t.Log(res)
}
