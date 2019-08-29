package main

import (
	"fmt"
	"sort"
)

var dict map[byte]int

type newString string
type strings []newString

func (s strings) Len() int {
	return len(s)
}
func (s strings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s strings) Less(i, j int) bool {
	s1 := string(s[i])
	s2 := string(s[j])
	if len(s1) == 0 {
		return false
	}
	if len(s1) != 0 && len(s2) == 0 {
		return true
	}
	l1 := len(s1)
	l2 := len(s2)
	var less = true
	var n int
	if l1 > l2 {
		n = l2
	} else {
		n = l1
	}
	for i := 0; i < n; i++ {
		if dict[s1[i]] > dict[s2[i]] {
			less = false
			break
		}
	}

	return less && l1 < l2
}

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"
	arrays := new(strings)
	for _, s := range []string{"cba", "cc", "ccc"} {
		*arrays = append(*arrays, newString(s))
	}
	// n := 0
	// fmt.Scan(&str)
	// fmt.Scan(&n)
	// for i := 0; i < n; i++ {
	// 	s := ""
	// 	fmt.Scan(&s)
	// 	*arrays = append(*arrays, newString(s))
	// }
	dict = make(map[byte]int, 0)
	for i := 0; i < len(str); i++ {
		dict[str[i]] = i
	}

	sort.Sort(arrays)
	for _, s := range *arrays {
		fmt.Println(s)
	}
}
