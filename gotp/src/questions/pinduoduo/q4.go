package main

import "fmt"

func main() {
	var t, n, m int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &m)
		numbers := make([]int, n)
		weeks := make([]int, m)
		for j := 0; j < m; j++ {
			x := 0
			fmt.Scan(&x)
			weeks[j] = x
		}
		for j := 0; j < m; j++ {
			x := 0
			fmt.Scan(&x)
			numbers[weeks[j] - 1] = x
		}

		total := 1
		for i := 0; i < n; i++ {
			if numbers[i] == 0 && i > 0 && i < n - 1{
				if numbers[i-1] == 0 && numbers[i+1] == 0 {
					total *= 3
				} else if numbers[i-1] != 0 && numbers[i+1] == 0 {
					total *= 3
				} else if numbers[i-1] == 0 && numbers[i+1] != 0 {
					temp := total
					total *= 3
					total = total -  temp / 3 * 2
				} else {
					if numbers[i-1] == numbers[i+1] {
						total *= 3
					} else {
						total *= 2
					}
				}
			}
			if numbers[i] == 0 && i == 0 {
				if numbers[i+1] != 0 {
					total *= 3
				} else {
					total *= 4
				}
			}
			if numbers[i] == 0 && i == n - 1 {
				total *= 3
			}
			total = total % 1000000007
		}

		if total == 1 {
			total = 0
		}
		fmt.Println(total)
	}
}
