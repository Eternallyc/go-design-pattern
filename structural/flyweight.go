package structural

import (
	"fmt"
	"strconv"
)

/**
享元模式
简介：享元模式就是将大量对象共有的部分抽取出来，供这些对象共享使用。而这些元素不同的部分，会以参数的形式注入具体享元的相关方法中。
本质：缓存共享对象，降低内存消耗
*/
//People接口
type People interface {
	PrintInfo(info PeopleInfo)
}

//共有的信息
type Student struct {
	Id string //身份证
}

//不同的信息
type PeopleInfo struct {
	Name string //姓名
	Sex  string //性别
}

func NewStudent(id string) *Student {
	return &Student{Id: id}
}
func (p *Student) PrintInfo(info PeopleInfo) {
	//打印共有的信息
	fmt.Println("id:" + p.Id)
	//打印不同的信息
	fmt.Println("name" + info.Name + "\n" + "sex:" + info.Sex)
}

type PeopleFactory struct {
	p map[string]People
}

//建立工厂时需要建立可以存入大量的People实例的map
func NewPeopleFactory() *PeopleFactory {
	return &PeopleFactory{p: make(map[string]People)}
}

//获得实例
func (p *PeopleFactory) GetPeople(id string) People {
	people := p.p[id]
	if people == nil {
		people = NewStudent(id)
		p.p[id] = people
		fmt.Println("the flyweight object is created")
	} else {
		fmt.Println("the flyweight object has been created")
	}
	return people
}

func TestFlyweight() {
	f := NewPeopleFactory()
	for i := 0; i < 5; i++ {
		f.GetPeople("a").PrintInfo(PeopleInfo{
			Name: "name" + strconv.Itoa(i),
			Sex:  "sex" + strconv.Itoa(i),
		})
	}
}
