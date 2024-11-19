package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-user/biz/application/service"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

type IMoneyController interface {
	SetRemain(ctx context.Context, req *user.SetRemainReq) (res *user.SetRemainResp, err error)
	GetAccountByTxId(ctx context.Context, req *user.GetAccountByTxIdReq) (res *user.GetAccountByTxIdResp, err error)
}

type MoneyController struct {
	MoneyService   *service.MoneyService
	AccountService *service.AccountService
}

var MoneyControllerSet = wire.NewSet(
	wire.Struct(new(MoneyController), "*"),
	wire.Bind(new(IMoneyController), new(*MoneyController)),
)

func (c *MoneyController) SetRemain(ctx context.Context, req *user.SetRemainReq) (res *user.SetRemainResp, err error) {
	return c.MoneyService.SetRemain(ctx, req)
}

func (c *MoneyController) GetAccountByTxId(ctx context.Context, req *user.GetAccountByTxIdReq) (*user.GetAccountByTxIdResp, error) {
	return c.AccountService.GetAccountByTxId(ctx, req)
}
