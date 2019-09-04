package main
//
// import "fmt"
//
// func main() {
// 	var N, startX, startY, endX, endY int
// 	var s string
// 	fmt.Scan(&N)
//
// 	vis := make([][]bool, N)
// 	for i := 0; i < N; i++ {
// 		vis[i] = make([]bool, N)
// 	}
// 	for i := 0; i < N; i++ {
// 		fmt.Scanln(s)
// 		for j := range s {
// 			switch s[j] {
// 			case '#':
// 				vis[i][j] = true
// 			case 'S':
// 				startX = i
// 				startY = j
// 			case 'E':
// 				endX = i
// 				endY = j
// 			default:
// 				continue
// 			}
// 		}
// 	}
// 	fmt.Println(ShortestPath(N, startX, startY, endX, endY, vis))
// }
//
// func ShortestPath(N, sx, sy, ex, ey int, vis [][]bool) int{
// 	ans := shortestPath(N, sx, sy, vis)
// 	if vis[ex][ey] {
// 		return ans
// 	}
// 	return 0
// }
//
// func shortestPath(N int, row, col int, vis [][]bool) int{
// 	if row < 0 || row >= N ||
// 		col < 0 || col >= N ||
// 		vis[row][col] {
// 		return 0
// 	}
// 	vis[row][col] = true
//
// 	return 1 + min(shortestPath(N, row, col+1, vis),
// 		shortestPath(N, row, col-1, vis),
// 		shortestPath(N, row+1, col, vis),
// 		shortestPath(N, row-1, col, vis),
// 		)
// }
//
// func min(a, b, c ,d int) int{
// 	var n int
// 	if a < b {
// 		n = a
// 	} else {
// 		n = b
// 	}
// 	if c < n {
// 		n = c
// 	}
// 	if d < n {
// 		n = d
// 	}
// 	return n
// }
