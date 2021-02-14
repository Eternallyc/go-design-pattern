package structural

import "fmt"

/**
适配器模式
简介：用于将一个对象转换成另外一个对象可识别的对象
应用场景：
	1.系统需要使用现有的类，而此类的接口不符合系统的需要。
案例解释：下面我用日志记录来显示适配器模式的作用
		一般在公司中基础平台组中，我们会自己封装一些开源库而不直接使用，主要是因为开源库的一些接口不符合我们公司使用习惯，或者晦涩难懂等情况，
		所以我们会做一层适配，提高产品中信和研发中心开发人员的工作效率;
注意点：适配器模式是解决新旧接口不一致导致出现了客户端无法得到满足的问题
*/
//自己用的日志
type ILog interface {
	//日志记录
	Logging()
}
type MyLog struct {
	zap *Zap
}

func (m MyLog) Logging() {
	fmt.Println("进行日志数据处理")
	//进行日志记录
	m.zap.ZapLogging()
}

//mock zap日志框架
type Zap struct {
}

func (z Zap) ZapLogging() {
	fmt.Println("zap 日志框架记录日志")
}

func TestAdapter() {
	zap := Zap{}
	myLog := MyLog{zap: &zap}
	//进行日志记录
	myLog.Logging()
}
