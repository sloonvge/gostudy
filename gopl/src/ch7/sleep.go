package main

import (
	"flag"
	"fmt"
	"time"
)

/** 
* Created by wanjx in 2019/4/25 22:20
**/
 
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
