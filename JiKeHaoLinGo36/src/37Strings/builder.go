package main

import (
	"fmt"
	"strings"
	"strconv"
)

// Strings.Builder和string的区别
/*
E example
1. 已存在内容不可变，可以拼接更多的内容（主要优势在于“字符串拼接”上面）
2. 减少内存分配和内容拷贝的次数
3. 可将内容重置，可重用值
*/

func Union(a string, b string) string{
	as := strings.Split(a, " ")
	bs := strings.Split(b, " ")

	cs := as[:0]
	if len(as) == 2 && len(bs) == 2 {
		var str2float = func(s string, bitSize int) (float64){
			r, err := strconv.ParseFloat(s,64)
			if err != nil {
				fmt.Println(err)
			}
			return r
		}

		aq := str2float(as[0], 64)
		bq := str2float(bs[0], 64)
		if as[1] == bs[1] {
			cq := aq + bq
			c := strconv.FormatFloat(cq, 'f', 2, 64)
			cs = append(cs, c)
			cs = append(cs, as[1])
		} else {
			return strings.Join(as, " ") + "|" + strings.Join(bs, " ")
		}
	}

	return strings.Join(cs, " ")
}

type E string

func (e *E) Join(s string) {
	*e = *e + E(s)
}

func (e *E) Builder(s string) {
	var sb strings.Builder
	sb.WriteString(string(*e))
	sb.WriteString(s)
	*e = E(sb.String())
}

func main() {
	a := "2.13 Kg"
	b := "0.12 Kg"

	c := Union(a, b)

	fmt.Println(c)
}
