package pattern

import "fmt"

/** 
* Created by wanjx in 2018/12/31 10:53
**/
 
type NBABrowse struct {
	Name string
	Sub Subscriber
}

func (n *NBABrowse) Update() {
	fmt.Printf("%s %s. Close pc!\n", n.Sub.ActionT(), n.Name)
}