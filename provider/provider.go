package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-user/biz/adaptor/controller"
	"github.com/xh-polaris/openapi-user/biz/application/service"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/account"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/key"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/user"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/transaction"
)

var UserServerProvider = wire.NewSet(
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
	service.AccountServiceSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	MapperSet,
	TransactionSet,
)

var MapperSet = wire.NewSet(
	key.NewMongoMapper,
	user.NewMongoMapper,
	account.NewMongoMapper,
)

var TransactionSet = wire.NewSet(
	transaction.NewUserTransaction,
)
