package config

import (
	"gomicrox/infra/database"
	"gomicrox/infra/oss"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Database database.Config    `json:"database"`
	OSS      oss.Config         `json:"oss"`
	AuthRpc  zrpc.RpcClientConf `json:"authRpc"` // 认证中心 RPC 客户端配置
}
