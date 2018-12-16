package pattern

type AddFactory Operation

func NewAddFactory() (a *AddFactory){
	a = &AddFactory{}
	return
}

func (a *AddFactory) CreateOperation() (o Operate) {
	o = NewOperationAdd()
	return
}

type SubFactory Operation

func NewSubFactory() (a *SubFactory){
	a = &SubFactory{}
	return
}

func (s *SubFactory) CreateOperation() (o Operate) {
	o = NewOperationSub()
	return
}

type MulFactory Operation
func NewMulFactory() (a *MulFactory){
	a = &MulFactory{}
	return
}
func (s *MulFactory) CreateOperation() (o Operate) {
	o = NewOperationMul()
	return
}

type DivFactory Operation
func NewDivFactory() (a *DivFactory){
	a = &DivFactory{}
	return
}
func (s *DivFactory) CreateOperation() (o Operate) {
	o = NewOperationDiv()
	return
}
