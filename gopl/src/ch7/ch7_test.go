package main

import (
	"testing"
	"fmt"
)

/** 
* Created by wanjx in 2019/4/23 22:19
**/
 
func TestByteCounter_Write(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("Hello"))
	c.Write([]byte("world"))
	t.Log(c)
}

func TestSleep(t *testing.T) {
	main()
}

func TestWordCounter_CountWord(t *testing.T) {
	name := "word.txt"
	var c WordCounter
	c.CountWord(name)

	fmt.Println(c)
}
