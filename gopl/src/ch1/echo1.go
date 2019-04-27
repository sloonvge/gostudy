package main

import (
	"fmt"
	"strings"
	"os"
)

func echo11() {
	fmt.Println(strings.Join(os.Args, " "))
}

func echo12() {

	for i, arg := range os.Args {
		fmt.Println(i, " ", arg)
	}
}
