<!-- @format -->

# trytry

Putting go-zero into practice locally and personally.

### API related

1. Basic template generation

```bash
goctl api -o xxx.api --remote="https://github.com/v3nooonn/trytry-template.git"
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
goctl api go -api xxx.api -dir ../
```

---

### RPC Related

1. Basic template generation

```
// .../rpc/pb
goctl rpc -o xxx.proto --remote="https://github.com/v3nooonn/trytry-template.git"
```

2. Edit `xxx.proto`
3. Code Generation

```
// Path: /apps/customer/cmd/rpc/pb
// --go_out 与 --go-grpc_out 生成的最终目录必须一致
// --go_out & --go-grpc_out 和 --zrpc_out 的生成的最终目录必须不为同一目录
goctl rpc protoc xxx.proto --go_out=. --go-grpc_out=. --zrpc_out=../
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

### Model related

#### MySQL

1. Model Generation

```bash
// Path: apps/production
// With: cache, fundamental table, target foler
goctl model mysql datasource --url="user:pwd@tcp(host:port)/db" -table="table" -c -dir=./model_dir --remote="https://github.com/v3nooonn/trytry-template"
```

---

#### PostgreSQL

1. Model Generation

```bash
// Path: apps/production
// With: schema, cache, fundamental table, target foler
 goctl model pg datasource --url="postgres://user:pwd@127.0.0.1:5434/db?sslmode=disable" --schema="schema" -table="table" -c -dir="target" --remote="https://github.com/v3nooonn/trytry-template"
```
