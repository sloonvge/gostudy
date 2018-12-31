package pattern

/** 
* Created by wanjx in 2018/12/22 21:36
**/

type PersonDirector struct {
	Person Builder
}

func NewPersionDirector(builder Builder) (p *PersonDirector) {
	p = &PersonDirector{Person:builder}
	return
}

func (p *PersonDirector) CreatePerson()  {
	p.Person.BuildHead()
	p.Person.BuildBody()
	p.Person.BuildArmLeft()
	p.Person.BuildArmRight()
	p.Person.BuildLegLeft()
	p.Person.BuildLegRight()
}
