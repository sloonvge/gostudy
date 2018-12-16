package pattern

import (
	"encoding/json"
	"log"
)

type Resume struct {
	Name string
	Age int
	Sex string
	Work Work
	//TimeArea string
	//Company string
}

func NewResume(name string) (r *Resume) {
	work := NewWork()
	r = &Resume{Name:name, Work:work}
	return
}

func (r *Resume) SetPersonalInfo(sex string, age int) {
	r.Sex = sex
	r.Age = age
}

func (r *Resume) SetWorkInfo(timeArea string, company string) {
	r.Work.TimeArea = timeArea
	r.Work.Company = company
}

func (r *Resume) toString() (ret string){
	bytes, err := json.Marshal(r)
	if err != nil {
		log.Print(err)
	}
	ret = string(bytes)
	return
}

func (r *Resume) Clone() (rc Resume){
	rc.Name = r.Name
	rc.Sex = r.Sex
	rc.Age = r.Age
	//rc.Company = r.Company
	//rc.TimeArea = r.TimeArea
	rc.Work = r.Work
	return
}

