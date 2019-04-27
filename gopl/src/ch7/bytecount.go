package main

/** 
* Created by wanjx in 2019/4/23 22:18
**/
 
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}