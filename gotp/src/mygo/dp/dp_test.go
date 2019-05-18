package dp

import (
	"fmt"
	"testing"
)

/** 
* Created by wanjx in 2019/3/30 23:08
**/

// https://blogarithms.github.io/articles/2019-03/cracking-dp-part-one
//
//
// 1 2 3
// 4 5 6
// 7 8 9
// f(node) = val(node) + max(f(upper), f(left))

var dp [][]int
var m, n int
var A = [][]int{
	{2, 5, 1},
	{4, 6, 7},
}

func Init() {
	m = 3
	n = 2
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp =append(dp, make([]int, n))
	}
	fmt.Printf("%v\n", dp)
	dp[0][0] = A[0][0]
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func TestGrid(t *testing.T) {
	// Init()
	n = 2
	m = 3
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		if dp[i] == nil {
			dp[i] = make([]int, m)
		}
	}

	dp[0][0] = A[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 && j != 0{
				dp[i][j] = dp[i][j - 1] + A[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i - 1][j] + A[i][j]
			} else {
				dp[i][j] = max(dp[i][j - 1], dp[i - 1][j]) + A[i][j]
			}
		}
	}
	fmt.Printf("%v\n", dp)
	fmt.Printf("%v\n", dp[n-1][m-1])
}