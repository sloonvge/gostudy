package cash

import "fmt"

type Item struct {
	N int
	Name string
	Price float64
	H string
}

func (i *Item) CalCost() (cost float64){
	var dz float64

	switch i.H {
	case "八折":
		dz = 0.8
	case "七折":
		dz = 0.7
	default:
		dz = 1.0
		i.H = "正常收费"
	}
	cost = i.Price * float64(i.N) *dz
	return
}

func (i *Item) String() (s string){
	s = fmt.Sprintf("商品: %s 数量: %d 单价: %.3f 打折: %s 总计: %.3f",
		i.Name, i.N, i.Price, i.H, i.CalCost())
	return
}


type Items struct {
	I []*Item
	H string
}


func (is Items) CalCost()  (cost float64){
	for _, i := range is.I {
		cost += i.CalCost()
	}
	return
}


func (is Items) Confirm() {
	for _, i := range is.I {
		fmt.Println(i)
	}
	fmt.Printf("合计: %.3f", is.CalCost())
}
