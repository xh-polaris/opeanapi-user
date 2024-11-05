package service

import (
	"context"
	"github.com/google/wire"
	usermapper "github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/user"
	usertransaction "github.com/xh-polaris/openapi-user/biz/infrastructure/transaction"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"strconv"
)

type IMoneyService interface {
	SetRemain(ctx context.Context, req *user.SetRemainReq) (*user.SetRemainResp, error)
}

type MoneyService struct {
	UserMongoMapper *usermapper.MongoMapper
	UserTransaction *usertransaction.UserTransaction
}

var MoneyServiceSet = wire.NewSet(
	wire.Struct(new(MoneyService), "*"),
	wire.Bind(new(IMoneyService), new(*MoneyService)),
)

func (s *MoneyService) SetRemain(ctx context.Context, req *user.SetRemainReq) (*user.SetRemainResp, error) {
	id := req.UserId
	increment := req.Increment
	err := s.UserTransaction.UpdateRemain(ctx, id, increment)
	if err != nil {
		return &user.SetRemainResp{
			Done: false,
			Msg:  "余额更新失败",
		}, err
	}
	return &user.SetRemainResp{
		Done: true,
		Msg:  "余额更新成功" + strconv.FormatInt(increment, 10),
	}, nil
}
