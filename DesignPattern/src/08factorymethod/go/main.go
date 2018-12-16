package main

import (
	"08factorymethod/go/pattern"
	"fmt"
)

func main() {
	factory := pattern.NewAddFactory()
	op := factory.CreateOperation()
	op.Set(1, 2)
	fmt.Println(op.GetResult())
}
