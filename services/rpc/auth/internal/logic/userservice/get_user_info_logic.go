package userservicelogic

import (
	"context"

	"gomicrox/infra/errors"
	"gomicrox/services/rpc/auth/model/user"
	"gomicrox/services/rpc/auth/internal/interceptor"
	"gomicrox/services/rpc/auth/internal/svc"
	"gomicrox/services/rpc/auth/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// 强制使用 interceptor 注入的当前用户，忽略请求中的 UserId，防止水平越权
	u, ok := l.ctx.Value(interceptor.ContextKeyUser).(*user.Users)
	if !ok || u == nil {
		l.Logger.Error("user not found in context")
		return nil, errors.ErrUnauthorized
	}

	return &pb.GetUserInfoResp{
		UserId:   u.Id,
		Username: u.Username,
		Email:    u.Email,
		Avatar:   u.Avatar,
		Phone:    u.Phone,
	}, nil
}
