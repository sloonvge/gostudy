package main

import "testing"

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