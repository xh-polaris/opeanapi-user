package service

import (
	"context"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	usermapper "github.com/xhpolaris/opeanapi-user/biz/infrastructure/mapper/user"
)

type UserService struct {
	UserMongoMapper *usermapper.MongoMapper
}

func (s *UserService) SignUp(ctx context.Context, req *user.SignUpReq) (*user.SignUpResp, error) {

}

func (s *UserService) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {

}

func (s *UserService) SetUserInfo(ctx context.Context, req *user.SetUserInfoReq) (*user.SetUserInfoResp, error) {

}
