package questions

import "fmt"

/** 
* Created by wanjx in 2019/4/5 20:59
**/

var coins = []int{1, 2, 5, 10}
var dp1 = make(map[int]int, 0)
func main()  {

	fmt.Println(f1(20))
	fmt.Printf("%v\n", dp1)
	fmt.Println(max1(dp1))
}

func f1(sum int) int{

	if _, ok := dp1[sum]; ok {
		return dp1[sum]
	}
	ans := 10000000
	if sum <= 0 {
		return 0
	}
	for _, x := range coins {
		if sum - x >= 0 {
			ans = min1(ans, 1 + f1(sum - x))
		}
	}
	dp1[sum] = ans
	return ans
}

func min1(a int, b int) int{
	if a > b {
		return b
	} else {
		return a
	}
}

func max1(s map[int]int) int {
	m := 0
	for _, v := range s {
		if v > m {
			m = v
		}
	}

	return m
}