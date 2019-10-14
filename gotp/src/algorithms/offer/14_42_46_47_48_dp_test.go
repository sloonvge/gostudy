package offer

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// 14

func MaxProductAfterCutting1(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	a := maxProductAfterCutting1(n)
	return a
}
func maxProductAfterCutting1(n int) int {
	if n <= 3 {
		return n
	}
	a := 0
	for i := 1; i < n; i++ {
		a = intMax(a, maxProductAfterCutting1(i) * maxProductAfterCutting1(n - i))
	}
	return a
}

func MaxProductAfterCutting2(a int) (b int){
	if a < 2 {
		return 0
	}
	if a == 2 {
		return 1
	}
	if a == 3 {
		return 2
	}
	products := make([]int, a + 1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	var max int
	for i := 4; i <= a; i++ {
		max = 0
		for j := 1; j <= i / 2; j++ {
			max = intMax(max, products[j] * products[i - j])
		}
		products[i] = max
	}
	return products[a]

}

func MaxProductAfterCutting3(a int) (b int) {
	if a < 2 {
		return 0
	}
	if a == 2 {
		return 1
	}
	if a == 3 {
		return 2
	}

	timesOf3 := a / 3
	if a - timesOf3 *  3 == 1 {
		timesOf3 -= 1
	}
	timesOf2 := (a - timesOf3 * 3) / 2
	return int(math.Pow(3, float64(timesOf3)) * math.Pow(2, float64(timesOf2)))
}

func TestMaxProduct(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4, 8}
	wants := []int{0, 0, 1, 2, 4, 18}
	t.Run("1", func(t *testing.T) {
		for i, input := range inputs {
			got := MaxProductAfterCutting1(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i, input := range inputs {
			got := MaxProductAfterCutting2(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
	t.Run("3", func(t *testing.T) {
		for i, input := range inputs {
			got := MaxProductAfterCutting3(input)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}

func intMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 42
func FindGreatestSumOfSubArray(data []int, n int) {
	maxSum := 0
	cur := 0
	for i := 0; i < n; i++ {
		if cur <= 0 {
			cur = data[i]
		} else {
			cur += data[i]
		}
		if cur > maxSum {
			maxSum = cur
		}
	}
	fmt.Println(maxSum)
}

func TestFindGreatestSumOfSubArray(t *testing.T) {
	FindGreatestSumOfSubArray([]int{1,-2,3,10,-4,7,2,-5}, 8)
}

// 46
func TransNumbersBT(s string) {
	t := make([]byte, 0)
	transNumbersBT(s, t, 0, 0)
}
func transNumbersBT(s string, t []byte, start, end int) {
	if end > len(s) {
		if start == len(s) && end == len(s) + 1 {
			t = t[:len(t) - 1]
			fmt.Printf("%d %d %q\n", start, end, string(t))
		}
		return
	}
	if start != end {
		t = append(t, s[start:end]...)
		t = append(t, ' ')
	}
	transNumbersBT(s, t, end, end + 1)
	transNumbersBT(s, t, end, end + 2)
}

func TranslationNumbers(s string) int {
	return translationNumbers(s, 0)
}
func translationNumbers(s string, i int) int {
	if i + 2 > len(s) {
		return 1
	}
	return translationNumbers(s, i + 1) + f(s, i) * translationNumbers(s, i + 2)
}

func f(s string, i int) int{
	n, _ := strconv.Atoi(string(s[i:i+2]))
	if n <= 26 && n >= 10 {
		return 1
	}
	return 0
}

func TransNumbers(s string) int{
	return transNumbers(s)
}
func transNumbers(s string) int{
	n := len(s)
	counts := make([]int, n)
	count := 0
	for i := n - 1; i >= 0; i-- {
		count = 0
		if i < n - 1 {
			count = counts[i + 1]
		} else {
			count = 1
		}
		if i < n - 1 {
			num1 := s[i] - '0'
			num2 := s[i + 1] - '0'
			num := num1 * 10 + num2
			if num >= 10 && num <= 25 {
				if i < n - 2 {
					count += counts[i + 2]
				} else {
					count += 1
				}
			}
		}
		counts[i] = count
	}
	count = counts[0]

	return count
}

func TestTranslationNumbers(t *testing.T) {
	//
	t.Run("0", func(t *testing.T) {
		TransNumbersBT("12258")
	})
	t.Run("1", func(t *testing.T) {
		fmt.Println(TranslationNumbers("12258"))
	})
	t.Run("2", func(t *testing.T) {
		fmt.Println(TransNumbers("12258"))
	})
}

// 47
func GetMaxValue(a [][]int) int {
	rows := len(a)
	cols := len(a[0])
	return getMaxValue(a, rows - 1, cols - 1, 3, 3)
}
func getMaxValue(a [][]int, rows, cols, row, col int) int {
	if row < 0 || row > rows ||
		col < 0 || col > cols {
		return 0
	}
	return int(math.Max(float64(getMaxValue(a, rows, cols, row, col-1)),
		float64(getMaxValue(a, rows, cols, row-1, col)))) + a[row][col]
}

func GetMaxValue2(a [][]int) int {
	rows := len(a)
	cols := len(a[0])
	return getMaxValue2(a, rows, cols)
}
func getMaxValue2(a [][]int, rows, cols int) int {
	if rows <= 0 || cols <= 0 {
		return 0
	}
	maxValues := make([][]int, rows)
	for i := 0; i < rows; i++ {
		maxValues[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			left := 0
			down := 0
			if i > 0 {
				down = maxValues[i-1][j]
			}
			if j > 0 {
				left = maxValues[i][j-1]
			}
			maxValues[i][j] = int(math.Max(float64(left), float64(down))) + a[i][j]
		}
	}

	return maxValues[rows-1][cols-1]
}

func TestGetMaxValue(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 10, 3, 8},
			{12, 2, 9, 6},
			{5, 7, 4, 11},
			{3, 7, 16, 5},
		},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			input := inputs[i]
			fmt.Println(GetMaxValue(input))
		}
	})
	t.Run("2", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			input := inputs[i]
			fmt.Println(GetMaxValue2(input))
		}
	})
}

// 48