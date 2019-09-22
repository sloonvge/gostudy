package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 64
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
	vis := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		vis[i] = make([]bool, cols)
	}
	minPath := math.MaxInt64
	pathSum2(grid, vis, &minPath, 0, rows - 1, cols - 1, 0, 0)

	return minPath
}

func pathSum2(grid [][]int, vis [][]bool, minPath *int, sum, rows, cols, row, col int) {
	if row == rows && col == cols {
		sum += grid[row][col]
		if sum < *minPath {
			*minPath = sum
		}
		return
	}
	sum += grid[row][col]
	if row > 0 && !vis[row - 1][col] {
		vis[row - 1][col] = true
		pathSum2(grid, vis, minPath, sum, rows, cols, row - 1, col)
		vis[row - 1][col] = false
	}
	if row < rows && !vis[row + 1][col] {
		vis[row + 1][col] = true
		pathSum2(grid, vis, minPath, sum, rows, cols, row + 1, col)
		vis[row + 1][col] = false
	}
	if col > 0 && !vis[row][col - 1] {
		vis[row][col - 1] = true
		pathSum2(grid, vis, minPath, sum, rows, cols, row, col - 1)
		vis[row][col - 1] = false
	}
	if col < cols && !vis[row][col + 1] {
		vis[row][col + 1] = true
		pathSum2(grid, vis, minPath, sum, rows, cols, row, col + 1)
		vis[row][col + 1] = false
	}

}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	min := func(a, b int) int{
		if a < b {
			return a
		}
		return b
	}
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	var a, b int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > 0 {
				a = dp[i - 1][j]
			} else {
				a = 0
			}
			if j > 0 {
				b = dp[i][j - 1]
			} else {
				b = 0
			}
			if i > 0 && j > 0 {
				dp[i][j] = min(a, b) + grid[i][j]
			} else {
				dp[i][j] = a + b + grid[i][j]
			}
		}
	}

	fmt.Println(dp[rows - 1][cols - 1])
	return dp[rows - 1][cols - 1]
}


func TestMinPathSum(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}})
	})
	t.Run("2", func(t *testing.T) {
		minPathSum2([][]int{{1,4,8,6,2,2,1,7},{4,7,3,1,4,5,5,1},{8,8,2,1,1,8,0,1},{8,9,2,9,8,0,8,9},{5,7,5,7,1,8,5,5},{7,0,9,4,5,6,5,6},{4,9,9,7,9,1,9,0}})
	})
}

// 494
func findTargetSumWays1(nums []int, S int) int {
	n := 0
	targetSum(nums, 0, &n, 0, S)
	fmt.Println(n)
	return n
}

func targetSum(nums []int, i int, n *int, s, S int) {
	if i == len(nums){
		if s == S {
			*n++
		}
		return
	}
	s += nums[i]
	targetSum(nums, i + 1, n, s, S)
	s -= 2 * nums[i]
	targetSum(nums, i + 1, n, s, S)
}

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum < S || (S + sum) % 2 > 0 {
		return 0
	}
	return sebsetSum(nums, (S + sum) >> 1)
}
func sebsetSum(nums []int, S int) int {
	dp := make([]int, S + 1)
	dp[0] = 1
	for _, n := range nums {
		for i := S; i >= n; i-- {
			dp[i] += dp[i - n]
		}
	}

	return dp[S]
}

func TestFindTargetSumWays(t *testing.T) {
	findTargetSumWays([]int{1,1,1,1,1}, 3)
}

// 5
// O(n^3) pass:876ms 2.3MB
func longestPalindrome2(s string) string {
	l := 0
	sl := ""
	for j := 1; j <= len(s); j++ {
		for i := 0; i <= len(s) - j; i++ {
			if isPalindrome(s[i:i+j]) {
				if j > l {
					l = j
					sl = s[i:i+j]
				}

			}
		}
	}

	return sl
}

func isPalindrome(s string) bool {
	var flag = true
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - 1 - i] {
			flag = false
			break
		}
	}

	return flag
}

// O(N^2) 48ms 6.5MB
func longestPalindrome3(s string) string {
	var ret string
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for j := 1; j <= len(s); j++ {
		for i := 0; i <= len(s) - j; i++ {
			if s[i] == s[i + j] && (j < 3 || dp[i + 1][i + j - 1]) {
				dp[i][i + j] = true
			}
			if dp[i][i + j] && j > len(ret) {
				ret = s[i:i+j]
			}
		}
	}

	return ret
}

