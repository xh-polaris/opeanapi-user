package service

import (
	"context"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"github.com/xhpolaris/opeanapi-user/biz/domain/service"
)

type AuthService struct {
	KeyService  *service.KeyService
	UserService *service.UserService
}

func (s *AuthService) SignUp(ctx context.Context, req *user.SignUpReq) (*user.SignUpResp, error) {
	return s.UserService.SignUp(ctx, req)
}

func (s *AuthService) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	return s.UserService.GetUserInfo(ctx, req)
}

func (s *AuthService) SetUserInfo(ctx context.Context, req *user.SetUserInfoReq) (*user.SetUserInfoResp, error) {
	return s.UserService.SetUserInfo(ctx, req)
}

func (s *AuthService) CreateKey(ctx context.Context, req *user.CreateKeyReq) (*user.CreateKeyResp, error) {
	return s.KeyService.CreateKey(ctx, req)
}

func (s *AuthService) GetKey(ctx context.Context, req *user.GetKeysReq) (*user.GetKeysResp, error) {
	return s.KeyService.GetKey(ctx, req)
}

func (s *AuthService) UpdateKey(ctx context.Context, req *user.UpdateKeyReq) (*user.UpdateKeyResp, error) {
	return s.KeyService.UpdateKey(ctx, req)
}

func (s *AuthService) UpdateHosts(ctx context.Context, req *user.UpdateHostsReq) (*user.UpdateHostsResp, error) {
	return s.KeyService.UpdateHosts(ctx, req)
}

func (s *AuthService) RefreshKey(ctx context.Context, req *user.RefreshKeyReq) (*user.RefreshKeyResp, error) {
	return s.KeyService.RefreshKey(ctx, req)
}

func (s *AuthService) DeleteKey(ctx context.Context, req *user.DeleteKeyReq) (*user.DeleteKeyResp, error) {
	return s.KeyService.DeleteKey(ctx, req)
}
