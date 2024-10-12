package service

import (
	"context"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"github.com/xhpolaris/opeanapi-user/biz/domain/service"
)

type MoneyService struct {
	RemainService *service.RemainService
}

func (s MoneyService) SetRemain(ctx context.Context, req *user.SetRemainReq) (*user.SetRemainResp, error) {
	return s.RemainService.SetRemain(ctx, req)
}
