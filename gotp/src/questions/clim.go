package questions

import "fmt"

// func main() {
// 	fmt.Println(climbStairs(7))
// }

func climbStairs(n int) int {

	dp := make(map[int]int, 0)
	dp[1] = 1
	dp[2] = 2
	return climb(n, dp)
}

func climb(n int, dp map[int]int) int {
	fmt.Printf("n:%d-%v\n", n, dp)

	if _, ok := dp[n]; ok {
		return dp[n]
	}
	dp[n] = climb(n - 1, dp) + climb(n - 2, dp)
	return dp[n]
}