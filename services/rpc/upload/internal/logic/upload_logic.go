package logic

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gomicrox/infra/errors"
	"gomicrox/services/rpc/upload/internal/interceptor"
	"gomicrox/services/rpc/upload/internal/svc"
	"gomicrox/services/rpc/upload/model/file"
	"gomicrox/services/rpc/upload/pb"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const maxFileSize = 100 * 1024 * 1024 // 100MB

// 允许上传的文件类型白名单（扩展名 -> MIME 类型）
var allowedFileTypes = map[string]string{
	".pdf":  "application/pdf",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
}

// 文件头 magic number 校验
var fileSignatures = map[string][]byte{
	".pdf":  {0x25, 0x50, 0x44, 0x46}, // %PDF
	".jpg":  {0xFF, 0xD8, 0xFF},       // JPEG
	".jpeg": {0xFF, 0xD8, 0xFF},       // JPEG
	".png":  {0x89, 0x50, 0x4E, 0x47}, // PNG
	".docx": {0x50, 0x4B, 0x03, 0x04}, // ZIP (docx 基于 zip)
	".doc":  {0xD0, 0xCF, 0x11, 0xE0}, // MS Office OLE
}

func (l *UploadLogic) Upload(in *pb.UploadReq) (*pb.UploadResp, error) {
	// 从 interceptor context 取 user_id，校验与请求中的 UploadUserID 是否一致
	userId, ok := l.ctx.Value(interceptor.ContextKeyUserID).(string)
	if !ok || userId == "" {
		return nil, errors.ErrUnauthorized
	}
	if userId != in.UploadUserID {
		return nil, errors.ErrUnauthorized
	}

	// 文件大小校验
	if in.FileSize > maxFileSize {
		return nil, fmt.Errorf("file size exceeds limit: %d bytes (max %d)", in.FileSize, maxFileSize)
	}

	// 文件扩展名校验
	ext := strings.ToLower(filepath.Ext(in.FileName))
	if _, allowed := allowedFileTypes[ext]; !allowed {
		return nil, fmt.Errorf("file type not allowed: %s", ext)
	}

	// 文件头 magic number 校验
	if sig, ok := fileSignatures[ext]; ok {
		if len(in.File) < len(sig) || !bytes.Equal(in.File[:len(sig)], sig) {
			return nil, fmt.Errorf("file signature mismatch: %s", ext)
		}
	}

	// 1. 上传到阿里云OSS
	fileUrl, err := l.uploadToAliOSS(in.File, in.FileName)
	if err != nil {
		logx.Error("OSS上传失败:", err)
		return nil, err
	}

	// 2. 写入数据库
	fileRecord := &file.File{
		FileName:     in.FileName,
		FilePath:     fileUrl,
		FileSize:     in.FileSize,
		FileType:     in.FileType,
		UploadUserID: in.UploadUserID,
		IsPrivate:    true,
	}
	if err := l.svcCtx.DB.Create(fileRecord).Error; err != nil {
		logx.Error("数据库入库失败:", err)
		return nil, err
	}

	// 3. 返回结果
	return &pb.UploadResp{
		Msg:      "上传成功",
		FilePath: fileUrl,
		FileName: in.FileName,
		FileSize: in.FileSize,
		FileType: in.FileType,
	}, nil
}

// 阿里云OSS上传工具
func (l *UploadLogic) uploadToAliOSS(fileBytes []byte, fileName string) (string, error) {
	// 生成唯一路径
	now := time.Now()
	dateDir := now.Format("20060102")
	uniqueID := strconv.FormatInt(now.UnixNano(), 10)
	ext := filepath.Ext(fileName)
	objectKey := "uploads/" + dateDir + "/" + uniqueID + ext

	// 上传文件
	err := l.svcCtx.OSS.Bucket.PutObject(objectKey, bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}

	// 生成带签名的临时访问URL（1小时有效）
	signedURL, err := l.svcCtx.OSS.Bucket.SignURL(objectKey, oss.HTTPGet, 3600)
	if err != nil {
		return "", err
	}

	return signedURL, nil
}
