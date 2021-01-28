package creational

import "fmt"

/**
工厂模式
简介：定义一个创建对象的接口，让其子类自己决定实例化哪一个工厂类，工厂模式使其创建过程延迟到子类进行。
使用场景：
	1、客户端可以根据自身需要创建很多同类型对象中的一种

目的：解决了接口的选择问题，我们现在只需要知道要获取的类名称，然后从工厂中直接获取就可以
*/

//工厂
type Factory struct {
}

func (f Factory) Create(typ int) Instance {
	if typ == 1 {
		return ConcreteInstance1{}
	} else {
		return ConcreteInstance2{}
	}
}

//
type Instance interface {
	DoSomething() //实例动作
}
type ConcreteInstance1 struct {
}

func (c ConcreteInstance1) DoSomething() {
	fmt.Println("concrete instance1 do something")
}

type ConcreteInstance2 struct {
}

func (c ConcreteInstance2) DoSomething() {
	fmt.Println("concrete instance2 do something")
}

//测试工厂模式
func TestFactory() {
	//初始化工厂
	f := Factory{}
	instance := f.Create(1)
	instance.DoSomething()
}
