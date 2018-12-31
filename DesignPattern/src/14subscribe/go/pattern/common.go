package pattern

import "fmt"

/** 
* Created by wanjx in 2018/12/31 10:50
**/

type Browse struct {
	Name string
	Sub *Subscribe
}

func (s *Browse) Update() {
	fmt.Printf("%s %s. Close pc!\n", s.Sub.ActionT(), s.Name)
}

type Browser interface {
	Update()
}

type Subscriber interface {
	Add(b Browser)
	Inform()
	ActionT() string
}