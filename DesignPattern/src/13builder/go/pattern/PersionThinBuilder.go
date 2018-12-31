package pattern

import "fmt"

/**
* Created by wanjx in 2018/12/22 21:30
**/
 
type PersonThinBuilder struct {

}

func (p *PersonThinBuilder) BuildHead() {
	fmt.Println("Head")
}

func (p *PersonThinBuilder) BuildBody() {
	fmt.Println("Thin Body")
}

func (p *PersonThinBuilder) BuildArmLeft() {
	fmt.Println("ArmLeft")
}

func (p *PersonThinBuilder) BuildArmRight() {
	fmt.Println("ArmRight")
}

func (p *PersonThinBuilder) BuildLegLeft()  {
	fmt.Println("LegLeft")
}

func (p *PersonThinBuilder) BuildLegRight()  {
	fmt.Println("LegRight")
}