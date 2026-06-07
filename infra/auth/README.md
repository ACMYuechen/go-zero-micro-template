# infra/auth — 认证工具

提供 JWT Token 生成/校验、bcrypt 密码哈希等认证相关工具函数。

## 功能

### JWT Token

```go
// 生成 Token
token, err := auth.GenerateToken(userId, secret, expire)

// 验证 Token
_, err := auth.ValidateToken(tokenString, secret)

// 检查是否过期
expired, err := auth.IsTokenExpired(tokenString, secret)

// 提取 user_id
userId, err := auth.GetUserIdFromToken(tokenString, secret)
```

### 密码

```go
// 加密密码
hashed, err := auth.HashPassword("plain_password")

// 验证密码
err := auth.ComparePassword(hashedPassword, "plain_password")
```

## 配置

```go
type Config struct {
    Secret string `json:"secret"` // JWT 签名密钥
    Expire int64  `json:"expire"` // Token 过期时间（秒）
}
```

## yaml 配置示例

```yaml
Auth:
  Secret: "your-secret-key-change-in-production"
  Expire: 86400  # 1天
```

> ⚠️ **生产环境**: 务必修改 Secret，不要使用默认值！
