package svc

import (
	"gomicrox/infra/oss"
	"gomicrox/services/rpc/auth/auth"
	"gomicrox/services/rpc/upload/internal/config"
	"gomicrox/services/rpc/upload/model/file"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	OSS     *oss.OSS
	AuthRpc auth.Auth // 认证中心 RPC 客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 PostgreSQL
	db, err := gorm.Open(postgres.Open(c.Database.DSN), &gorm.Config{})
	if err != nil {
		logx.Must(err)
	}
	// 自动迁移表
	db.AutoMigrate(&file.File{})

	// 初始化阿里云 OSS
	ossClient, err := oss.NewOSS(c.OSS)
	if err != nil {
		logx.Must(err)
	}

	// 初始化认证中心 RPC 客户端
	authRpc := auth.NewAuth(zrpc.MustNewClient(c.AuthRpc))

	return &ServiceContext{
		Config:  c,
		DB:      db,
		OSS:     ossClient,
		AuthRpc: authRpc,
	}
}
