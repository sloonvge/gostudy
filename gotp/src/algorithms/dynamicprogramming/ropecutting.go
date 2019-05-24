package main

import "fmt"

/** 
* Created by wanjx in 2019/4/27 15:20
**/
 
// Question: 如何把长度为n的绳子截成若干段,使得各段的乘积最大?
func RopeCut(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 || n == 3{
		return 2
	}

	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	m := 0
	for i := 4; i <= n; i++ {
		m = 0
		for j := 0; j <= i / 2; j++ {
			p := dp[j] * dp[i - j]
			m = max(p, m)
		}
		dp[i] = m
	}
	m = dp[n]

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	v := RopeCut(8)

	fmt.Println(v)
}