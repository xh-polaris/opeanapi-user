package provider

import (
	"github.com/google/wire"
	"github.com/xhpolaris/opeanapi-user/biz/adaptor/controller"
	"github.com/xhpolaris/opeanapi-user/biz/application/service"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/config"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/mapper/key"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/mapper/user"
)

var AllProvider = wire.NewSet(
	ControllerSet,
	ApplicationSet,
	InfrastructureSet,
)

var ControllerSet = wire.NewSet(
	controller.AuthControllerSet,
	controller.MoneyControllerSet,
)

var ApplicationSet = wire.NewSet(
	service.KeyServiceSet,
	service.UserServiceSet,
	service.MoneyServiceSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	MapperSet,
)

var MapperSet = wire.NewSet(
	key.NewMongoMapper,
	user.NewMongoMapper,
)
