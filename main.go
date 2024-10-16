package main

import (
	"github.com/xh-polaris/openapi-user/biz/infrastructure/config"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/util/log"
	"github.com/xh-polaris/openapi-user/provider"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user/user"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/xh-polaris/gopkg/kitex/middleware"
	logx "github.com/xh-polaris/gopkg/util/log"
)

func main() {
	klog.SetLogger(logx.NewKlogLogger())
	s, err := provider.NewProvider()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.GetConfig().ListenOn)
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(
		s,
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GetConfig().Name}),
		server.WithMiddleware(middleware.LogMiddleware(config.GetConfig().Name)),
	)

	err = svr.Run()

	if err != nil {
		log.Error(err.Error())
	}
}
