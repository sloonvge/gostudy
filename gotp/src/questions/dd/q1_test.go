package dd

import (
	"fmt"
	"testing"
)

func q1(values []int, ops []string) {
	for i := len(values) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			cur := values[j]
			next := values[j + 1]
			if 0 < j && j < i - 1 {
				if next < cur &&
					(ops[j - 1] == "+" || ops[j - 1] == "-") &&
					(ops[j + 1] == "+" || ops[j + 1] == "-") {
					values[j], values[j + 1] = values[j + 1], values[j]
				}
			} else if j > i - 1{
				if next < cur && j > 0 && (ops[j - 1] == "+" || ops[j - 1] == "-") &&
					(ops[j] == "+" || ops[j] == "-"){
					values[j], values[j + 1] = values[j + 1], values[j]
				}
			} else if j == 0 {
				if next < cur &&
					(ops[j + 1] == "+" || ops[j + 1] == "-"){
					values[j], values[j + 1] = values[j + 1], values[j]
				}
			}

		}
	}

}

func TestQ1(t *testing.T) {
	// n := 0
	// fmt.Scan(&n)
	// values := make([]int, 0)
	// ops := make([]string, 0)
	// var s string
	// for {
	// 	_, err := fmt.Scan(&s)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 	} else {
	// 		switch s {
	// 		case "+", "/", "*", "-":
	// 			ops = append(ops, s)
	// 		default:
	// 			i, err := strconv.Atoi(s)
	// 			if err != nil {
	// 				panic(fmt.Sprintf("convert string %s to int, err:%s", s, err))
	// 			}
	// 			values = append(values, i)
	// 		}
	// 	}
	// }
	values := []int{3,2,1,-4,-5,-7}
	ops := []string{"+","+","+","+","*"}
	q1(values, ops)
	fmt.Printf("%v\n", values)
}