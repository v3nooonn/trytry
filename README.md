<!-- @format -->

# trytry

Putting go-zero into practice locally and personally.

### MakeFile
All the related cmds are gathered in the makefile in each of the service. Commands as below if you want to execute them manually. 

### API related

1. Basic api file generation

```bash
goctl api -o api_dir xxx.api 
```

2. Api files import and group

```go
// ./production.outgoing
import (
	"./car/car.outgoing"
	"./brand/brand.outgoing"
)

// ./car/car.outgoing
@server(
   middleware: Authentication, Authorization, Language, RemoteAddr
   group: GROUP
   prefix: PREFIX
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
goctl api go -api ./def/bff.api -dir .
```

---

### RPC Related

1. Basic template generation

```
// .../rpc/pb
goctl rpc -o ./pb/brand.proto
```

2. Edit `xxx.proto`
3. Code Generation
```bash
goctl rpc protoc ./pb/brand.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
```

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
