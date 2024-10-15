package adaptor

import (
	"github.com/xh-polaris/opeanapi-user/biz/adaptor/controller"
)

type UserServer struct {
	controller.IAuthController
	controller.IMoneyController
}
