package structural

import "fmt"

/**
代理模式
简介：当无法或不想直接引用某个对象或访问某个对象存在困难时，可以通过代理对象来间接访问。
目标：：一是保护目标对象，二是增强目标对象。
案例场景：在游戏公司中，如果该公司内部如果使用DDD和中台，那么会将登陆作为通用业务，在这种场景下，每个游戏会依赖到这个通用业务，
		那么，每个游戏的登录其实可以看做是一种代理，每个游戏根据自身需求基于登录通用业务进行拓展
*/
type IUser interface {
	Login(userId, password string) (bool, error)
}
type User struct {
}

func (u User) Login(userId, password string) (bool, error) {
	fmt.Println("校验帐号密码")
	return true, nil
}

type Game struct {
	user User
}

func (g Game) Login(userId, password string) (bool, error) {
	fmt.Println("校验请求的合法性")
	fmt.Println("校验帐号是否冻结")
	fmt.Println("校验帐号数据是否异常")
	return g.user.Login(userId, password)
}

func TestProxy() {
	user := User{}
	game := Game{user: user}
	game.Login("1", "1")
}
