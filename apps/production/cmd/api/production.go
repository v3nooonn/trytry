package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/config"
	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/handler"
	"github.com/v3nooonn/trytry/apps/production/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/production-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// global middelware registration
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			logx.Info("--- Production:Globle: Middleware Before")
			next(w, r)
			logx.Info("--- Production:Globle: Middleware After")
		}
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
