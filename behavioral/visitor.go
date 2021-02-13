package behavioral

import "fmt"

/**
访问者模式
简介：主要将数据结构与数据操作分离。
应用场景：
	1、对象结构中对象对应的类很少改变，但经常需要在此对象结构上定义新的操作。
2、需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免让这些操作"污染"这些对象的类，也不希望在增加新操作时修改这些类。
*/

type IEntity interface {
	FullName()
}
type Entity struct {
	Name string
	Age  int
}

func (e Entity) FullName() {

}

type ICalculation interface {
	//计算操作1
	Func1(entity IEntity)
}
type Calculation struct {
}

func (c Calculation) Func1(entity Entity) {
	fmt.Println(entity.Name, entity.Age)
}

func TestVisitor() {
	entity := Entity{
		Name: "visitor",
		Age:  22,
	}
	calculate := Calculation{}
	calculate.Func1(entity)
}
