package structural

import "fmt"

/**
 桥接模式
	简介：
*/
//书本
type book interface {
	printColor()
	setColor(color color)
}

//浪潮之巅
type lCZD struct {
	color color
}

func (l *lCZD) printColor() {
	fmt.Print("浪潮之巅: ")
	l.color.printColor()
}
func (l *lCZD) setColor(color color) {
	l.color = color
}

//见识
type experience struct {
	color color
}

func (e *experience) printColor() {
	fmt.Print("见识: ")
	e.color.printColor()
}
func (e *experience) setColor(color color) {
	e.color = color
}

//打印机
type color interface {
	printColor()
}

//蓝色
type colorBlue struct {
}

func (c colorBlue) printColor() {
	fmt.Println("蓝色")
}

//红色
type colorRed struct {
}

func (c colorRed) printColor() {
	fmt.Println("红色")
}

func TestBridge() {
	//初始化颜色
	blue := colorBlue{}
	red := colorRed{}
	//初始化浪潮之巅这本书，设置为蓝色书籍
	l := lCZD{}
	l.setColor(blue)
	l.printColor()
	//设置为红色书籍
	l.setColor(red)
	l.printColor()

	//初始化见识这本书
	e := experience{}
	e.setColor(blue)
	e.printColor()

	e.setColor(red)
	e.printColor()
}
