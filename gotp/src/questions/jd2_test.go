package questions

func pop(mat [][]int, rows, cols int, s string) bool {
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
func hasPath(mat [][]int, rows, cols, row, col int,
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
