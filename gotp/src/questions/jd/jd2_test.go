package jd

import "testing"

func jd2(a [][]int) int {
	for {
		res := 0
		label := 1
		maxLabel := 0
		labels := make([][]int, 5)
		for i := 0; i < 5; i++ {
			labels[i] = make([]int, 5)
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if labels[i][j] == 0 && a[i][j] != -1 {
					n := dfs(a, i, j, a[i][j], label, labels)
					if n > res {
						res = n
						maxLabel = label
					}
					label += 1
				}
			}
		}

		if res >= 3 {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if labels[i][j] == maxLabel {
						a[i][j] = -1
					}
				}
			}

			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					k := i - 1
					for a[i][j] == -1 && k >= 0 {
						a[i][j], a[k][j] = a[k][j], a[i][j]
						k--
					}
				}
			}
		}
		if res < 3 {
			break
		}
	}

	ans := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if a[i][j] != -1 {
				ans++
			}
		}
	}
	return ans
}
func dfs(a [][]int, i, j, num int, label int, labels [][]int) int {
	if i < 0 || i >=5 ||
		j < 0 || j >= 5 ||
		labels[i][j] != 0 || a[i][j] != num {
		return 0
	}

	labels[i][j] = label

	return 1 + dfs(a, i - 1, j, num, label, labels) +
		dfs(a, i + 1, j, num, label, labels) +
		dfs(a, i, j - 1, num, label, labels) +
		dfs(a, i, j + 1, num, label, labels)
}


func TestJd2(t *testing.T) {
	inputs := [][][]int{
		{
			{3, 1, 2, 1, 1},
			{1, 1, 1, 1, 3},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{3, 1, 2, 2, 2},
		},
	}
	wants := [][]int{
		{3},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := jd2(inputs[i]); got != wants[i][j] {
					t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i][j], got)
				}
			}
		}
	})
}


