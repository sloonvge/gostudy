package main

import "fmt"

func main() {
	var m, n, x int
	fmt.Scanf("%d %d", &n, &m)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		numbers[i] = x
	}
	sum := 0
	for i := 0; i < m; i++ {
		sum += numbers[i] * numbers[2 * m - 1 - i]
	}

	fmt.Println(sum)
}
