package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c += 1
	return int(*c), nil
}

func (c *WordCounter) Value() int {
	return int(*c)
}

func (c *WordCounter) CountWord(name string) {
	var line int
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c += 1
	}
	reader := bufio.NewReader(f)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		line += 1

	}

	fmt.Println(line)
}


