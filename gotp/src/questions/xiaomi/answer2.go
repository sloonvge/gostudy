package main

import (
    "bufio"
    "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Time: O(nlogn + n) 	真实的实力
func solution(line string) int {
	numbers := make([]int, 0)
	oneLine := strings.Split(line, ",")
	for _, s := range oneLine {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	if len(numbers) < 2 {
		return len(numbers)
	}
	sort.Ints(numbers)
	var ans int
	n := 0
	for i := 0; i < len(numbers); i++ {
		cur := numbers[i]
		if i > 0 {
			diff := cur - numbers[i - 1]
			if diff == 1 {
				n++
			} else {
				if n + 1 > ans {
					ans = n + 1
				}
				n = 0
			}
		}
	}
	if n != 0 {
		ans = n + 1
	}
	return ans
}

// 解答连接：https://blog.csdn.net/qq_20801369/article/details/60469833
// Time: O(n)	虚假的实力
func solution2(line string) int {
	numbers := make([]int, 0)
	oneLine := strings.Split(line, ",")
	for _, s := range oneLine {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	var ans int

	set := make(map[int]struct{}, 0)
	for _, n := range numbers {
		set[n] = struct{}{}
	}
	for _, n := range numbers {
		if _, ok := set[n-1]; ok {
			continue
		}
		temp := n
		for {
			if _, ok := set[temp]; !ok {
				break
			}
			delete(set, temp)
			temp++
		}
		diff := temp - n
		if diff > ans {
			ans = diff
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
