package behavioral

import "fmt"

/**
模板方法
简介：定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。
	模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。
应用：1、有多个子类共有的方法，且逻辑相同。
	 2、重要的、复杂的方法，可以考虑作为模板方法。
*/
type ITemplate interface {
	DoInit()
}
type Init struct {
}

func (t Init) DoInit() {

}
func (t Init) InitBg() {
	fmt.Println("初始化背景")
}
func (t Init) InitMusic() {
	fmt.Println("初始化音乐")
}
func (t Init) InitBoss() {
	fmt.Println("初始化Boss")
}

type Role struct {
	Init
}

func (r Role) DoInit() {
	fmt.Println("角色1初始化")
}

func doInit(r Role) {
	r.InitBg()
	r.InitMusic()
	r.DoInit()
	r.InitBoss()
}

func TestTemplate() {
	r := Role{}
	doInit(r)
}
