package main

import (
	"flag"
	"fmt"

	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/config"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/handler"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hospital-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
