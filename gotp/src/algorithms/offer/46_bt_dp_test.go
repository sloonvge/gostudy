package offer

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

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

// return 1 + shortestPath(n, m, row, col+1, vis) +
// shortestPath(n, m, row, col-1, vis) +
// shortestPath(n, m, row+1, col, vis) +
// shortestPath(n, m, row-1, col, vis)