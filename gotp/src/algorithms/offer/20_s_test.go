package offer

import (
	"reflect"
	"regexp"
	"testing"
)

// 19
func Match(pattern string, b string) bool{
	if pattern == "" || b == "" {
		return false
	}
	return matchCore(pattern, b, 0, 0)
}
func matchCore(pattern string, b string, i, j int) bool {
	if i == len(pattern) - 1 || j == len(b) - 1 {
		return true
	}
	if i == len(pattern) - 1 && j != len(b) - 1 {
		return false
	}
	if pattern[i + 1] == '*' {
		if pattern[i] == b[j] || (pattern[i] == '.' && j != len(b) - 1) {
			return matchCore(pattern, b, i + 2, j + 1) ||
				matchCore(pattern, b, i, j + 1) ||
				matchCore(pattern, b, i + 2, j)
		} else {
			return matchCore(pattern, b, i + 2, j)
		}
	}
	if pattern[i] == b[j] || (pattern[i] == '.' && j != len(b) - 1) {
		return matchCore(pattern, b, i + 1, j + 1)
	}

	return false
}

func TestMatch(t *testing.T) {
	inputs := []string{
		"aaa",

	}
	helper := [][]string{
		{"a.a", "ab*ac*a", "aa.a", "ab*a"},
	}
	wants := [][]bool{
		{true, true, true, false},
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got := Match(helper[i][j],
					inputs[i]); got != wants[i][j]{
					t.Fatalf("fail %d %d, want:%v,got:%v\n", i, j, wants[i][j], got)
				}
			}
		}
	})
	t.Run("2", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			for j := 0; j < len(wants[i]); j++ {
				if got, _ := regexp.MatchString(helper[i][j],
					inputs[i]); got != wants[i][j]{
					t.Fatalf("fail %d %d, want:%v,got:%v\n", i, j, wants[i][j], got)
				}
			}
		}
	})
}

// 20
func IsNumeric(s string) bool {

	return true
}

func TestIsNumeric(t *testing.T) {
	inputs := []string{

	}
	wants := []bool{
	}
	t.Run("1", func(t *testing.T) {
		for i := 0; i < len(inputs); i++ {
			if got := IsNumeric(inputs[i]); !reflect.DeepEqual(got, wants[i]){
				t.Fatalf("fail %d, want:%v,got:%v\n", i, wants[i], got)
			}
		}
	})
}