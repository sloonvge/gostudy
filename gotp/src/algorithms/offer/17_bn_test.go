package offer

import (
	"fmt"
	"math"
	"testing"
)

func Print1ToN(n int) {
	x := int(math.Pow(10.0, float64(n)))
	for i := 0; i < x; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func Print1ToN2(n int) {
	if n <= 0 {
		return
	}
	number := make([]byte, n)
	for i := 0; i < len(number); i++ {
		number[i] = '0'
	}
	for !increment(&number) {
		printNumber(&number)
	}
}
func increment(a *[]byte) (isOverflow bool) {
	n := len(*a)
	var takeOver uint8
	for i := n - 1; i >= 0; i-- {
		iSum := (*a)[i] - '0' + takeOver
		if i == n -1 {
			iSum++
		}
		if iSum >= 10 {
			if i == 0 {
				isOverflow = true
			} else {
				iSum -= 10
				takeOver = 1
				(*a)[i] = '0' + iSum
			}
		} else {
			(*a)[i] = '0' + iSum
			break
		}
	}
	return
}
func printNumber(a *[]byte) {
	var isBegin0 = true
	for i := 0; i < len(*a); i++ {
		if isBegin0 && (*a)[i] != '0' {
			isBegin0 = false
		}
		if !isBegin0 {
			fmt.Printf("%s", string((*a)[i]))
		}
	}
	fmt.Printf(" ")
}

func TestPrint1ToN(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		Print1ToN(3)
	})
	t.Run("2", func(t *testing.T) {
		Print1ToN2(3)
	})
}
