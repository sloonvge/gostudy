package offer

import "testing"

// input: "We are happy."
// output: "We%20are%20happy."

func TestReplaceSpace(t *testing.T) {
	in := "We are happy."
	out := "We%20are%20happy."

	got := replace1(in)
	if out != got {
		t.Errorf("error, got(%s) != want %s",
			got, out)
	}
}

func replace1(s string) string {
	var ret []rune
	for _, str := range s {
		if str != ' ' {
			ret = append(ret, rune(str))
			continue
		}
		ret = append(ret, []rune("%20")...)
	}

	return string(ret)
}
