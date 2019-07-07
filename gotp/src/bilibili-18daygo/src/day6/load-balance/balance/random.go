package balance

import (
	"math/rand"
	"errors"
)

type RandomBalance struct {

}

func (p *RandomBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}