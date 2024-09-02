package main

import (
	"flag"
	"fmt"

	"github.com/Sanagiig/fox-admin-core/example_rpc/internal/config"
	"github.com/Sanagiig/fox-admin-core/example_rpc/internal/server"
	"github.com/Sanagiig/fox-admin-core/example_rpc/internal/svc"
	"github.com/Sanagiig/fox-admin-core/example_rpc/types/examplerpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/example_rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		examplerpc.RegisterExampleRpcServer(grpcServer, server.NewExampleRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
