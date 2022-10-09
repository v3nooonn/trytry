<!-- @format -->

# trytry-based-on-looklook

Putting go-zero into practice locally and personally.

### API related

1. Basic template generation

```bash
goctl api -o production.api
```

2. Api files import and group

```go
// ./production.api
import (
	"./car/car.api"
	"./brand/brand.api"
)

// ./car/car.api
@server(
	middleware: Authentication,Authorization
	group: carGrp
	prefix: api/production/car
)
service production-api {
	@handler Estb
	post /estb(CarEstbReq) returns(CarEstbResp)

	@handler Info
	post /info(CarInfoReq) returns(CarInfoResp)
}
```

3. Format

```bash
goctl api format --dir .
```

4. Code Generation

```bash
goctl api go -api production.api -dir ../
```

---

### Model related

1. Code Generation

```bash
// Path: apps/production
// With: cache, fundamental table, target foler
goctl model mysql datasource --url="root:root@tcp(127.0.0.1:3306)/trytry" -table="production_brand" -c -dir=./model
```

2. Update `etc/production-api.yaml`

```yaml
Name: production-api
Host: 0.0.0.0
Port: 10001
Mode: dev

DB:
  DataSource: root:root@tcp(127.0.0.1:3306)/micro?charset=utf8mb4&parseTime=true&loc=UTC
Cache:
  - Host: 127.0.0.1:6379
    Pass: redispwd
```

3. Update `internal/config/config.go`

```go
package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
```

---

### RPC Related

1. Basic template generation

```
// .../rpc/pb
goctl rpc -o customer.proto
```

2. Edit `customer.proto`
3. Code Generation

```
// Path: /apps/customer/cmd/rpc/pb
// --go_out 与 --go-grpc_out 生成的最终目录必须一致
// --go_out & --go-grpc_out 和 --zrpc_out 的生成的最终目录必须不为同一目录
goctl rpc protoc customer.proto --go_out=. --go-grpc_out=. --zrpc_out=../
```

---

### Middilware

1. Globle

```go
// global middelware registration at entrance of a single service.
server.Use(func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("--- Globle: Middleware Before")
		next(w, r)
		logx.Info("--- Globle: Middleware After")
	}
})
```

2. Local
   `middleware` tag will create all the middlewares at `apps/service/cmd/api/internal/middleware` in seperate files.
   And also, the middleware registration is handled at `apps/service/cmd/api/internal/handler/routes.go`.

```go
@server(
	middleware: Middle1,Middle2
	...
)
service xxx-api {
	...
}
```
