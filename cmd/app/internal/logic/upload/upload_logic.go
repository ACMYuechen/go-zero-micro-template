// Code scaffolded by goctl. Safe to edit.

package upload

import (
	"gomicrox/cmd/app/internal/svc"
	"gomicrox/cmd/app/internal/types"
	"gomicrox/infra/errors"
	"gomicrox/services/rpc/upload/pb"
	"context"
	"io"
	"mime/multipart"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件
func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadReq, file multipart.File, fileHeader *multipart.FileHeader) (resp *types.UploadResp, err error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 从 JWT context 取 user_id，不信任客户端传入的 req.UserId
	userId, ok := l.ctx.Value("user_id").(string)
	if !ok || userId == "" {
		return nil, errors.ErrUnauthorized
	}

	// 组装 RPC 请求
	rpcReq := &pb.UploadReq{
		File:         fileBytes,
		FileName:     fileHeader.Filename,
		FileSize:     fileHeader.Size,
		FileType:     fileHeader.Header.Get("Content-Type"),
		UploadUserID: userId,
	}

	// 将 token 通过 metadata 传递给 upload-rpc
	token := l.ctx.Value("token").(string)
	md := metadata.Pairs("authorization", "Bearer "+token)
	ctx := metadata.NewOutgoingContext(l.ctx, md)

	// 调用 RPC 服务
	rpcResp, err := l.svcCtx.UploadRpc.Upload(ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	// 组装 HTTP 响应
	return &types.UploadResp{
		Msg: rpcResp.Msg,
		Data: types.UploadData{
			FileName: rpcResp.FileName,
			FilePath: rpcResp.FilePath,
			FileSize: rpcResp.FileSize,
			FileType: rpcResp.FileType,
		},
	}, nil

}
