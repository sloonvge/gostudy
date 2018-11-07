package main

import (
	"./operation"
	"fmt"
)

func main() {
	op := operation.CreatOperate("+")

	op.Set(1, 2)
	fmt.Println(op.GetResult())


}
