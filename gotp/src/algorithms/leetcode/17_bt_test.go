package leetcode

import (
	"fmt"
	"testing"
)

// 17
func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	if digits == ""{
		return ans
	}

	dict := make(map[byte]string, 0)
	dict['2'] = "abc"
	dict['3'] = "def"
	dict['4'] = "ghi"
	dict['5'] = "jkl"
	dict['6'] = "mno"
	dict['7'] = "pqrs"
	dict['8'] = "tuv"
	dict['9'] = "wxyz"
	bt(ans, []byte(digits), dict, make([]byte, 0), 0, len(digits) - 1)
	return make([]string, 0)
}

func bt(ans []string, digits []byte, dict map[byte]string, b []byte, s, e int) {
	if len(b) == len(digits) {
		fmt.Println(string(b))
		ans  = append(ans, string(b))
		return
	}
	for i := s; i <= e; i++ {
		cur := digits[i]
		for j := 0; j < len(dict[cur]); j++ {
			b = append(b, dict[cur][j])
			bt(ans, digits, dict, b, i + 1, e)
			b = b[:len(b) - 1]
		}
	}

}

func TestLetterCombinations(t *testing.T)  {
	letterCombinations("23")
}

// 79
func exist(board [][]byte, word string) bool {
	var ans bool
	rows := len(board)
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			find := dfs(&board, word, 0, rows - 1, cols - 1, i, j)
			if find {
				ans = true
				return ans
			}
		}
	}

	return ans
}

func dfs(board *[][]byte, word string, index, rows, cols, row, col int) bool {
	if index == len(word) {
		return true
	}
	if row < 0 || row > rows ||
		col < 0 || col > cols ||
		(*board)[row][col] != word[index]{
		return false
	}
	(*board)[row][col] ^= byte(255)
	find := dfs(board, word, index + 1, rows, cols, row, col + 1) ||
		dfs(board, word, index + 1, rows, cols, row, col - 1) ||
		dfs(board, word, index + 1, rows, cols, row + 1, col) ||
		dfs(board, word, index + 1, rows, cols, row - 1, col)
	(*board)[row][col] ^= byte(255)
	return find
}

func TestExist(t *testing.T)  {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exist(board, "ASADFB")
}

// 62
func uniquePaths(m int, n int) int {
	nPath := 0
	backtrace(m - 1, n - 1, 0, 0, &nPath)
	fmt.Println(nPath)
	return nPath
}

func backtrace(rows, cols, row, col int, nPath *int) {
	if row == rows && col == cols {
		*nPath++
		return
	}
	if row < 0 || row > rows ||
		col < 0 || col > cols {
		return
	}
	backtrace(rows, cols, row + 1, col, nPath)
	backtrace(rows, cols, row, col + 1, nPath)
}

func TestUniquePaths(t *testing.T)  {
	uniquePaths(51, 9)
}