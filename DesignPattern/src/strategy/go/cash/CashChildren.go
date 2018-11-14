package cash

type Normal struct {

	Super
}

func NewNormal() (n *Normal) {
	n = &Normal{}
	return
}

func (n *Normal) Cash(m float64) (ret float64){
	ret = m
	return
}


type Rebate struct {
	Rate float64

	Super
}

func NewRebate(r float64) (n *Rebate) {
	n = &Rebate{Rate: r}
	return
}

func (r *Rebate) Cash(m float64) (ret float64) {
	if r.Rate != 0.0 {
		ret = m * r.Rate
	} else {
		panic("请重新输入折扣")
	}
	return
}


type Return struct {
	Condition float64
	Back float64

	Super
}

func NewReturn(c float64, b float64) (n *Return) {
	n = &Return{Condition: c, Back: b}
	return
}

func (r *Return) Cash(m float64) (ret float64) {
	if m > r.Condition {
		ret = m - r.Back * float64(int(m / r.Condition))
	} else {
		ret = m
	}
	return
}