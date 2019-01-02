package main

import "13builder/go/pattern"

/** 
* Created by wanjx in 2018/12/22 21:25
**/

func main()  {
	pt := new(pattern.PersonThinBuilder)
	pd := pattern.NewPersionDirector(pt)
	pd.CreatePerson()
}
