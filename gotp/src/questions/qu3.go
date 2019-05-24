package questions

import (
	"fmt"
	"sort"
)

func main() {
	var ans int

	m := 0
	n := 0
	fmt.Scan(&m)
	fmt.Scan(&n)
	tasks := make([]int, n)
	for i := 0; i < n; i++ {
		t := 0
		fmt.Scan(&t)
		tasks[i] = t
	}

	if m >= n {
		ans = maxInts(tasks)
	} else {
		sort.Ints(tasks)
		tmp := tasks[:m]
		start := 0
		l := tmp[start]
		ans += tmp[start]
		for i := m; i < n; i++ {
			for j := 0; j < m; j++ {
				tmp[j] -= l
			}
			tmp[start] = tasks[i]
			start++

			l = tmp[start]
			ans += l
		}
		ans += tmp[start]
	}
	fmt.Println(ans)

}

func maxInts(s []int) int {
	m := 0
	for _, x := range s {
		if x > m {
			m = x
		}
	}

	return m
}
