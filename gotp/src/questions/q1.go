package questions

import (
	"fmt"
	"sort"
)

func main() {
	//n := 2
	in := []string{"abc", "123456789"}

	out := make([]string, 0)
	for _, s1 := range in {
		for i := 0; i < len(s1); i += 8 {
			end := mmax(i+8, len(s1))
			if len(s1[i:end]) != 8 {
				out = append(out,  add0(s1[i:end], 8 - end + i))
			} else {
				out = append(out, s1[i:end])
			}

		}
	}
	sort.Strings(out)
	fmt.Printf("%v\n", out)
}

func mmax(a int, b int) int{
	if a < b {
		return a
	}
	return b
}

func add0(s string, n int) string{
	for i := 0; i < n; i++ {
		s += "0"
	}
	return s
}