package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/** 
* Created by wanjx in 2019/6/22 16:32
**/

func multiply() {
	for i := 1;i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func isPerfectNumber(n int) bool{
	factorSum := 0
	for i := 1; i < n; i++ {
		if n % i == 0 {
			factorSum += i
		}
	}

	if factorSum == n {
		return true
	}
	return false
}

func perfectNumber() {
	for i := 1; i < 1000; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()
}

func isPalindrome(s string) bool {
	var result = true
	runeS := []rune(s)
	for i := 0; i < len(runeS)/2; i++ {
		if runeS[i] != runeS[len(runeS)-i-1] {
			result = false
		}
	}

	return result
}

func palindrome() {
	var s string
	for {
		fmt.Scanf("%s\n", &s)
		if isPalindrome(s) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}

}

func charCount() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read error")
	}
	// fmt.Scanf("%s\n", &input)
	input = string(result)
	var spaceCount, numberCount, letterVCount, otherCount int
	for _, s := range input {
		// switch {
		// case s == 32:
		// 	spaceCount++
		// case 48 <= s && s < 58:
		// 	numberCount++
		// case 65 <= s && s < 91 || 97 <= s && s < 123:
		// 	letterVCount++
		// default:
		// 	otherCount++
		// }
		switch  {
		case s >= 'a' && s <= 'z':
			fallthrough
		case s >= 'A' && s <= 'z':
			letterVCount++
		case s == ' ':
			spaceCount++
		case s >= '0' && s <= '9':
			numberCount++
		default:
			otherCount++
		}

	}
	fmt.Printf("space count: %d\nnumber count: %d\n" +
		"letter count: %d\n other count: %d\n",
		spaceCount, numberCount, letterVCount, otherCount)
}

// todo -1 + 1
func addLargeNumber(a string, b string) string {
	var result string
	var j uint8
	if len(b) > len(a) {
		a ,b = b, a
	}
	for i := 0; i < len(b); i++ {
		va := a[len(a)-i-1] - '0'
		vb := b[len(b)-i-1] - '0'
		v := va + vb + j
		if v > 10 {
			j = v % 10
		}
		result = fmt.Sprintf("%d%s", v, result)
	}
	for i := 0; i < len(a) - len(b); i++ {
		result = fmt.Sprintf("%d%s", a[i] - '0', result)
	}
	fmt.Println(result)
	return result
}

func largeNumber() {
	var s string
	fmt.Scanf("%s\n", &s)
	strSlice := strings.Split(s, "+")
	if len(strSlice) != 2{
		fmt.Println("input a + b")
	}

	str1 := strings.TrimSpace(strSlice[0])
	str2 :=  strings.TrimSpace(strSlice[1])
	addLargeNumber(str1, str2)
}

func main() {
	// multiply()
	// perfectNumber()
	// palindrome()
	// charCount()
	largeNumber()
}
