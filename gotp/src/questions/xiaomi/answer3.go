package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solution(line string) int {
    bananas := make([]int, 0)
	oneLine := strings.Split(line, " ")
	for _, s := range oneLine {
		n, _ := strconv.Atoi(s)
		bananas = append(bananas, n)
	}
	// O( nlog(n) + n )
	var ans int
	sort.Ints(bananas)
	n := len(bananas)
	for i := 0; i < n; i++ {
		total := bananas[i] * (n - i)
		if total > ans {
			ans = total
		}
	}

    return ans
}

func main() {
    r := bufio.NewReaderSize(os.Stdin, 20480)
    for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
        fmt.Println(solution(string(line)))
    }
}
