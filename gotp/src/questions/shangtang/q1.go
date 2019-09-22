package main
//
// import "fmt"
//
// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	if n < 2 {
// 		fmt.Println(0)
// 		return
// 	}
// 	if n == 2 {
// 		fmt.Println(3)
// 		return
// 	}
// 	if n == 3 {
// 		fmt.Println(4)
// 		return
// 	}
// 	dp := make([]int, n + 1)
// 	nums := make([]int, n + 1)
// 	dp[0] = 0
// 	nums[0] = 0
// 	dp[1] = 1
// 	nums[1] = 1
// 	dp[2] = 2
// 	nums[2] = 2
// 	dp[3] = 3
// 	nums[3] = 2
// 	max := 0
// 	for i := 4; i <= n; i++ {
// 		max = 0
// 		for j := 1; j <= i / 2; j++ {
// 			temp := dp[j] * dp[i - j]
// 			if temp > max {
// 				max = temp
// 				nums[i] = nums[j] + nums[i - j]
// 			}
// 			dp[i] = max
// 		}
// 	}
//
// 	fmt.Println(dp[n] + nums[n] / 2)
// }
