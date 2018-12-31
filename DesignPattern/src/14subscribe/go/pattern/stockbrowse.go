package pattern

import "fmt"

/** 
* Created by wanjx in 2018/12/24 22:38
**/
 
type StockBrowse struct {
	Name string
	Sub Subscriber
}

func (s *StockBrowse) Update() {
	fmt.Printf("%s %s. Close pc!\n", s.Sub.ActionT(), s.Name)
}