package main

import "fmt"

//
// import "fmt"
//
// func main() {
// 	var n, m, r, x, a, b, c int
// 	fmt.Scan(&n)
// 	fmt.Scan(&m)
// 	fmt.Scan(&r)
// 	rs := make([]int, r)
// 	for i := 0; i < r; i++ {
// 		fmt.Scan(&x)
// 		rs[i] = x
// 	}
// 	adj := make(map[int]map[int]int, 0)
// 	for {
// 		v, _ := fmt.Scanf("%d %d %d", &a, &b, &c)
// 		if v == 0 {
// 			break
// 		} else {
// 			if _, ok := adj[a]; !ok {
// 				adj[a] = make(map[int]int, 0)
// 			}
// 			adj[a][b] = c
// 		}
// 	}
//
// }
// func shortestPath(adj map[int]map[int]int, n, a, b int) int {
// 	vis := make([]bool, n)
//
// }
// func bfs(adj map[int]map[int]int, s, e int, vis []bool) {
// 	var min int
// 	queue := make([]int, 0)
// 	vis[s] = true
// 	queue = append(queue, s)
// 	for len(queue) != 0 {
// 		n := queue[0]
// 		queue = queue[1:]
// 		for k, v := range adj[n] {
// 			if !vis[k] {
// 				vis[k] = true
// 				queue = append(queue, k)
// 			}
// 		}
// 	}
// }

func main() {
	t := 10
	m := 36
	s := 255

	for tx := 1; tx <= t; tx++ {
		for ty := 1; ty <= t; ty++ {
			ta := int((4 * tx + m) / 10)
			p := ta * 50 + 13 * ty
			if p >= s {
				to := tx + ty + ta
				if to <= 10 {
					fmt.Println(to, p)
					return
				}
			}
		}
	}
}