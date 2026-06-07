# services/rpc/upload — 文件上传服务

独立的文件上传 gRPC 服务，负责文件元数据持久化和云存储上传。

## 端口

| 协议 | 端口 | 说明 |
|------|------|------|
| gRPC | 10004 | 内部服务间调用 |

## protobuf 定义

```protobuf
service Upload {
  rpc Upload(UploadReq) returns (UploadResp);
}

message UploadReq {
  string filename = 1;
  string content_type = 2;
  bytes data = 3;
  int64 size = 4;
}

message UploadResp {
  string url = 1;
  string file_id = 2;
  string filename = 3;
}
```

## 上传流程

```
1. 客户端 → cmd/app POST /api/upload (multipart)
2. cmd/app → 读取文件字节
3. cmd/app → 调用 upload-rpc.Upload() (gRPC)
4. upload-rpc → 上传到云存储（OSS）
5. upload-rpc → 将元数据写入 files 表
6. upload-rpc → 返回签名 URL + 文件 ID
7. cmd/app → 返回给客户端
```

## 数据模型

| 表 | model | 说明 |
|----|-------|------|
| files | `services/rpc/upload/model/file/` | 文件元数据（filename, url, size, uploader_id） |

## 扩展方式

| 需求 | 修改位置 |
|------|----------|
| 更换云存储（腾讯云 COS/AWS S3） | `infra/oss/` 封装新存储驱动 |
| 增加图片压缩/水印 | `internal/logic/upload_logic.go` |
| 增加文件类型白名单 | `internal/interceptor/auth_interceptor.go` |
