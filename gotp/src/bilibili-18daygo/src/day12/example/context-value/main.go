package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 123456)
	ctx = context.WithValue(ctx, "session", "test_session")
	process(ctx)

}

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 987654
	}
	fmt.Printf("ret:%d\n", ret)

	s, _ := ctx.Value("session").(string)
	fmt.Printf("session:%s\n", s)
}
