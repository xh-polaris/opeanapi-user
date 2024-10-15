package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/opeanapi-user/biz/application/service"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
)

type IAuthController interface {
	SignUp(ctx context.Context, req *user.SignUpReq) (r *user.SignUpResp, err error)
	GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error)
	SetUserInfo(ctx context.Context, req *user.SetUserInfoReq) (r *user.SetUserInfoResp, err error)
	CreateKey(ctx context.Context, req *user.CreateKeyReq) (r *user.CreateKeyResp, err error)
	GetKey(ctx context.Context, req *user.GetKeysReq) (r *user.GetKeysResp, err error)
	UpdateKey(ctx context.Context, req *user.UpdateKeyReq) (r *user.UpdateKeyResp, err error)
	UpdateHosts(ctx context.Context, req *user.UpdateHostsReq) (r *user.UpdateHostsResp, err error)
	RefreshKey(ctx context.Context, req *user.RefreshKeyReq) (r *user.RefreshKeyResp, err error)
	DeleteKey(ctx context.Context, req *user.DeleteKeyReq) (r *user.DeleteKeyResp, err error)
}

type AuthController struct {
	KeyService  *service.KeyService
	UserService *service.UserService
}

var AuthControllerSet = wire.NewSet(
	wire.Struct(new(AuthController), "*"),
	wire.Bind(new(IAuthController), new(*AuthController)),
)

func (c *AuthController) SignUp(ctx context.Context, req *user.SignUpReq) (r *user.SignUpResp, err error) {
	return c.UserService.SignUp(ctx, req)
}

func (c *AuthController) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error) {
	return c.UserService.GetUserInfo(ctx, req)
}

func (c *AuthController) SetUserInfo(ctx context.Context, req *user.SetUserInfoReq) (r *user.SetUserInfoResp, err error) {
	return c.UserService.SetUserInfo(ctx, req)
}

func (c *AuthController) CreateKey(ctx context.Context, req *user.CreateKeyReq) (r *user.CreateKeyResp, err error) {
	return c.KeyService.CreateKey(ctx, req)
}

func (c *AuthController) GetKey(ctx context.Context, req *user.GetKeysReq) (r *user.GetKeysResp, err error) {
	return c.KeyService.GetKey(ctx, req)
}

func (c *AuthController) UpdateKey(ctx context.Context, req *user.UpdateKeyReq) (r *user.UpdateKeyResp, err error) {
	return c.KeyService.UpdateKey(ctx, req)
}

func (c *AuthController) UpdateHosts(ctx context.Context, req *user.UpdateHostsReq) (r *user.UpdateHostsResp, err error) {
	return c.KeyService.UpdateHosts(ctx, req)
}

func (c *AuthController) RefreshKey(ctx context.Context, req *user.RefreshKeyReq) (r *user.RefreshKeyResp, err error) {
	return c.KeyService.RefreshKey(ctx, req)
}

func (c *AuthController) DeleteKey(ctx context.Context, req *user.DeleteKeyReq) (r *user.DeleteKeyResp, err error) {
	return c.KeyService.DeleteKey(ctx, req)
}
