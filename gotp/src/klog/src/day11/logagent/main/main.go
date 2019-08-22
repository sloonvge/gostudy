package main

import "fmt"

func main() {
	err := loadConf()
	if err != nil {
		fmt.Println("load config failed, err:", err)
		return
	}
}
