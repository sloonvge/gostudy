package proxy

type Proxy struct {
	GG *Persuit
}

func NewProxy(mm *Girl) (p *Proxy) {
	p = &Proxy{GG: NewPersuit(mm)}
	return
}

func (p *Proxy) GiveDolls() {

	p.GG.GiveDolls()
}

func (p *Proxy) GiveFlowers() {
	p.GG.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	p.GG.GiveChocolate()
}


