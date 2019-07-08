package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputReaderEx1() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input: ")
	str, err :=inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("read: %s", str)
}

func main() {
	inputReaderEx1()
}