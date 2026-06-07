package config

import (
	"gomicrox/infra/auth"
	"gomicrox/infra/redis"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis redis.Config       `json:"redis"`
	Auth  auth.Config        `json:"auth"`

	// rpc 配置
	UploadRpc zrpc.RpcClientConf `json:"uploadRpc"`
	AuthRpc   zrpc.RpcClientConf `json:"authRpc"`
}
