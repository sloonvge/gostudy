package pattern

/** 
* Created by wanjx in 2018/11/27 22:49
**/

type ICloneable interface {
	Clone() (r Resume)
	SetPersonalInfo(sex string, age int)
	SetWorkInfo(timeArea string, company string)
}
 
