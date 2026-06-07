package main

import (
	"flag"
	"fmt"

	"gomicrox/services/rpc/upload/internal/config"
	"gomicrox/services/rpc/upload/internal/interceptor"
	"gomicrox/services/rpc/upload/internal/server"
	"gomicrox/services/rpc/upload/internal/svc"
	"gomicrox/services/rpc/upload/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUploadServer(grpcServer, server.NewUploadServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 添加认证拦截器
	s.AddUnaryInterceptors(interceptor.AuthInterceptor(ctx))
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
