package main

import "14subscribe/go/pattern"

/** 
* Created by wanjx in 2018/12/24 22:34
**/
 
func main() {
	qt := new(pattern.Subscribe)
	ts1 := &pattern.StockBrowse{"Q", qt}
	ts2 := &pattern.StockBrowse{"W", qt}
	ts3 := &pattern.StockBrowse{"E", qt}

	qt.Add(ts1)
	qt.Add(ts2)
	qt.Add(ts3)

	qt.Action = "Boss Comming!"
	qt.Inform()

}