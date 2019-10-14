package offer

import (
	"fmt"
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

// 58
// func ReverseSentence(data string) string{
// 	if data == "" {
// 		return data
// 	}
//
// 	bytes := []byte(data)
// 	var begin, end int
// 	begin = 0
// 	end = len(bytes) - 1
// 	reverse(bytes, begin, end)
// 	begin = 0
// 	for i, b := range bytes {
// 		if b != ' ' {
// 			continue
// 		}
// 		if i > 1{
// 			end = i - 1
// 			reverse(bytes, begin, end)
// 		}
// 		begin = i + 1
// 	}
//
// 	return string(bytes)
// }
func ReverseSentence(s string) string {
	if s == "" {
		return ""
	}
	var ans []byte
	ans = []byte(s)
	reverse(ans, 0, len(s) - 1)
	var begin, end int
	for begin < len(s) {
		if ans[begin] == ' ' {
			begin++
			end++
		} else if end == len(s) || ans[end] == ' ' {
			reverse(ans, begin, end - 1)
			end++
			begin = end
		} else {
			end++
		}
	}

	return string(ans)
}
func reverse(b []byte, i, j int) {
	if i < 0 || j >= len(b) {
		return
	}
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

func TestReverseSentence(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		fmt.Printf("%q\n", ReverseSentence("I am a student."))
		fmt.Printf("%q\n", ReverseSentence(" I am a student."))
		fmt.Printf("%q\n", ReverseSentence("I am a student. "))
		fmt.Printf("%q\n", ReverseSentence("    "))
	})
}

func LeftRotateString(s string, n int) string {
	if s == "" {
		return s
	}
	bytes := []byte(s)
	l := len(bytes)
	if n > 0 && l > 0 && n < l{
		reverse(bytes, 0, n - 1)
		reverse(bytes, n, l - 1)
		reverse(bytes, 0, l - 1)
	}

	return string(bytes)
}

func TestLeftRotateString(t *testing.T) {

	t.Run("1", func(t *testing.T) {
		fmt.Printf("%q\n", LeftRotateString("world", 0))
		fmt.Printf("%q\n", LeftRotateString("world", 1))
		fmt.Printf("%q\n", LeftRotateString("world", 4))
		fmt.Printf("%q\n", LeftRotateString("world", 6))
	})
}