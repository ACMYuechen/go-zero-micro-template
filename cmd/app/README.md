# cmd/app — 应用网关

面向客户端的 REST API Gateway，聚合所有对外提供的接口。

## 当前模块

| 模块 | 路径 | 说明 | 认证 |
|------|------|------|------|
| health | `/api/health` | 健康检查 | 公开 |
| upload | `/api/upload` | 文件上传（代理到 upload-rpc） | JWT |

## 新增模块步骤

1. 在 `desc/` 下创建 `<module>/<module>.api`
2. 在 `desc/app.api` 中 `import` 并添加 `service` 块
3. 运行 `make api-all` 生成代码
4. 在 `internal/logic/<module>/` 下编写业务逻辑

## 配置说明

```yaml
# cmd/app/etc/config.yaml
Auth:
  Secret: "your-secret-key"      # JWT 签名密钥
  Expire: 86400                   # Token 过期时间（秒）

Redis:
  Address: "127.0.0.1:6379"      # Redis 地址

# RPC 客户端配置（app 不直连数据库，全部走 RPC）
AuthRpc:
  Endpoints:
    - 127.0.0.1:10003
UploadRpc:
  Endpoints:
    - 127.0.0.1:10004
```
