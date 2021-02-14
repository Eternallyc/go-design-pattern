package behavioral

import "fmt"

/**
命令模式
简介：将一系列的请求命令封装起来，不直接调用真正执行者的方法，这样比较好扩展。
应用场景：需要向某些对象发送请求，但是并不知道请求的接收者是谁，也不知道被请求的操作是什么。
		在基础平台组，会为各大业务组提供SDK，然而SDK内部是调用另外的服务，所以这里也可以用命令模式
*/

//遥控器
type RemoteControl struct {
	command Command
}

//按下按钮
func (r RemoteControl) press() {
	r.command.Execute()
}

type Command interface {
	//执行命令
	Execute()
}

//打开命令
type onCommand struct {
	//具体执行者
	device Device
}

func (on onCommand) Execute() {
	on.device.On()
}

//关闭命令
type offCommand struct {
	//具体执行者
	device Device
}

func (off offCommand) Execute() {
	off.device.Off()
}

type Device interface {
	On()
	Off()
}

//具体命令执行者-电视
type tv struct {
}

func (tv tv) On() {
	fmt.Println("打开电视")
}
func (tv tv) Off() {
	fmt.Println("关闭电视")
}

//测试命令模式
func TestCommand() {
	//初始化具体命令执行者
	tv := tv{}
	//初始化开关命令
	on := onCommand{device: tv}
	off := offCommand{device: tv}
	//初始化遥控器
	control := RemoteControl{command: on}
	control.press()

	control = RemoteControl{command: off}
	control.press()
}
