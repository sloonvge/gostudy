package main

import (
	"fmt"
	"sort"
)

/** 
* Created by wanjx in 2019/4/14 10:13
**/
 
func main() {
	n := 0
	fmt.Scan(&n)
	exs := make([][]int, n)

	for l := 0; l < n; l++ {
		ln := 0
		fmt.Scan(&ln)

		lns := make([]int, ln)
		for k := 0; k < ln; k ++ {
			a := 0
			fmt.Scan(&a)

			lns[k] = a
		}
		exs[l] = lns
	}

	fmt.Printf("%v\n", exs)
	for _, x := range exs {
		t := q5(x)
		fmt.Println(t)
	}
}

func q5(exs []int) int {
	var ans int

	n := len(exs)
	sort.Ints(exs)

	i := 0
	j := n - 1
	if n % 3 == 0 {
		for i < j {
			ans += exs[j]
			j -= 2
			i += 1
		}
	}

	if n % 3 == 1 {
		k := n / 3
		for t := 0; t < k - 1 ; t++ {
			ans += exs[j - 1]
			j -= 2
			i += 1
		}
		ans = ans + exs[i + 1] + exs[j - 1]
		tmp := ans

		ans = 0
		i = 0
		j = n
		ans = ans + exs[i + 1] + exs[j - 1]
		i += 1
		j -= 2
		for t := 0; t < k - 1 ; t++ {
			ans += exs[j - 1]
			j -= 2
			i += 1
		}

		ans = minc(tmp, ans)
	}

	if n % 3 == 2 {
		k := n / 3
		if k == 0 {
			ans += exs[i - 1]
		} else {
			for t := 0; t < k - 1; t++ {
				ans += exs[j - 1]
				j -= 2
				i += 1
			}
			ans = ans + exs[i + 1] + exs[j - 1]
		}

	}

	return ans
}

func minc(a int, b int) int {
	if a < b {
		return a
	}
	return b
}