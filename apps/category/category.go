package main

import (
	"flag"
	"fmt"

	"github.com/v3nooonn/trytry/apps/category/internal/config"
	"github.com/v3nooonn/trytry/apps/category/internal/server"
	"github.com/v3nooonn/trytry/apps/category/internal/svc"
	"github.com/v3nooonn/trytry/apps/category/pb/category"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/category.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		category.RegisterCategoryServer(grpcServer, server.NewCategoryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
