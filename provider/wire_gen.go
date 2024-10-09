// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/config"
)

// Injectors from wire.go:

func NewProvider() (*Provider, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	providerProvider := &Provider{
		Config: configConfig,
	}
	return providerProvider, nil
}
