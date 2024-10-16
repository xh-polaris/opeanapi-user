package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/opeanapi-user/biz/infrastructure/consts"
	usermapper "github.com/xh-polaris/opeanapi-user/biz/infrastructure/mapper/user"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"strconv"
)

type IMoneyService interface {
	SetRemain(ctx context.Context, req *user.SetRemainReq) (*user.SetRemainResp, error)
}

type MoneyService struct {
	UserMongoMapper *usermapper.MongoMapper
}

var MoneyServiceSet = wire.NewSet(
	wire.Struct(new(MoneyService), "*"),
	wire.Bind(new(IMoneyService), new(*MoneyService)),
)

func (s *MoneyService) SetRemain(ctx context.Context, req *user.SetRemainReq) (*user.SetRemainResp, error) {
	id := req.User.UserId
	increment := req.Increment
	var msg string
	aUser, err := s.UserMongoMapper.FindOne(ctx, id)
	if err != nil || aUser == nil {
		return &user.SetRemainResp{
			Done: false,
			Msg:  "用户不存在或已删除",
		}, err
	}
	remain := aUser.Remain
	if (increment > 0) || (increment+remain > 0) {
		remain += increment
		msg = consts.RemainIncrease + strconv.FormatInt(increment, 10)
	} else {
		return &user.SetRemainResp{
			Done: false,
			Msg:  "余额不足",
		}, err
	}
	aUser.Remain = remain
	err = s.UserMongoMapper.Update(ctx, aUser)
	if err != nil {
		return &user.SetRemainResp{
			Done: false,
			Msg:  "余额更新失败",
		}, err
	}
	return &user.SetRemainResp{
		Done: true,
		Msg:  msg,
	}, nil
}
