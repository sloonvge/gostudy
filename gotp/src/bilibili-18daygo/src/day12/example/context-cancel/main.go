package main

import (
	"context"
	"fmt"
)

func test() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("exit")
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := gen(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}


func main() {
	test()
}
