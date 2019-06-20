package main

import (
	"os"
	"fmt"
)

func main() {
	var o = os.Getenv("GOOS")
	fmt.Printf("The os: %s\n", o)
	p := os.Getenv("PATH")
	fmt.Printf("The path: %s\n", p)
}