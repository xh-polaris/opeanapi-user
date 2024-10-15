//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/opeanapi-user/biz/adaptor"
	"github.com/xh-polaris/opeanapi-user/provider"
)

func NewProvider() (*adaptor.UserServer, error) {
	wire.Build(
		wire.Struct(new(adaptor.UserServer), "*"),
		provider.UserServerProvider,
	)
	return nil, nil
}
