package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/account"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

type IAccountService interface {
	GetAccountByTxId(ctx context.Context, req *user.GetAccountByTxIdReq) (*user.GetAccountByTxIdResp, error)
}

type AccountService struct {
	AccountMongoMapper *account.MongoMapper
}

var AccountServiceSet = wire.NewSet(
	wire.Struct(new(AccountService), "*"),
	wire.Bind(new(IAccountService), new(*AccountService)),
)

func (s *AccountService) GetAccountByTxId(ctx context.Context, req *user.GetAccountByTxIdReq) (*user.GetAccountByTxIdResp, error) {
	txId := req.Id
	a, err := s.AccountMongoMapper.FindOneByTxId(ctx, txId)
	if err != nil {
		return nil, err
	}
	return &user.GetAccountByTxIdResp{Account: &user.Account{
		Id:         a.ID.Hex(),
		TxId:       a.TxId,
		Increment:  a.Increment,
		UserId:     a.UserId,
		CreateTime: a.CreateTime.Unix(),
	}}, nil
}
