package main

import (
	"fmt"
	"strings"
)

func main() {
	var S int
	fmt.Scan(&S)
	poker := ""
	ns := make([]string, 0)
	ms := make([]string, 0)
	result := make([][]string, S)
	for i := 0; i < 2 * S; i++ {
		fmt.Scanln(&poker)
		if i & 1 == 0 {
			ns = append(ns, poker)
		} else {
			ms = append(ms, poker)
		}
	}

	for i := 0; i < S; i++ {
		p1 := ns[i]
		p2 := ms[i]
		result[i] = playPoker(p1, p2)
	}

	printResult(result)
}

func playPoker(p1, p2 string) (ret []string){
	if len(p1) != len(p2) {
		poker := make([]byte, 0)
		pre := ""
		for i := 0; i < len(p1); i++ {
			if !strings.Contains(p2, string(p1[i])) {
				continue
			}
			cur := string(p1[i])
			if left, ok := isLeft(p2, pre, cur); ok {
				if left {
					poker = append(poker, 'l')
				} else {
					poker = append(poker, 'r')
				}
				pre = cur
			}
		}
		s := string(poker)
		if len(s) == len(p2) {
			ret = append(ret, s)
		}
	} else {
		poker := make([]byte, 0)
		pre := ""
		for i := 0; i < len(p1); i++ {
			if !strings.Contains(p2, string(p1[i])) {
				continue
			}
			cur := string(p1[i])
			if left, ok := isLeft(p2, pre, cur); ok {
				if left {
					poker = append(poker, 'l')
				} else {
					poker = append(poker, 'r')
				}
				pre = cur
			}
		}
		s := string(poker)
		if len(s) == len(p2) {
			ret = append(ret, s)
		}
	}
	return
}

func isLeft(s string, pre, cur string) (left bool, ok bool) {
	if !strings.Contains(s, cur) {
		return false, false
	}
	index := strings.Index(s, pre)
	if index < len(s) - 1 {
		if string(s[index + 1]) == cur {
			return false, true
		} else {
			return false, false
		}
	}
	if index > 0 {
		if string(s[index - 1]) == cur {
			return true, true
		} else {
			return false, false
		}
	}

	return
}

func printResult(a [][]string) {
	for _, pokers := range a {
		fmt.Printf("{\n")
		for _, poker := range pokers {
			for _, s := range poker {
				fmt.Printf("%s ", string(s))
			}
		}
		fmt.Printf("}\n")
	}
}
