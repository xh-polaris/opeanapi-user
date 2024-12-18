// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/xh-polaris/openapi-user/biz/adaptor"
	"github.com/xh-polaris/openapi-user/biz/adaptor/controller"
	"github.com/xh-polaris/openapi-user/biz/application/service"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/account"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/key"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/user"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/transaction"
)

// Injectors from wire.go:

func NewProvider() (*adaptor.UserServer, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	mongoMapper := key.NewMongoMapper(configConfig)
	userMongoMapper := user.NewMongoMapper(configConfig)
	keyService := &service.KeyService{
		KeyMongoMapper:  mongoMapper,
		UserMongoMapper: userMongoMapper,
	}
	userService := &service.UserService{
		UserMongoMapper: userMongoMapper,
	}
	authController := &controller.AuthController{
		KeyService:  keyService,
		UserService: userService,
	}
	userTransaction := transaction.NewUserTransaction(configConfig)
	moneyService := &service.MoneyService{
		UserMongoMapper: userMongoMapper,
		UserTransaction: userTransaction,
	}
	accountMongoMapper := account.NewMongoMapper(configConfig)
	accountService := &service.AccountService{
		AccountMongoMapper: accountMongoMapper,
	}
	moneyController := &controller.MoneyController{
		MoneyService:   moneyService,
		AccountService: accountService,
	}
	userServer := &adaptor.UserServer{
		IAuthController:  authController,
		IMoneyController: moneyController,
	}
	return userServer, nil
}
