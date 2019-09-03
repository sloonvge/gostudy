package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, k int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	fmt.Println(maxK(n, m ,k))
}

func maxK(n, m, k int) int {
	a := make([]int, 0)
	for i := 1; i < m + 1; i++ {
		for j := 1; j < n + 1; j++ {
			a = append(a, i * j)
		}
	}
	sort.Ints(a)
	return a[len(a) - k]
}
