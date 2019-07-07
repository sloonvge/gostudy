package balance

type Balancer interface {
	DoBalance() (inst *Instance, err error)
}