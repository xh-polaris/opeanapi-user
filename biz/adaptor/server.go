package adaptor

import "github.com/xhpolaris/opeanapi-user/biz/adaptor/controller"

type UserServer struct {
	controller.IAuthController
	controller.IMoneyController
}

func NewUserServer() *UserServer {
	return &UserServer{
		IAuthController:  controller.NewAuthController(),
		IMoneyController: controller.NewMoneyController(),
	}
}
