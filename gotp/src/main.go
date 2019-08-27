package main

import (
	"fmt"
)

func q1(values []int, ops []string) {
	for i := len(values) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			cur := values[j]
			next := values[j + 1]
			if next < cur && (ops[j + 1] == "+" || ops[j + 1] == "-") {
				values[j], values[j + 1] = values[j + 1], values[j]
			}
		}
	}

}

func main() {
	// n := 0
	// fmt.Scanln(&n)
	// values := make([]int, 0)
	// ops := make([]string, 0)
	// var s string
	// for i := 0; i < 2 * n - 1; i++ {
	// 	fmt.Scan(&s)
	// 	switch s {
	// 	case "+", "/", "*", "-":
	// 		ops = append(ops, s)
	// 	default:
	// 		i, err := strconv.Atoi(s)
	// 		if err != nil {
	// 			panic(fmt.Sprintf("convert string %s to int, err:%s", s, err))
	// 		}
	// 		values = append(values, i)
	// 	}
	// }
	values := []int{3,2,1,-4,-5,1}
	ops := []string{"+","+","+","*","+"}
	q1(values, ops)
	for i := 0; i < len(values) - 1; i++ {
		fmt.Printf("%d %s ", values[i], ops[i])
	}
	fmt.Printf("%d", values[len(values) - 1])
}