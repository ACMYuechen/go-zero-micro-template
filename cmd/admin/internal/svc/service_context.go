// Code scaffolded by goctl. Safe to edit.

package svc

import (
	"gomicrox/cmd/admin/internal/config"
	"gomicrox/cmd/admin/internal/middleware"
	"gomicrox/services/rpc/auth/auth"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Validator *validator.Validate

	// RPC 客户端（所有数据操作通过 auth-rpc，admin 不直接访问 DB）
	AuthRpc auth.Auth

	// 中间件配置
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	valid := validator.New(validator.WithRequiredStructEnabled())

	authRpc := auth.NewAuth(zrpc.MustNewClient(c.AuthRpc))

	return &ServiceContext{
		Config:         c,
		Validator:      valid,
		AuthRpc:        authRpc,
		AuthMiddleware: middleware.NewAuthMiddleware(authRpc).Handle,
	}
}
