package behavioral

import "fmt"

/*
	简介：简单来说责任链就是当你需要进行某项操作的时候，而这项操作是一套处理流程，而你必须沿着这个处理流程进行执行
	例子：比如游戏世界你买一样东西，也可以用责任链模式来做，在游戏微服务时代，各个业务模块划分为不同的BC
			首先，调用仓库BC检验你的元宝是否足够，不够就报错，够就进行下一步
			其次，去商店BC（也就是本BC）进行购买
			再次，调用仓库BC进行扣除元宝
			最后，进行日志记录
		但是在平时这类操作就是一个接口~
	这里为了更加容易理解，采用病人去看病的流程来解释责任链模式
	参考：https://refactoringguru.cn/design-patterns/chain-of-responsibility/go/example
*/

//病人
type patient struct {
	Name string //姓名
}

//处理者接口
type Department interface {
	Execute(*patient)
	SetNext(Department)
}

//接待处
type Reception struct {
	NextHandle Department //下个部门
}

func NewReception() Department {
	return &Reception{}
}
func (c *Reception) Execute(p *patient) {
	fmt.Println("接待病人")
	c.NextHandle.Execute(p)
}
func (c *Reception) SetNext(handle Department) {
	c.NextHandle = handle
}

//医生
type Doctor struct {
	NextHandle Department //下个部门
}

func NewDoctor() Department {
	return &Doctor{}
}
func (d *Doctor) Execute(p *patient) {
	fmt.Println("看病")
	d.NextHandle.Execute(p)
}

func (d *Doctor) SetNext(handle Department) {
	d.NextHandle = handle
}

//药房
type Pharmacy struct {
	NextHandle Department //下个部门
}

func NewPharmacy() Department {
	return &Pharmacy{}
}
func (pharmacy *Pharmacy) Execute(p *patient) {
	fmt.Println("拿药")
	pharmacy.NextHandle.Execute(p)
}

func (pharmacy *Pharmacy) SetNext(handle Department) {
	pharmacy.NextHandle = handle
}

//缴费处
type PaymentOffice struct {
	NextHandle Department //下个部门
}

func NewPaymentOffice() Department {
	return &PaymentOffice{}
}
func (paymentOffice *PaymentOffice) Execute(p *patient) {
	fmt.Println("缴费")
}

func (paymentOffice *PaymentOffice) SetNext(handle Department) {
	paymentOffice.NextHandle = handle
}

//测试责任链模式
func TestChainResponsibility() {
	//接待处
	reception := NewReception()
	//医生
	doctor := NewDoctor()
	//药房
	pharmacy := NewPharmacy()
	//缴费处
	paymentOffice := NewPaymentOffice()
	//责任链
	//药房的下个部门是缴费处
	pharmacy.SetNext(paymentOffice)
	//医生的下个部门是药房
	doctor.SetNext(pharmacy)
	//接待处的下个部门是医生
	reception.SetNext(doctor)

	//一个病人
	patient := patient{
		Name: "我是病人",
	}
	//看病
	reception.Execute(&patient)
}