func longestPalindrome4(s string) string {
	var ret string
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j - i < 3 || dp[i + 1][j - 1]) {
				dp[i][j] = true
			}
			if dp[i][j] && j - i + 1 > len(ret) {
				ret = s[i:j+1]
			}
		}
	}

	return ret
}

// O(N^2) 4ms 2.2MB
func longestPalindrome(s string) string {
	var lo, max int
	n := len(s)
	if n < 2 {
		return s
	}
	for i := 0; i < n - 1; i++ {
		extendPalindrome(s, i, i, lo, max)
		extendPalindrome(s, i, i + 1, lo, max)
	}

	return s[lo:lo + max]
}
func extendPalindrome(s string, i, j, lo, max int) {
	for i > 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	if j - i + 1 > max {
		max = j - i + 1
		lo = i
	}
}

// 152
//
func maxProduct1(nums []int) int {
	max := math.MinInt64
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			p := 1
			for _, num := range nums[i:j+1] {
				p *= num
			}
			if p > max {
				max = p
			}
		}
	}

	return max
}

// O(N^2)
func maxProduct2(nums []int) int {
	max := math.MinInt64
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := 1
		for j := i; j < n; j++ {
			temp *= nums[j]
			if temp > max {
				max = temp
			}
		}

	}

	return max
}

func maxProduct3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	maxs := make([]int, n)
	mins := make([]int, n)
	mins[0], maxs[0] = nums[0], nums[0]
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			maxs[i], mins[i] = mins[i], maxs[i]
		}
		maxs[i] = int(math.Max(float64(maxs[i - 1] * nums[i]), float64(nums[i])))
		mins[i] = int(math.Min(float64(mins[i - 1] * nums[i]), float64(nums[i])))
		max = int(math.Max(float64(maxs[i]), float64(max)))
	}

	return max
}

func maxProduct(nums []int) int {
	max := nums[0]
	for i, imax, imin := 1, max,  max; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		if nums[i] > imax * nums[i] {
			imax = nums[i]
		} else {
			imax = imax * nums[i]
		}
		if nums[i] < imin * nums[i] {
			imin = nums[i]
		} else {
			imin = imin * nums[i]
		}

		if imax > max {
			max = imax
		}
	}

	return max
}

// 322
func coinChange2(coins []int, amount int) int {
	dp := make(map[int]int)
	ans :=  dynamic(coins, amount, dp)
	if ans == math.MaxInt32 {
		ans = -1
	}
	return ans
}
func dynamic(coins []int, amount int, dp map[int]int) int {
	if value, ok := dp[amount]; ok {
		return value
	}
	ans := math.MaxInt32
	if amount <= 0 {
		return 0
	}
	for _, coin := range coins {
		if amount - coin >= 0 {
			t := 1 + dynamic(coins, amount - coin, dp)
			if t < ans {
				ans = t
			}
		}
	}
	dp[amount] = ans

	return ans
}

func coinChange3(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount)
	return minimumCoins(coins, amount, dp)
}
func minimumCoins(coins []int, amount int, dp []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if dp[amount - 1] != 0 {
		return dp[amount - 1]
	}
	min := math.MaxInt64
	for _, coin := range coins {
		val := minimumCoins(coins, amount - coin, dp)
		if val >= 0 && val < min {
			min = 1 + val
		}
	}
	if min == math.MaxInt64 {
		min = -1
	}
	dp[amount - 1] = min

	return dp[amount - 1]
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		var val int
		for _, coin := range coins {
			if i >= coin {
				val = dp[i - coin]
				if val < dp[i] {
					dp[i] = val
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 221
func maximalSquare1(matrix [][]byte) int {
	max := 0

	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				area := 1
				all := true
				for area + i < m && area + j < n && all{
					for k := i; k <= i + area; k++ {
						if matrix[k][j + area] != '1' {
							all = false
							break
						}
					}
					if all {
						for k := j; k <= j + area; k++ {
							if matrix[i + area][k] != '1' {
								all = false
								break
							}
						}
					}
					if all {
						area++
					}
				}
				if area > max {
					max = area
				}
			}
		}
	}

	return max * max
}

func maximalSquare2(matrix [][]byte) int {
	maxSquareLen := 0
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	dp := make([][]int, rows + 1)
	for i := 0; i <= rows; i++ {
		dp[i] = make([]int, cols + 1)
	}
	min := func(a, b, c int) int{
		m := a
		if m > b {
			m = b
		}
		if m > c {
			m = c
		}

		return m
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i - 1][j - 1] == '1' {
				dp[i][j] = min(dp[i - 1][j], dp[i - 1][j - 1], dp[i][j - 1]) + 1
				if dp[i][j] > maxSquareLen {
					maxSquareLen = dp[i][j]
				}
			}
		}
	}

	return maxSquareLen * maxSquareLen
}

