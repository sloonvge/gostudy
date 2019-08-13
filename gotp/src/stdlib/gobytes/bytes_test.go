package gobytes

import (
	"bytes"
	"log"
	"testing"
)

func TestBuffer(t *testing.T) {
	t.Run("demo", func(t *testing.T) {
		var buf bytes.Buffer

		s := "test bytes buffer."
		log.Printf("write: %q\n", s)
		buf.WriteString(s)
		log.Printf("len: %d\n", buf.Len())
		log.Printf("cap: %d\n", buf.Cap())

		p := make([]byte, 8)
		n, _ := buf.Read(p)
		log.Printf("read bytes: %d\n", n)
		log.Printf("len: %d\n", buf.Len())
		log.Printf("cap: %d\n", buf.Cap())

		err := buf.UnreadByte()
		log.Printf("unread 1 byte\n")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("len: %d\n", buf.Len())
		log.Printf("cap: %d\n", buf.Cap())

	})

	t.Run("bytes", func(t *testing.T) {
		s := "ab"
		buf := bytes.NewBufferString(s)
		log.Printf("contant: %q cap: %d\n", buf.String(), buf.Cap())
		unreadBytes := buf.Bytes()
		log.Printf("unread bytes: %v\n", unreadBytes)
		// change
		buf.WriteString("cdefg")
		log.Printf("contant: %q cap: %d\n", buf.String(), buf.Cap())
		unreadBytes = unreadBytes[:cap(unreadBytes)]
		log.Printf("unread bytes: %v\n", unreadBytes)
		//
		unreadBytes[len(unreadBytes)-2] = 'X'
		log.Printf("unread bytes: %v\n", unreadBytes)
	})

	t.Run("stringtobyte", func(t *testing.T) {
		b := make([]byte, 8)
		s := "ab"
		copy(b, s)
		log.Printf("%v\n", b)
	})
}
