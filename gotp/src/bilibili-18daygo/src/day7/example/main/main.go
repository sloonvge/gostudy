package main

import (
	"os"
	"fmt"
)

func CmdArgsExample() {
	fmt.Printf("len of args: %d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("%d: %s\n", i, v)
	}
}

func main() {
	CmdArgsExample()
}