func maximalSquare(matrix [][]byte) int {
	maxSquareLen := 0
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	dp := make([]int, cols + 1)
	min := func(a, b, c int) int{
		m := a
		if m > b {
			m = b
		}
		if m > c {
			m = c
		}

		return m
	}
	prev := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			temp := dp[j]
			if matrix[i - 1][j - 1] == '1' {
				dp[j] = min(dp[j], prev, dp[j - 1]) + 1
				if dp[j] > maxSquareLen {
					maxSquareLen = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}

	return maxSquareLen * maxSquareLen
}

// 300
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
		return a
	}
		return b
	}
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n;  i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}
	answer := 0
	for _, n := range dp {
		answer = max(answer, n)
	}

	return answer
}


// 416
// func canPartition1(nums []int) bool {
// 	return helper(nums, 0, 0, 0)
// }
// func helper(nums []int, curpos, taken, noTaken int) bool {
// 	if curpos == len(nums) {
// 		if taken == noTaken {
// 			return true
// 		}
// 		return false
// 	}
// 	return helper(nums, curpos + 1, taken + nums[curpos], noTaken) ||
// 		helper(nums, curpos + 1, taken, noTaken + nums[curpos])
// }

func canPartition2(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum & 1 == 1 {
		return false
	}
	sum /= 2

	n := len(nums)
	dp := make([][]bool, n + 1)
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]bool, sum + 1)
	}

	dp[0][0] = true
	for i := 1; i < n + 1; i++ {
		dp[i][0] = true
	}
	for j := 1; j < sum + 1; j++ {
		dp[0][j] = false
	}

	for i := 1; i < n + 1; i++ {
		for j := 1; j < sum + 1; j++ {
			dp[i][j] = dp[i - 1][j]
			if j > nums[i] {
				dp[i][j] = dp[i][j] || dp[i - 1][j - nums[i]]
			}
		}
	}

	return dp[n][sum]
}

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum & 1 == 1 {
		return false
	}
	sum /= 2

	dp := make([]bool, sum + 1)
	dp[0] = true
	for _, n := range nums {
		for i := sum; i >= sum; i-- {
			dp[i] = dp[i] || dp[i-n]
		}
	}

	return dp[sum]
}

// 279
func numSquares(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		min := math.MaxInt64
		j := 1
		for i - j * j >= 0 {
			v := dp[i - j * j] + 1
			if v < min {
				min = v
			}
			j++
		}
		dp[i] = min
	}

	return dp[n]
}

// 309 Best Time to Buy and Sell Stock with Cooldown
func maxProfit(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var buy, rest, sell int
	buy = math.MinInt64
	for _, p := range prices {
		prevBuy := buy
		buy = max(rest - p, buy)
		rest = max(sell, rest)
		sell = prevBuy + p
	}

	return max(sell, rest)
}

// 62
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 && j >= 1 {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}

	return dp[m - 1][n - 1]
}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i >= 1 && j >= 1 {
				dp[j] = dp[j - 1] + dp[j]
			}
		}
	}

	return dp[n - 1]
}

// 198 House Robber
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	robs := make([]int, n)
	noRobs := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	robs[0] = nums[0]
	for i := 1; i < n; i++ {
		robs[i] = noRobs[i - 1] + nums[i]
		noRobs[i] = max(noRobs[i - 1], robs[i - 1])
	}

	return max(robs[n - 1], noRobs[n - 1])
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	robs := nums[0]
	noRobs := 0
	for i := 1; i < n; i++ {
		prevRobs := robs
		robs = noRobs + nums[i]
		noRobs = max(noRobs, prevRobs)
	}

	return max(robs, noRobs)
}


func TestT(t *testing.T) {
	fmt.Println(rob([]int{7,1}))
}