package pattern

import "log"

type Add Operation

func NewOperationAdd() (o *Add) {
	return &Add{}
}

func (o *Add) GetResult() (result float64) {
	result = o.numberA + o.numberB
	return
}
func (o *Add) Set(numberA float64, numberB float64) {
	o.numberA = numberA
	o.numberB = numberB
}


type Sub Operation

func NewOperationSub() (o *Sub) {
	return &Sub{}
}

func (o *Sub) GetResult() (result float64) {
	result = o.numberA - o.numberB
	return
}
func (o *Sub) Set(numberA float64, numberB float64) {
	o.numberA = numberA
	o.numberB = numberB
}

type Mul Operation

func NewOperationMul() (o *Mul) {
	return &Mul{}
}
func (o *Mul) Set(numberA float64, numberB float64) {
	o.numberA = numberA
	o.numberB = numberB
}

func (o *Mul) GetResult() (result float64) {
	result = o.numberA * o.numberB
	return
}


type Div Operation

func NewOperationDiv() (o *Div) {
	return &Div{}
}
func (o *Div) Set(numberA float64, numberB float64) {
	o.numberA = numberA
	o.numberB = numberB
}
func (o *Div) GetResult() (result float64) {
	if o.numberB == 0 {
		log.Println("除数不能为0")
	}
	result = o.numberA / o.numberB
	return
}
