package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solution(line string) string {
	var ans = "true"
	oneLine := strings.Split(line, " ")
	short, long := oneLine[0], oneLine[1]
	charMap := make(map[rune]int, 0)
	for _, s := range []rune(long) {
		charMap[s]++
	}
	for _, s := range []rune(short) {
		count, ok := charMap[s]
		if !ok {
			ans = "false"
			return ans
		}
		count--
		if count < 0 {
			ans = "false"
			return ans
		}
		charMap[s] = count

	}

	return ans
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
