package controller

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"github.com/xhpolaris/opeanapi-user/biz/application/service"
)

type IAuthController interface {
	SignUp(ctx context.Context, Req *user.SignUpReq) (r *user.SignUpResp, err error)
	GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error)
	SetUserInfo(ctx context.Context, Req *user.SetUserInfoReq) (r *user.SetUserInfoResp, err error)
	CreateKey(ctx context.Context, Req *user.CreateKeyReq) (r *user.CreateKeyResp, err error)
	GetKey(ctx context.Context, Req *user.GetKeysReq) (r *user.GetKeysResp, err error)
	UpdateKey(ctx context.Context, Req *user.UpdateKeyReq) (r *user.UpdateKeyResp, err error)
	UpdateHosts(ctx context.Context, Req *user.UpdateHostsReq) (r *user.UpdateHostsResp, err error)
	RefreshKey(ctx context.Context, Req *user.RefreshKeyReq) (r *user.RefreshKeyResp, err error)
	DeleteKey(ctx context.Context, Req *user.DeleteKeyReq) (r *user.DeleteKeyResp, err error)
}

type AuthController struct {
	KeyService  *service.KeyService
	UserService *service.UserService
}

var AuthControllerSet = wire.NewSet(
	wire.Struct(new(AuthController), "*"),
	wire.Bind(new(IAuthController), new(*AuthController)),
)

func (c *AuthController) SignUp(ctx context.Context, Req *user.SignUpReq) (r *user.SignUpResp, err error) {
	return c.UserService.SignUp(ctx, Req)
}

func (c *AuthController) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error) {
	return c.UserService.GetUserInfo(ctx, Req)
}

func (c *AuthController) SetUserInfo(ctx context.Context, Req *user.SetUserInfoReq) (r *user.SetUserInfoResp, err error) {
	return c.UserService.SetUserInfo(ctx, Req)
}

func (c *AuthController) CreateKey(ctx context.Context, Req *user.CreateKeyReq) (r *user.CreateKeyResp, err error) {
	return c.KeyService.CreateKey(ctx, Req)
}

func (c *AuthController) GetKey(ctx context.Context, Req *user.GetKeysReq) (r *user.GetKeysResp, err error) {
	return c.KeyService.GetKey(ctx, Req)
}

func (c *AuthController) UpdateKey(ctx context.Context, Req *user.UpdateKeyReq) (r *user.UpdateKeyResp, err error) {
	return c.KeyService.UpdateKey(ctx, Req)
}

func (c *AuthController) UpdateHosts(ctx context.Context, Req *user.UpdateHostsReq) (r *user.UpdateHostsResp, err error) {
	return c.KeyService.UpdateHosts(ctx, Req)
}

func (c *AuthController) RefreshKey(ctx context.Context, Req *user.RefreshKeyReq) (r *user.RefreshKeyResp, err error) {
	return c.KeyService.RefreshKey(ctx, Req)
}

func (c *AuthController) DeleteKey(ctx context.Context, Req *user.DeleteKeyReq) (r *user.DeleteKeyResp, err error) {
	return c.KeyService.DeleteKey(ctx, Req)
}
