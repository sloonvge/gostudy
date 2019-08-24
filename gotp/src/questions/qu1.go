package questions

import (
	"strconv"
	"strings"
)

// func main() {
// 	input := ""
// 	fmt.Scan(&input)
//
// 	array := generateArray(input)
// 	sort.Ints(array)
//
// 	if !valid(array) {
// 		fmt.Println("invalid")
// 	} else {
// 		var ans string
// 		for i, x := range array {
// 			if x < 3 {
// 				ans += strconv.Itoa(x)
// 				array = array[1:]
// 			} else if x < 6 {
// 				ans += strconv.Itoa(x)
// 				array = array[1:]
// 			} else {
// 				ans += strconv.Itoa(x)
// 				array = array[1:]
// 			}
// 			if i == 1 || i == 3{
// 				ans += ":"
// 			}
// 		}
//
// 		fmt.Println(ans)
// 	}
// }

func generateArray(s string) []int {
	s = strings.Trim(s, "[")
	s = strings.Trim(s, "]")
	return ints2strings(strings.Split(s, ","))
}

func ints2strings(s []string) []int {
	ret := make([]int, 0)
	for _, x := range s {
		i, _ := strconv.Atoi(x)
		ret = append(ret, i)
	}

	return ret
}

func valid(s []int) bool {
	for i := 6; i < 10; i++ {
		t := 0
		for _, x := range s {
			if x == i {
				t++
			}
		}
		if t > 2 {
			return false
		}
	}

	return true
}