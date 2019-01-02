package main

import (
	"fmt"
	"os"
)

func main() {
	f := os.Args[1]

	fmt.Printf("Hello, %s", f)
}
