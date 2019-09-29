package leetcode

import (
	"fmt"
	"testing"
)

// 0-1 背包 Knapsack
func knapsack(weights, values []int, m int) int {
	n := len(weights)
	dp := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, m + 1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if j < weights[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weights[i-1]] + values[i-1])
			}
		}
	}

	return dp[n][m]
}


// 1049

func TestKnapsack(t *testing.T)  {
	fmt.Println(knapsack([]int{5,4,7,2,6}, []int{12,3,10,3,6}, 10))
}