package main
//
// import "fmt"
//
// func main() {
// 	var n, x int
// 	fmt.Scan(&n)
// 	if n == 0 {
// 		fmt.Println(0)
// 		return
// 	}
// 	numbers1 := make([]int, n)
// 	numbers2 := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&x)
// 		numbers1[i] = x
// 	}
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&x)
// 		numbers2[i] = x
// 	}
// 	answer := numbers1[0] + numbers2[0]
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if i != 0 || j != 0 {
// 				answer ^= numbers1[i] + numbers2[i]
// 			}
//
// 		}
// 	}
//
// 	fmt.Println(answer)
// }
//
// func helper(nums1, nums2 []int) {
//
// }
