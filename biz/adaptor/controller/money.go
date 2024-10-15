package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/opeanapi-user/biz/application/service"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

type IMoneyController interface {
	SetRemain(ctx context.Context, req *user.SetRemainReq) (res *user.SetRemainResp, err error)
}

type MoneyController struct {
	MoneyService *service.MoneyService
}

var MoneyControllerSet = wire.NewSet(
	wire.Struct(new(MoneyController), "*"),
	wire.Bind(new(IMoneyController), new(*MoneyController)),
)

func (c *MoneyController) SetRemain(ctx context.Context, req *user.SetRemainReq) (res *user.SetRemainResp, err error) {
	return c.MoneyService.SetRemain(ctx, req)
}
