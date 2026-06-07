# infra/oss — 阿里云 OSS 封装

阿里云对象存储的客户端封装。

## 配置

```go
type Config struct {
    Endpoint        string `json:"endpoint"`        // OSS 端点
    AccessKeyID     string `json:"accessKeyID"`     // AccessKey ID
    AccessKeySecret string `json:"accessKeySecret"` // AccessKey Secret
    BucketName      string `json:"bucketName"`      // Bucket 名称
    Domain          string `json:"domain"`          // 自定义域名（用于生成 URL）
}
```

## 使用

```go
ossClient, err := oss.NewOSS(c.OSS)
if err != nil {
    return err
}

// 上传文件
err = ossClient.Bucket.PutObject(key, reader)

// 生成签名 URL
url, err := ossClient.Bucket.SignURL(key, oss.HTTPGet, 3600)
```

## yaml 配置示例

```yaml
OSS:
  Endpoint: "oss-cn-hangzhou.aliyuncs.com"
  AccessKeyID: "your-ak-id"
  AccessKeySecret: "your-ak-secret"
  BucketName: "your-bucket"
  Domain: "https://your-bucket.oss-cn-hangzhou.aliyuncs.com"
```

## 扩展

如需支持其他云存储（腾讯云 COS、AWS S3），可：
1. 在 `infra/` 下新建包（如 `infra/cos/`）
2. 实现相同的接口契约
3. 在 `svc/service_context.go` 中选择性注入
