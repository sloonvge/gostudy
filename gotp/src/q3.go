package main

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
