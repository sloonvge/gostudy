package goio

import (
	"io"
	"log"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {
	src := strings.NewReader("from src to dst.")
	dst := new(strings.Builder)
	written, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	log.Printf("Written(%d): %q\n", written, dst.String())
}

func TestLimitedReader(t *testing.T) {
	src := strings.NewReader("test limited reader.")
	rd := io.LimitReader(src, 4)
	dst := new(strings.Builder)
	wr, err := io.Copy(dst, rd)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Writen(%d): %q\n", wr, dst.String())
	rd = io.LimitReader(src, 4)
	dst = new(strings.Builder)
	wr, _ = io.Copy(dst, rd)
	log.Printf("Read: %v Writen(%d): %q\n", *src, wr, dst.String())
}
