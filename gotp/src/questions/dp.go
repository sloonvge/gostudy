package questions

/**
* Created by wanjx in 2019/4/5 20:45
**/


var N = 4
var wts = []int{1, 2, 2, 2, 2, 4, 8}
var vals = []int{1, 2, 2, 2, 2, 1, 2}
// var dp map[int]map[int]int
var dp = make(map[int]int, 0)
var MAX int

// func main()  {
//
// 	ma := max(wts)
// 	MAX = ma
// 	fmt.Println(f(MAX))
// 	fmt.Printf("%v\n", dp)
// }

func f(sum int) int{

	if _, ok := dp[sum]; ok {
		return dp[sum]
	}
	ans := 10000000
	if sum <= 0 {
		return 0
	}
	for _, x := range wts {
		if sum - x >= 0 {
			ans = min(ans, 1 + f(sum - x))
		}
	}
	dp[sum] = ans
	return ans
}

// ans = min(ans, coin[idx] + f(left-value[idx], idx + 1))
// func f(got int, idx int) int{
// 	if idx >= N {
// 		return 0
// 	}
//
// 	if _, ok := dp[got]; !ok {
// 		dp[got] = make(map[int]int)
// 	}
//
//
// 	if _, ok := dp[got][idx]; ok {
// 		return dp[got][idx]
// 	}
//
// 	ans := 10*10000
// 	// ans = f(got, idx+1)
// 	if got < MAX {
// 		ans = min(ans, vals[idx] + f(got + wts[idx], idx + 1))
// 	}
// 	dp[got][idx] = ans
// 	return ans
// }

func max(a []int) int {
	m := 0
	for _, x := range a {
		if x > m {
			m = x
		}
	}

	return m
}

func min(a int, b int) int{
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a int, b int) int{
	if a < b {
		return b
	} else {
		return a
	}
}
