package behavioral

import "fmt"

/**
迭代器模式
简介：迭代器模式是一种行为设计模式， 让你能在不暴露集合底层表现形式 （列表、 栈和树等） 的情况下遍历集合中所有的元素。
*/
type ICollection interface {
	createIterator()
}
type Collection struct {
	ElemList []*Elem
}

func (c Collection) createIterator() IIterator {
	return &Iterator{
		index:    0,
		elemList: c.ElemList,
	}
}

type IIterator interface {
	hasNext() bool
	getNext() *Elem
}
type Iterator struct {
	index    int
	elemList []*Elem
}

func (i Iterator) hasNext() bool {
	if i.index < len(i.elemList) {
		return true
	}
	return false
}
func (i *Iterator) getNext() *Elem {
	if i.hasNext() {
		i.index++
		return i.elemList[i.index-1]
	}
	return nil
}

type Elem struct {
	Name string
}

func TestIterator() {
	elem1 := Elem{
		Name: "elem1",
	}
	elem2 := Elem{
		Name: "elem2",
	}
	collection := Collection{ElemList: []*Elem{&elem1, &elem2}}
	iterator := collection.createIterator()

	for iterator.hasNext() {
		fmt.Println(iterator.getNext().Name)
	}
}
