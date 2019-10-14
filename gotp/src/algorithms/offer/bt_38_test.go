package offer

import (
	"fmt"
	"testing"
)

func Permutation(s string) {
	permutation([]byte(s), 0)
}
func permutation(s []byte, i int) {
	if i == len(s) {
		fmt.Println(string(s))
		return
	}

	for j := i; j < len(s); j++ {
		s[i], s[j] = s[j], s[i]
		permutation(s, i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func TestPermutation(t *testing.T) {
	Permutation("abc")
}