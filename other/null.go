package other

import "fmt"

/**
空对象模式
简介：一个空对象取代nil对象实例的检查，nil对象不是检查空值，而是反应一个不做任何动作的关系。这样的nil对象也可以在数据不可用的时候提供默认的行为。
应用场景：我们在获取配置静态库的时候，如果获取不到配置静态库，可以返回一个默认的配置
*/
type IConfig interface {
	GetName() string
}

type Config struct {
	Name string
}

func NewConfig(name string) IConfig {
	return Config{Name: name}
}

func (c Config) GetName() string {
	return c.Name
}

type NullConfig struct {
	Name string
}

func NewNullConfig() IConfig {
	return NullConfig{}
}

func (n NullConfig) GetName() string {
	return "空配置"
}

//工厂
type CustomerFactory struct {
	configs []string
}

func (c CustomerFactory) GetConfig(name string) IConfig {
	for _, value := range c.configs {
		if value == name {
			return NewConfig(name)
		}
	}
	return NewNullConfig()
}

func TestNull() {
	//初始化工厂
	factory := CustomerFactory{
		configs: []string{"配置1", "配置2"},
	}

	//获取配置
	config1 := factory.GetConfig("配置1")
	config2 := factory.GetConfig("配置2")
	config3 := factory.GetConfig("配置3")

	//使用配置
	fmt.Println(config1.GetName())
	fmt.Println(config2.GetName())
	fmt.Println(config3.GetName())
}
