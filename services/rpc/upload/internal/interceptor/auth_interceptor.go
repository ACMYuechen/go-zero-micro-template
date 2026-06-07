package interceptor

import (
	"context"
	"strings"

	"gomicrox/infra/errors"
	"gomicrox/services/rpc/auth/pb"
	"gomicrox/services/rpc/upload/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// contextKey 用于在 context 中存储认证信息
type contextKey string

const ContextKeyUserID contextKey = "user_id"

// AuthInterceptor 从 gRPC metadata 中提取 Authorization token，
// 通过 auth-rpc 验证并注入 user_id
func AuthInterceptor(svcCtx *svc.ServiceContext) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 从 metadata 取 Authorization
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.ErrInvalidToken
		}

		var tokenString string
		if vals := md.Get("authorization"); len(vals) > 0 {
			tokenString = trimBearerToken(vals[0])
		}

		if tokenString == "" {
			return nil, errors.ErrInvalidToken
		}

		// 通过 auth-rpc 验证 token
		resp, err := svcCtx.AuthRpc.ValidateToken(ctx, &pb.ValidateTokenReq{Token: tokenString})
		if err != nil {
			logx.WithContext(ctx).Errorf("validate token via auth-rpc failed: %v", err)
			return nil, errors.ErrInvalidToken
		}
		if resp.User == nil || resp.User.Id == "" {
			return nil, errors.ErrInvalidToken
		}

		// 注入 user_id 到 context
		ctx = context.WithValue(ctx, ContextKeyUserID, resp.User.Id)

		return handler(ctx, req)
	}
}

func trimBearerToken(authorization string) string {
	authorization = strings.TrimSpace(authorization)
	if authorization == "" {
		return ""
	}
	if len(authorization) >= 7 && strings.EqualFold(authorization[:7], "Bearer ") {
		return strings.TrimSpace(authorization[7:])
	}
	return authorization
}
