package behavioral

import "fmt"

/**
观察者模式
简介:观察者模式是一种行为设计模式， 允许你定义一种订阅机制， 可在对象事件发生时通知多个 “观察” 该对象的其他对象。
应用场景：公众号订阅，游戏活动订阅等
*/

//被观察者
type iSubject interface {
	AddObserve(observe iObserve) //添加观察者
	NotifyAll()                  //当有修改的时候，通知所有订阅者
}
type subject struct {
	observes []iObserve
}

func (s *subject) AddObserve(observe iObserve) {
	if len(s.observes) == 0 {
		s.observes = make([]iObserve, 0)
	}
	s.observes = append(s.observes, observe)
}
func (s *subject) NotifyAll() {
	//主题更新
	fmt.Println("主题更新")
	for _, value := range s.observes {
		value.Notify()
	}
}

//订阅者
type iObserve interface {
	Notify() //当观察者主题发生变化时，通知订阅者
}
type observe1 struct {
}

func (o observe1) Notify() {
	fmt.Println("订阅者1 收到变更")
}

type observe2 struct {
}

func (o observe2) Notify() {
	fmt.Println("订阅者2 收到变更")
}

//测试观察者模式
func TestObserve() {
	//初始化主题
	subject := subject{}

	//初始化订阅者1并订阅主题
	o1 := observe1{}
	subject.AddObserve(o1)
	//初始化订阅者2
	o2 := observe2{}
	subject.AddObserve(o2)

	//主题更新
	subject.NotifyAll()
}
