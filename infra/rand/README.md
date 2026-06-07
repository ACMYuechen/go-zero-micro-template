# infra/rand — 随机工具

提供密码学安全的随机字符串和验证码生成。

## 使用方式

```go
import "gomicrox/infra/rand"

// 生成指定长度的随机字符串
str := rand.GenerateRandomString(16)

// 生成 32 位随机字符串
str := rand.GenerateRandomString32()

// 生成 6 位数字验证码
code := rand.GenCode()  // 例如: "042981"
```

## 特性

- 使用 `crypto/rand` 密码学安全随机源
- 验证码自动补零（始终 6 位）
