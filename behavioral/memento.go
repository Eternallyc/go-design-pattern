package behavioral

import "fmt"

/**
备忘录模式
	简介：保存一个对象的某个状态，以便在适当的时候恢复对象。
	例子：比如VMware里的快照，游戏世界中单机游戏的存档
	原理：用单机游戏中的存档来解释这个模式，即知需要实现这个功能，需要一个存档中心来保存所有的存档，自己电脑玩的是一个新的存档（还没有存）
		 并且电脑可以将当前游戏状态进行存档
	应用场景：浏览器回退，数据库备份与还原，编辑器撤销与重做，虚拟机生成快照与恢复，Git版本管理，棋牌游戏悔棋
	下面用单机游戏的例子来编码实现
*/
//自己当前的游戏进度（没有存入存档中心）
type originator struct {
	state string //游戏进度
}

//将当前游戏的进度进行存档
func (o originator) createMemento() *memento {
	return &memento{state: o.state}
}

//读档
func (o *originator) restoreMemento(m *memento) {
	o.state = m.getSavedState()
}
func (o *originator) setState(state string) {
	o.state = state
}
func (o *originator) getState() string {
	return o.state
}

//存档
type memento struct {
	state string //游戏进度
}

//返回游戏进度
func (m *memento) getSavedState() string {
	return m.state
}

//存档中心
type caretaker struct {
	mementoList []*memento
}

//增加存档
func (c *caretaker) addMemento(m *memento) {
	c.mementoList = append(c.mementoList, m)
}

//获取存档
func (c *caretaker) getMemento(index int) *memento {
	return c.mementoList[index]
}

//测试备忘录模式
func TestMemento() {
	//创建存档中心
	caretaker := &caretaker{
		mementoList: make([]*memento, 0),
	}
	//玩游戏（并初始化当前游戏进度）
	originator := &originator{
		state: "关卡1",
	}
	//存档
	fmt.Printf("当前游戏进度: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	//再存档
	originator.setState("关卡2")
	fmt.Printf("当前游戏进度: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	//再存档
	originator.setState("关卡3")
	fmt.Printf("当前游戏进度: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	//读档m
	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("读档。。。当前进度: %s\n", originator.getState())
	//读档
	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("读档。。。当前进度: %s\n", originator.getState())
}
