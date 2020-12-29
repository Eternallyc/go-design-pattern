package behavioral

import (
	"context"
	"fmt"
)

/**
中介者模式
	简介：用一个中介对象来封装一系列的对象交互，中介者使各对象不需要显示地互相引用，从而使其耦合松散，而且可以独立地改变他们之间的交互
	案例解释：当具体对象注册到中介者中，客户端需要更新某个具体对象值的时候，就可以直接将注册的type传入中介者来调用具体的对象
	应用场景：游戏中的道具、皮肤和英雄等，其实可以看成一个游戏商品，比如英雄联盟中英雄，皮肤，道具，精粹和点券
	优点：可以实现调用者与一系列同类对象解耦
*/
//商品种类
type ItemType uint8

var (
	HeroItem        ItemType = 1 //英雄
	SkinItem        ItemType = 2 //皮肤
	PropsItem       ItemType = 3 //游戏道具
	EssenceItem     ItemType = 4 //精粹
	PointCouponItem ItemType = 5 //点券

)

//中介者接口
type Mediator interface {
	//执行更新操作
	Update(ctx context.Context, typ ItemType, count int)
}

//中介者实现--游戏商品
type GameMerchandise struct {
	Item map[ItemType]IItem
}

func (game GameMerchandise) AddItem(typ ItemType, item IItem) {
	game.Item[typ] = item
}
func (game GameMerchandise) Update(ctx context.Context, typ ItemType, count int) {
	//获得具体执行者(一些错误判断就忽略了)
	executor := game.Item[typ]
	//执行操作
	executor.Update(ctx, count)
}

//商品接口
type IItem interface {
	//更新物品数量
	Update(ctx context.Context, count int)
}

//英雄
type Hero struct {
}

func NewHero(mediator GameMerchandise) Hero {
	//将自己注册进去
	h := Hero{}
	mediator.AddItem(HeroItem, h)
	return h
}
func (h Hero) Update(ctx context.Context, count int) {
	//进行购买英雄
	fmt.Println("购买英雄")
}

func TestMediator() {
	ctx := context.Background()
	//初始化游戏
	g := GameMerchandise{
		Item: make(map[ItemType]IItem, 0),
	}
	h := NewHero(g)
	//将英雄注册
	g.AddItem(HeroItem, h)
	//执行更新操作
	g.Update(ctx, HeroItem, 1)
}
