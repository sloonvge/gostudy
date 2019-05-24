package questions

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var ans int
	a := ""
	b := ""
	k := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&k)


	inputs := ints2strings(strings.Split(a,","))
	sales := ints2strings(strings.Split(b, ","))
	n := len(inputs)

	for i := 0; i < n; i++ {
		ans += sales[i] - inputs[i]
	}

	fmt.Println(k + ans)

}

//func ints2strings(s []string) []int {
//	ret := make([]int, 0)
//	for _, x := range s {
//		i, _ := strconv.Atoi(x)
//		ret = append(ret, i)
//	}
//
//	return ret
//}

