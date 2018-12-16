package pattern

import "fmt"

/**
* Created by wanjx in 2018/12/8 21:09
**/

//
//type PaperA struct {
//
//}
//
//func (p *PaperA) Question1() {
//	fmt.Println("Question 1")
//	fmt.Println("Answer: A")
//}
//
//func (p *PaperA) Question2() {
//	fmt.Println("Question 1")
//	fmt.Println("Answer B")
//}
//
//type PaperB struct {
//
//}
//
//func (p *PaperB) Question1() {
//	fmt.Println("Question 1")
//	fmt.Println("Answer C")
//}
//
//func (p *PaperB) Question2() {
//	fmt.Println("Question 2")
//	fmt.Println("Answer D")
//}

type IPaper interface {
	Answer1() (string)
	Answer2() (string)
}

type Paper struct {
	Answer IPaper
}

func (p *Paper) Question1() {
	fmt.Println("Question1")
	fmt.Println("Answer" + p.Answer.Answer1())
}

func (p *Paper) Question2(){
	fmt.Println("Question2")
	fmt.Println("Answer" + p.Answer.Answer2())
}

func (p *Paper) Answer1() (ret string){
	ret = ""
	return
}

func (p *Paper) Answer2() (ret string) {
	ret = ""
	return
}

type PaperA struct {
	Paper
}

func NewPaperA() (p *PaperA){
	p = &PaperA{}
	p.Answer = p
	return
}

func (p *PaperA) Answer1() (ret string){
	ret = "A"
	return
}

func (p *PaperA) Answer2() (ret string) {
	ret = "B"
	return
}

type PaperB struct {
	Paper
}

func NewPaperB() (p *PaperB){
	p = &PaperB{}
	p.Answer = p
	return
}

func (p *PaperB) Answer1() (ret string){
	ret = "C"
	return
}

func (p *PaperB) Answer2() (ret string){
	ret = "D"
	return
}