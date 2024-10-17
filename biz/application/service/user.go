package service

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/consts"
	usermapper "github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/user"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IUserService interface {
	SignUp(ctx context.Context, req *user.SignUpReq) (*user.SignUpResp, error)
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoResp, error)
	SetUserInfo(ctx context.Context, req *user.SetUserInfoReq) (*user.SetUserInfoResp, error)
}

type UserService struct {
	UserMongoMapper *usermapper.MongoMapper
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (s *UserService) SignUp(ctx context.Context, req *user.SignUpReq) (*user.SignUpResp, error) {
	id, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}

	role := req.Role
	var defaultName string
	switch role {
	case user.Role_DEVELOPER:
		defaultName = consts.DefaultDeveloperName
	case user.Role_ENTERPRISE:
		defaultName = consts.DefaultEnterpriseName
	}

	now := time.Now()
	newUser := &usermapper.User{
		ID:         id,
		Username:   defaultName + id.Hex()[:8],
		Role:       int(role.Number()),
		Auth:       false,
		AuthId:     "",
		Remain:     0,
		Status:     0,
		CreateTime: now,
		UpdateTime: now,
	}
	err = s.UserMongoMapper.Insert(ctx, newUser)
	if err != nil {
		return &user.SignUpResp{
			Done: false,
			Msg:  "创建用户失败",
		}, err
	}
	return &user.SignUpResp{
		Done: true,
		Msg:  "创建用户成功",
	}, nil
}

func (s *UserService) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	id := req.UserId
	aUser, err := s.UserMongoMapper.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.GetUserInfoResp{
		Username: aUser.Username,
		Role:     user.Role(aUser.Role),
		Auth:     aUser.Auth,
		AuthId:   aUser.AuthId,
		Status:   user.UserStatus(aUser.Status),
	}, err
}

func (s *UserService) SetUserInfo(ctx context.Context, req *user.SetUserInfoReq) (*user.SetUserInfoResp, error) {
	id := req.UserId
	aUser, err := s.UserMongoMapper.FindOne(ctx, id)
	if err != nil {
		return &user.SetUserInfoResp{
			Done: false,
			Msg:  "用户不存在或已删除",
		}, err
	}
	if req.Username != nil {
		aUser.Username = *req.Username
	}
	if req.Status != nil {
		aUser.Status = int(req.Status.Number())
	}
	err = s.UserMongoMapper.Update(ctx, aUser)
	if err != nil {
		return &user.SetUserInfoResp{
			Done: false,
			Msg:  "修改用户信息失败",
		}, err
	}
	return &user.SetUserInfoResp{
		Done: true,
		Msg:  "修改用户信息成功",
	}, nil

}
