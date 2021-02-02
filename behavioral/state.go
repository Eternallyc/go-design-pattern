package behavioral

import "fmt"

/**
状态模式
简介：在状态模式中，类的行为是基于它的状态改变的。这种类型的设计模式属于行为模式。
应用场景：代码中包含大量与对象状态有关的条件语句。
	1.在LOL中，我们会有很多个英雄，而这些英雄都有四种技能，这样我们其实只需要提供四个接口，
      然后根据我们所选择的具体英雄执行具体技能。
	2.在RPG游戏中，我们也会有男女角色角色，会有不同的剧情
*/
type IState interface {
	GameStory() //游戏剧情
}

type Woman struct {
}

func (w Woman) GameStory() {
	fmt.Println("女性角色剧情....")
}

type Man struct {
}

func (m Man) GameStory() {
	fmt.Println("男性角色剧情....")
}

type IGame interface {
	GameStory()
}
type Game struct {
	Role IState
}

func (g *Game) GameStory() {
	fmt.Println("开始游戏...")
	g.Role.GameStory()
}
func (g *Game) SetRole(role IState) {
	g.Role = role
}

//测试状态模式
func TestState() {
	//初始化游戏
	game := Game{}
	//初始化男性角色
	man := Man{}
	game.SetRole(man)

	game.GameStory()
}
