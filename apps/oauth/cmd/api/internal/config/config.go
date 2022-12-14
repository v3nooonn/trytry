package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CustomerRpc zrpc.RpcClientConf

	OAuth struct {
		PrivateKey string
		ExpireDay  uint16
	}
}
