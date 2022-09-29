<!-- @format -->

# trytry-based-on-looklook

Putting go-zero into practice locally and personally.

### API related

1. Template generation

```bash
goctl api -o production.api
```

2. Api files import and group

```go
import (
	"./car/car.api"
	"./brand/brand.api"
)

@server(
    group: carGrp
    prefix: api/production/car
)
service production-api {
	@handler CarNew
	post /estb(CarNewReq) returns(CarNewResp)

	@handler CarInfo
	post /info(CarInfoReq) returns(CarInfoResp)
}
```

3. Format

```bash
goctl api format --dir .
```

4. Generation

```bash
goctl api go -api production.api -dir ../
```

---

### Model related
1. Template Generation
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
