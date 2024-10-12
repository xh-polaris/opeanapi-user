package controller

import (
	"context"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"github.com/xhpolaris/opeanapi-user/biz/application/service"
)

type IMoneyController interface {
	SetRemain(ctx context.Context, req *user.SetRemainReq) (res *user.SetRemainResp, err error)
}

type MoneyController struct {
	MoneyService *service.MoneyService
}

func (c *MoneyController) SetRemain(ctx context.Context, req *user.SetRemainReq) (res *user.SetRemainResp, err error) {
	return c.MoneyService.SetRemain(ctx, req)
}
