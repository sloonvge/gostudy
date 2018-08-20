package main

//import "ch1/ch13"
//
//func main() {
//	ch13.Dup3()
//}

type Param map[string]interface {}

type Show struct {
	 Param
}

func main() {
	s := Show{}
	s.Param = make(map[string]interface{})
	s.Param["RMB"] = 1000
}