// Code scaffolded by goctl. Safe to edit.

package upload

import (
	"github.com/zeromicro/go-zero/core/logx"

	"net/http"

	"gomicrox/cmd/app/internal/logic/upload"
	"gomicrox/cmd/app/internal/svc"
	"gomicrox/cmd/app/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上传文件
func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			in  = new(types.UploadReq)
			ctx = r.Context()
		)

		if err := httpx.Parse(r, in); err != nil {
			logx.WithContext(ctx).Errorf("parse params failed: %v", err)
			httpx.Error(w, err)
			return
		}

		if err := svcCtx.Validator.Struct(in); err != nil {
			logx.WithContext(ctx).Errorf("validate params failed: %v", err)
			httpx.Error(w, err)
			return
		}

		// 解析文件（内存限制32MB，超过部分写入临时文件）
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			logx.WithContext(ctx).Errorf("parse multipart form failed: %v", err)
			httpx.Error(w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			logx.WithContext(ctx).Errorf("get file from form failed: %v", err)
			httpx.Error(w, err)
			return
		}
		defer file.Close()

		l := upload.NewUploadLogic(ctx, svcCtx)
		resp, err := l.Upload(in, file, fileHeader)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
