package pattern

/** 
* Created by wanjx in 2018/11/27 23:27
**/

type Work struct {
	TimeArea string
	Company string
}

func NewWork() (w Work) {
	w = Work{}
	return
}

func (w *Work) SetWorkInfo(timeArea string, company string) {
	w.TimeArea = timeArea
	w.Company = company
}
 
