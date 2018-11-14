package main

import (
	"fmt"
	"reflect"
	"strategy/go/cash"
)

func main() {
	cf := cash.CreateCash("正常收费")

	fmt.Println(reflect.TypeOf(cf))

	fmt.Println(cf.Cash(0.5))
}
