package main
//
// import (
// 	"fmt"
// )
//
// type xhs [][2]int
//
// func NewXhs(n int) xhs {
// 	return make([][2]int, n)
// }
//
// func (x xhs) Len() int { return len(x) }
// func (x xhs) Swap(i, j int) {x[i], x[j] = x[j], x[i]}
// func (x xhs) Less(i, j int) bool {
// 	if x[i][0] < x[j][0]{
// 		return true
// 	}
// 	return x[i][1] < x[j][1]
// }
//
// func main() {
// 	var N, x, h int
// 	fmt.Scan(&N)
// 	hxs := NewXhs(N)
// 	for i := 0; i < N; i++ {
// 		fmt.Scanf("%d %d", &x, &h)
// 		hxs[i] = [2]int{x, h}
// 	}
// 	fmt.Println(maxNumbers(hxs))
// }
//
// func maxNumbers(hxs xhs) int {
// 	var ans = 1
// 	x0 := 0
// 	h0 := 0
// 	for _, hx := range hxs {
// 		x, h := hx[0], hx[1]
// 		if x > x0 && h > h0 {
// 			x0 = x
// 			h0 = h
// 			ans++
// 		}
// 	}
// 	return ans
// }
