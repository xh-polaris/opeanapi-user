package controller

type IMoneyController interface {
	// 实现idl中定义的所有rpc方法
}

type MoneyController struct {
	// 引入domain层的service
}

func NewMoneyController() *MoneyController {
	// 需要依赖注入
	return &MoneyController{}
}
