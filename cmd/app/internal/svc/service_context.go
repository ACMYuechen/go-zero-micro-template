// Code scaffolded by goctl. Safe to edit.

package svc

import (
	"gomicrox/cmd/app/internal/config"
	"gomicrox/cmd/app/internal/middleware"
	"gomicrox/services/rpc/auth/auth"

	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Validator *validator.Validate
	Redis     redis.UniversalClient

	// RPC 客户端
	AuthRpc auth.Auth

	// 中间件配置
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config, redisClient redis.UniversalClient) *ServiceContext {
	valid := validator.New(validator.WithRequiredStructEnabled())

	return &ServiceContext{
		Config:    c,
		Validator: valid,
		Redis:     redisClient,

		AuthRpc: auth.NewAuth(zrpc.MustNewClient(c.AuthRpc)),

		AuthMiddleware: middleware.NewAuthMiddleware(auth.NewAuth(zrpc.MustNewClient(c.AuthRpc))).Handle,
	}
}
