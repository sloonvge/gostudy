package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		s := ""
		fmt.Scan(&s)
		isQQPhone(s, n)
	}
}

func isQQPhone(s string, n int) {
	flag := helper(s, n)
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func helper(s string, n int) bool {
	if len(s) < 11 {
		return false
	}
	for i := 0; i < n; i++ {
		if s[i] == '8' {
			if n - i < 11 {
				fmt.Println(n, i)
				return false
			} else {
				return true
			}
		}
	}

	return true
}
