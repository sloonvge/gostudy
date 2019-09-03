package main

import (
	"fmt"
	"math"
)

var min = math.MaxInt64
var find = false
func main() {
	var n, m, k int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)
	ks := make([][2]int, k)
	for i := 0; i < k; i++ {
		x := 0
		y := 0
		fmt.Scanf("%d %d", &x, &y)
		ks[i] = [2]int{x, y}
	}
	fmt.Println(ShortestPath(n, m, ks))
}

func ShortestPath(n, m int, ks [][2]int) int {
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}
	for _, xy := range ks {
		x, y := xy[0], xy[1]
		vis[x][y] = true
	}
	shortestPath(n, m, 0, 0,0, vis)
	if find {
		return min
	}
	return 0
}

func shortestPath(n, m int, row, col,s int, vis [][]bool){
	if row == n - 1 && col == m - 1{
		find = true
		if s < min {
			min = s
		}
		return
	}

	if col < m - 1 && !vis[row][col+1]{
		vis[row][col+1] = true
		shortestPath(n, m, row, col+1, s + 1,vis)
		vis[row][col+1] = false
	}
	if row < n - 1 && !vis[row+1][col]{
		vis[row+1][col] = true
		shortestPath(n, m, row+1, col, s + 1,vis)
		vis[row+1][col] = false
	}
	if col > 0 && !vis[row][col-1]{
		vis[row][col-1] = true
		shortestPath(n, m, row, col-1, s + 1,vis)
		vis[row][col-1] = false
	}
	if row > 0 && !vis[row-1][col]{
		vis[row-1][col] = true
		shortestPath(n, m, row-1, col, s + 1,vis)
		vis[row-1][col] = false
	}

	return
}
