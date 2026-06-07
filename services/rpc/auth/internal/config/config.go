package config

import (
	"gomicrox/infra/auth"
	"gomicrox/infra/database"
	"gomicrox/infra/redis"

	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Gateway  gateway.GatewayConf `json:"gateway,optional"`
	Database database.Config
	Redis    redis.Config
	Auth     auth.Config
}
