# infra/errors — 通用错误变量

定义项目中使用的通用错误变量，统一错误语义。

## 使用方式

```go
import "gomicrox/infra/errors"

// 在 logic 中返回
case err == errors.ErrUserNotFound:
    return nil, errors.New("用户不存在")
case err == errors.ErrInvalidPassword:
    return nil, errors.New("密码错误")
```

## 错误分类

| 类别 | 错误变量 | 说明 |
|------|----------|------|
| 系统 | `ErrInternal` | 内部服务器错误 |
| 系统 | `ErrDatabase` | 数据库错误 |
| 用户 | `ErrUserNotFound` | 用户不存在 |
| 用户 | `ErrUserExists` | 用户已存在 |
| 用户 | `ErrInvalidPassword` | 密码错误 |
| Token | `ErrTokenGeneration` | Token 生成失败 |
| Token | `ErrInvalidToken` | Token 无效 |
| 权限 | `ErrUnauthorized` | 未授权 |
| 限流 | `ErrTooManyRequests` | 请求过于频繁 |
| 验证码 | `ErrCodeInvalid` | 验证码无效 |
| 验证码 | `ErrCodeExpired` | 验证码已过期 |
| 邮箱 | `ErrInvalidEmail` | 邮箱格式无效 |
| 邮箱 | `ErrEmailSendFailed` | 邮件发送失败 |

## 扩展

新增业务错误时，在此文件中统一定义：

```go
// 订单相关
var ErrOrderNotFound = errors.New("order not found")
var ErrOrderPaid = errors.New("order already paid")
```
