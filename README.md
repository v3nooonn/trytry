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
