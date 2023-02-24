package main

import (
	"flag"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/v3nooonn/trytry/apps/production/internal/config"
	"github.com/v3nooonn/trytry/apps/production/internal/server"
	"github.com/v3nooonn/trytry/apps/production/internal/svc"
	"github.com/v3nooonn/trytry/apps/production/pb/production"
	"github.com/v3nooonn/trytry/pkg/interceptor/from"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/production.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		production.RegisterProductionServiceServer(grpcServer, server.NewProductionServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(
		from.Tenant,
	)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
