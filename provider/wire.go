//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/openapi-user/biz/adaptor"
)

func NewProvider() (*adaptor.UserServer, error) {
	wire.Build(
		wire.Struct(new(adaptor.UserServer), "*"),
		UserServerProvider,
	)
	return nil, nil
}
