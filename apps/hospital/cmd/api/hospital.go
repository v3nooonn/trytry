package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/config"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/handler"
	"github.com/v3nooonn/trytry/apps/hospital/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/hospital-api.yaml", "the config file")

var a1 struct {
	Name string
	Age  int8
}

var a2 = struct {
	Name string
	Age  int8
}{}

func main() {
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(a2))

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
