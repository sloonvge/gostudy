package dd

import "fmt"

func q2(a1 []int, a2 []int, total, cost int) {

}

func t() {
	var n, total, cost, v int
	fmt.Scan(&n)
	fmt.Scan(&total)
	fmt.Scan(cost)

	a1 := make([]int, n)
	a2 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		a1[i] = v
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		a2[i] = v
	}
}