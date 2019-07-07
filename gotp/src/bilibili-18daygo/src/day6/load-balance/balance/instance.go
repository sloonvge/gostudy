package balance

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (p *Instance) Port() int {
	return p.port
}

func (p *Instance) Host() string {
	rerturn p.host
}