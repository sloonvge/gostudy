package main

import (
	"fmt"
)

func badCall() {
	panic("bad call")
}

func testRecover() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing %s\r\n", e)
		}
	}()

	badCall()
	fmt.Printf("after bad call\r\n")
}

func main() {
	fmt.Printf("call test\r\n")
	testRecover()
}
