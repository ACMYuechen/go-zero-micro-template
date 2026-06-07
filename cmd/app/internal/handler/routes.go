// Code scaffolded by goctl. Safe to edit.

package handler

import (
	"net/http"

	health "gomicrox/cmd/app/internal/handler/health"
	"gomicrox/cmd/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 健康检查
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: health.HealthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}
