package main

import (
	"09prototype/go/pattern"
	"fmt"
)

/*
 * Created by  in 2018/11/27 
 */

func main(){
	resume1 := pattern.NewResume("Tom")
	resume1.SetPersonalInfo("male", 26)
	resume1.SetWorkInfo("xxx-xxx", "zzz")

	//resume2 := pattern.NewResume("Tom")
	//resume2.SetPersonalInfo("male", 26)
	//resume2.SetWorkInfo("xxx-xxx", "yyy")
	resume2 := resume1.Clone()
	resume2.SetWorkInfo("zzz-zzz", "xxx")


	fmt.Println(resume1)
	fmt.Println(resume2)
}
