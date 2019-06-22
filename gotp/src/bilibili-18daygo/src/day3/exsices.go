package main

import (
	"fmt"
	"time"
)

/** 
* Created by wanjx in 2019/6/22 12:42
**/

func nowTime() {
	now := time.Now()
	fmt.Println(now.Format("2006:01:02"), now)
}

func main() {
	nowTime()
}