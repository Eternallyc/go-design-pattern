package creational

import (
	"fmt"
	"strconv"
	"sync"
)

/**
单例模式
简介：单例模式用来保证某个对象的实例个数为1
实现方式：go语言有些语言特性，比如init和sync.once ，所以相比较其他语言而言比较好实现
	1.Using init() functions
	2.Using sync.Once in sync package (golang 最好的线程安全的实现方式 )
*/
var people People

var once sync.Once

func TestSingleton() {
	for i := 0; i < 5; i++ {
		once.Do(func() {
			people = People{
				Hat:     "帽子" + strconv.Itoa(i+1),
				Clothes: "衣服" + strconv.Itoa(i+1),
				Pants:   "裤子" + strconv.Itoa(i+1),
				Sock:    "袜子" + strconv.Itoa(i+1),
				Shoes:   "鞋子" + strconv.Itoa(i+1),
			}
			fmt.Println("执行次数  ", i+1)
		})
	}
	//输出信息
	people.ShowPeopleInfo()
	//总结：虽然我用for循环执行了几次，但是sync.once的do函数保证里面的逻辑只执行一次，也就达到了只实例化一次对象的效果
}
