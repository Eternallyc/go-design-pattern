package behavioral

import "fmt"

/**
策略模式
简介：   定义一系列的算法，把它们一个个封装起来，并且使它们可相互替换。
应用实例:
	1.单个游戏中的场景换肤
	2.游戏中的对于同种玩法的多种计算公式
*/
type IStrategy interface {
	ChangeSkin() //改变皮肤
}

//主场景
type Main struct {
	Skin Skin //当前皮肤
}

func (m Main) ChangeSkin() {
	m.Skin.ShowSkin()
}

type Skin interface {
	ShowSkin() //显示皮肤
}

type SpringSkin struct {
}

func (s SpringSkin) ShowSkin() {
	fmt.Println("spring skin")
}

type SummerSkin struct {
}

func (s SummerSkin) ShowSkin() {
	fmt.Println("summer skin")
}

func TestStrategy() {
	spring := SpringSkin{}
	summer := SummerSkin{}

	main1 := Main{Skin: spring}
	main1.ChangeSkin()

	main2 := Main{Skin: summer}
	main2.ChangeSkin()
}
