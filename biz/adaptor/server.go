package adaptor

import (
	"github.com/xh-polaris/openapi-user/biz/adaptor/controller"
)

type UserServer struct {
	controller.IAuthController
	controller.IMoneyController
}
