package pattern

type Operation struct {
	numberA float64
	numberB float64
}

type Operate interface {
	GetResult() float64
	Set(numberA float64, numberB float64)
}

type IFactory interface {
	CreateOperation() Operate
}
