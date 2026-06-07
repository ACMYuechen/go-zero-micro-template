package logic

import (
	"bytes"
	"context"
	"path/filepath"
	"runtime"
	"testing"

	aliOSS "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/conf"

	"gomicrox/infra/oss"
	"gomicrox/services/rpc/upload/internal/config"
	"gomicrox/services/rpc/upload/internal/svc"
	"gomicrox/services/rpc/upload/pb"
)

func loadConfig(t *testing.T) config.Config {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configPath := filepath.Join(dir, "../../etc/config.yaml")

	var c config.Config
	conf.MustLoad(configPath, &c)
	return c
}

// TestOSSConnectivity 仅验证 OSS 配置是否正确、能否成功上传并生成签名URL
// 不依赖数据库，适合快速验证 AK/SK/Endpoint/Bucket 是否配对了
func TestOSSConnectivity(t *testing.T) {
	c := loadConfig(t)

	ossClient, err := oss.NewOSS(c.OSS)
	if err != nil {
		t.Fatalf("init OSS client failed: %v", err)
	}
	t.Logf("OSS client init success, endpoint: %s, bucket: %s", c.OSS.Endpoint, c.OSS.BucketName)

	objectKey := "test/oss_connectivity_test.txt"
	err = ossClient.Bucket.PutObject(objectKey, bytes.NewReader([]byte("OSS connectivity test")))
	if err != nil {
		t.Fatalf("OSS upload failed: %v", err)
	}

	// 生成签名URL（1小时有效）
	signedURL, err := ossClient.Bucket.SignURL(objectKey, aliOSS.HTTPGet, 3600)
	if err != nil {
		t.Fatalf("generate signed URL failed: %v", err)
	}

	t.Logf("OSS upload success, signed URL: %s", signedURL)
	t.Logf("请复制上面的 URL 到浏览器验证是否能访问")
}

// TestUploadLogic_Upload 测试完整的上传逻辑（含数据库写入）
// 需要 MySQL 和 upload 数据库可连接
func TestUploadLogic_Upload(t *testing.T) {
	c := loadConfig(t)

	ctx := svc.NewServiceContext(c)
	l := NewUploadLogic(context.Background(), ctx)

	req := &pb.UploadReq{
		File:         []byte("Hello, this is an OSS upload test!"),
		FileName:     "test_oss.txt",
		FileSize:     36,
		FileType:     "text/plain",
		UploadUserID: "test-user",
	}

	resp, err := l.Upload(req)
	if err != nil {
		t.Fatalf("upload failed: %v", err)
	}

	t.Logf("upload success: msg=%s, filePath=%s, fileName=%s, fileSize=%d, fileType=%s",
		resp.Msg, resp.FilePath, resp.FileName, resp.FileSize, resp.FileType)

	if resp.FilePath == "" {
		t.Fatal("FilePath is empty, OSS upload may have failed")
	}
}
