package main

// <<<<<<< Updated upstream
import (
	"fmt"
)

func main() {
	var ans int
	a := 0

	vals := make([]int, 0)
	for {
		n, _ := fmt.Scanf(",",&a)
		if n == 0 {
			break
		} else {
			vals = append(vals, a)
		}
	}
	n := len(vals)
	inputs := vals[:(n-1)/2]
	sales := vals[(n-1)/2:]
	k := vals[n]

	for i := 0; i < (n - 1) / 2; i++ {
		ans += sales[i] - inputs[i]
	}

	fmt.Println(k + ans)

}
// =======
// import "fmt"
//
// /**
// * Created by wanjx in 2019/4/14 11:16
// **/
//
// func main() {
// 	n := 0
// 	fmt.Scan(&n)
//
// 	vals := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		x := 0
// 		fmt.Scan(&x)
//
// 		vals[i] = x
// 	}
//
// 	fmt.Println(q3(vals))
// }
//
//
// func q3(values []int) int{
// 	ans := 0
// 	values[0] = 0
// 	for i := 0; i < len(values) - 1; i++ {
// 		ans += - values[i + 1] + values[i]
// 	}
//
// 	return  -ans
// }
// >>>>>>> Stashed changes
