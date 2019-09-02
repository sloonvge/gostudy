package offer

import (
	"fmt"
	"testing"
)

func GenerateAll(a []byte) {
	generateAll(a, 0)
}
func generateAll(a []byte, pos int) {
	if pos == 6 {
		fmt.Println(string(a), len(a), cap(a))
	} else {
		a[pos] = '('
		generateAll(a, pos + 1)
		a[pos] = ')'
		generateAll(a, pos + 1)
	}

}

func TestGenerateAll(t *testing.T) {
	GenerateAll(make([]byte, 6))
}

// 12
func HasPath(mat [][]string, rows, cols int, s string) bool {
	if len(mat) == 0 || rows < 1 || cols < 1 ||
		s == "" {
		return false
	}
	visited := make([][]bool, len(mat))
	for i := 0; i < len(mat); i++ {
		visited[i] = make([]bool, len(mat[0]))
	}

	l := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if hasPath(mat, rows, cols, row, col,
				s, l, visited) {
				return true
			}
		}
	}
	return false
}
func hasPath(mat [][]string, rows, cols, row, col int,
	s string, l int, visited [][]bool) bool{
	if len(s) == l {
		return true
	}
	if row < 0 || col < 0 ||
		row >= rows || col >= cols ||
		visited[row][col]{
		return false
	}

	var path bool
	if row >= 0 && row < rows &&
		col >= 0 && col < cols &&
		mat[row][col] == string(s[l]) &&
		!visited[row][col]{
		l++
	}

	visited[row][col] = true
	path = hasPath(mat, rows, cols, row, col - 1,
			s, l, visited) ||
		hasPath(mat, rows, cols, row - 1, col,
			s, l, visited) ||
		hasPath(mat, rows, cols, row, col + 1,
			s, l, visited) ||
		hasPath(mat, rows, cols, row + 1, col,
			s, l, visited)
	if !path {
		l--
		visited[row][col] = false
	}

	return path
}

func TestHasPath(t *testing.T)  {
	inputs := [][][]string{
		{
			{"a", "b", "t", "g"},
			{"c", "f", "c", "s"},
			{"j", "d", "e", "h"},
		},
		{
			{"a", "b", "t", "g"},
		},
		{
			{"a"},
			{"c"},
			{"j"},
		},
		{
			{"a", "a", "a", "a"},
			{"a", "a", "a", "a"},
			{"a", "a", "a", "a"},
		},
	}
	strs := [][]string{
		{"bfce", "abfb"},
		{"a", "z"},
		{"a", "z"},
		{"aaa", "z"},
	}
	wants := [][]bool{
		{true, false},
		{true, false},
		{true, false},
		{true, false},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			rows := len(inputs[i])
			cols := len(inputs[i][0])
			for j := range strs[i] {
				got := HasPath(inputs[i], rows, cols, strs[i][j])
				if got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}

		}
	})
}

// 13
func MovingCount(threshold int, rows int, cols int) int {
	if rows < 1 || cols < 1 || threshold < 0{
		return 0
	}
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	count := movingCount(threshold, rows, cols, 0, 0, visited)

	return count
}
func movingCount(threshold int, rows, cols, row, col int, visited [][]bool) int {
	if row < 0 || row >= rows ||
		col < 0 || col >= cols ||
		visited[row][col] {
		return 0
	}

	var count int
	if check(threshold, rows, cols, row, col, visited) {
		visited[row][col] = true
		count = 1 + movingCount(threshold, rows, cols, row, col - 1, visited) +
			movingCount(threshold, rows, cols, row, col + 1, visited) +
			movingCount(threshold, rows, cols, row - 1, col, visited) +
			movingCount(threshold, rows, cols, row + 1, col, visited)
	}

	return count
}

func check(threshold int, rows, cols, row, col int, visited [][]bool) bool {
	return (getNumberSum(row) + getNumberSum(col)) <= threshold
}
func getNumberSum(a int) (b int) {
	for {
		c := a % 10
		if c == 0 {
			break
		}
		b += c
		a = c / 10
	}
	return b
}

func TestMovingCount(t *testing.T) {
	inputs := [][2]int{
		{2, 3},
		{1, 2},
		{3, 1},
	}
	wants := []int{
		6,
		2,
		3,
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			rows := inputs[i][0]
			cols := inputs[i][1]
			// fmt.Printf("%d %d\n", rows, cols)
			got := MovingCount(18, rows, cols)
			if got != wants[i] {
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}