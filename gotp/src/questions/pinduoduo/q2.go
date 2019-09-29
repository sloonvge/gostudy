package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	numbers := make([]int, n + 1)
	prev := 0
	cur := 0
	for i := 1; i <= m ; i++ {
		a := 0
		b := 0
		fmt.Scanf("%d %d", &a, &b)
		cur = 0
		for j := a; j <= b; j++ {
			if numbers[j] == 0 {
				numbers[j] = 1
				cur++
			}
		}
		fmt.Println(prev + cur)
		prev += cur
	}
}
