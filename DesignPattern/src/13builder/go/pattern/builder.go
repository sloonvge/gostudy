package pattern

/** 
* Created by wanjx in 2018/12/22 21:26
**/
 
type Builder interface {
	BuildHead()
	BuildBody()
	BuildArmLeft()
	BuildArmRight()
	BuildLegLeft()
	BuildLegRight()
}