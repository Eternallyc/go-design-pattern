package main

import (
	"go-design-pattern/behavioral"
	"go-design-pattern/structural"
)

func main() {

	//行为型
	//责任链模式
	//behavioral.TestChainResponsibility()
	//中介者模式
	//behavioral.TestMediator()
	//备忘录模式
	//behavioral.TestMemento()
	//策略模式
	//behavioral.TestStrategy()
	//模版方法模式
	behavioral.TestTemplate()
	//创建型
	//建造者模式
	//creational.TestBuilder()

	//结构型
	//桥接模式
	//structural.TestBridge()
	//享元模式
	//structural.TestFlyweight()
	//代理模式
	structural.TestProxy()
}
