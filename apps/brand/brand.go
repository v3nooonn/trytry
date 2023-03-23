package main

import (
	"flag"
	"fmt"

	"github.com/v3nooonn/trytry/apps/brand/internal/config"
	"github.com/v3nooonn/trytry/apps/brand/internal/server"
	"github.com/v3nooonn/trytry/apps/brand/internal/svc"
	"github.com/v3nooonn/trytry/apps/brand/pb/brand"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/brand.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		brand.RegisterBrandServer(grpcServer, server.NewBrandServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
