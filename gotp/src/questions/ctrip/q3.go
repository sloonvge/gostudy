package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ReplaceAll(str, "sip", "")
	answer := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] < 'A' || str[i] > 'z' {
			continue
		}
		answer = append(answer, byte(i))
	}
	fmt.Println(string(answer))
}
