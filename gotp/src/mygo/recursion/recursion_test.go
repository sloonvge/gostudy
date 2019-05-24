package recursion

import (
	"testing"
	"math"
)

// Coin Change
func coinChange(N int, denom []int) int{
	ans := math.MaxInt64
	if N <= 0 {
		return 0
	}
	for _, v := range denom {
		if N - v >= 0 {
			ans = min(ans, coinChange(N - v, denom) + 1)
		}
	}
	return ans
}

func min(i int, j int) int{
	if i <= j {
		return i
	} else {
		return j
	}
}

func TestCoinChange(t *testing.T) {
	ans := coinChange(10, []int{1, 3, 4})
	t.Logf("%d\n", ans)
}


// 0-1 背包问题
var N = 5;
var dp map[int]map[int]int
var wts, vals []int

func initBrute() {
	dp = make(map[int]map[int]int)
	wts = make([]int, 0)
	vals = make([]int, 0)
	wts = []int{5, 4, 7, 2, 6}
	vals = []int{12, 3, 10, 3, 6}
	//for i := 1; i < 5; i++ {
	//	wts = append(wts, i)
	//	vals = append(vals, 5 - i)
	//}
}

func brute(weightleft int, idx int) int {
	if idx >= N {
		return 0
	}

	ans := math.MaxInt64
	ans = brute(weightleft, idx + 1)
	if dp[weightleft][idx] != 0 {
		return dp[weightleft][idx]
	}
	if wts[idx] <= weightleft {
		ans = max(ans, vals[idx] + brute(weightleft - wts[idx], idx + 1))
	}
	if _, ok := dp[weightleft]; !ok{
		t := make(map[int]int)
		if _, ok := t[idx]; !ok {
			t[idx] = ans
		}
		dp[weightleft] = t
	}
	// dp[weightleft][idx] = ans
	return ans
}

func max(i int, j int) int{
	if i <= j {
		return j
	} else {
		return i
	}
}

func TestKnapsackProblem(t *testing.T) {
	initBrute()
	ans := brute(10, 0)
	t.Logf("0-1背包\n%d\n", ans)
	t.Logf("%v", dp)
}