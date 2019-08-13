package gounicode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("range", func(t *testing.T) {
		s := "Go爱好者"
		for i, v := range s {
			fmt.Printf("%d: %q %s %s\n", i, v, string(v), string(byte(v)))
		}
	})

	t.Run("range", func(t *testing.T) {
		str := "Go 爱好者 "
		fmt.Printf("The string: %q\n", str)
		fmt.Printf("  => runes(char): %q\n", []rune(str))
		fmt.Printf("  => runes(hex): %x\n", []rune(str))
		fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
	})

	t.Run("utf-8", func(t *testing.T) {
		fmt.Printf("a: %s\n", string([]byte{97, 98}))
		fmt.Printf("len('a') = %d, len('我') = %d\n", len("a"), len("我"))
	})
}
