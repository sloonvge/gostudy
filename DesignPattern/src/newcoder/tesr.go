package main

import (
	"fmt"
	"strconv"
)

/**
* Created by wanjx in 2019/3/10 23:00
**/

func main(){
	// for i:=1; i < math.MaxUint64; i++ {
	// 	if check2N(-i) {
	// 		fmt.Println(-i)
	// 		break
	// 	}
	// }
	// fmt.Println(check2N(-16))
	fmt.Println(num2bit(4))
}
//
func num2bit(s int) int{
	var res string;
	var tmp int;
	for i := s; i >= 1 ; i /= 2 {
		if i % 2 == 0 {
			tmp = 0;
		} else {
			tmp = 1;
		}
		res = strconv.Itoa(tmp) + res
	}
	ret, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(ret)
}

func check2N(s int) bool{
	return (s & (s - 1)) == 0
}