package main

import "fmt"

func main() {
	var m, n, a, b int
	fmt.Scan(&m)
	fmt.Scan(&n)
	n++
	m++
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for {
		n, _ := fmt.Scanf("%d %d", &a, &b)
		if n != 0 {
			if b == 0 {
				for i := a; i < m; i++ {
					dp[i][0] = 0
				}
			} else if a == 0 {
				for i := b; i < n; i++ {
					dp[0][i] = 0
				}
			} else {
				dp[a][b] = 0
			}
		} else {
			break
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 && j >= 1 && dp[i][j] != 0 {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
				fmt.Println(i, j, dp[i][j])
			}
		}
	}

	fmt.Println(dp[m - 1][n - 1])
}


// func main() {
// 	var m, n, a, b int
// 	fmt.Scan(&m)
// 	fmt.Scan(&n)
// 	n++
// 	m++
// 	dp := make([]int, n)
//
// 	for {
// 		n, _ := fmt.Scanf("%d %d", &a, &b)
// 		if n != 0 {
// 			dp[a][b] = 0
// 		} else {
// 			break
// 		}
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if i >= 1 && j >= 1 {
// 				dp[j] += dp[j - 1]
// 			}
// 		}
// 	}
//
// 	fmt.Println(dp[n - 1])
// }
