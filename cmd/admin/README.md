# cmd/admin — 管理后台网关

面向运营管理人员的 REST API Gateway，提供用户管理、数据统计等后台功能。

## 当前模块

| 模块 | 路径 | 说明 | 认证 |
|------|------|------|------|
| health | `/api/health` | 健康检查 | 公开 |
| user | `/api/admin/users` | 用户管理（列表/详情/状态/删除） | JWT + AdminRole |

## 与 cmd/app 的区别

| 维度 | cmd/app | cmd/admin |
|------|---------|-----------|
| 面向用户 | 客户端用户 | 运营管理员 |
| 认证中间件 | `AuthMiddleware`（仅校验 Token） | `AuthMiddleware`（校验 Token + AdminRole） |
| 依赖的 RPC | auth-rpc | auth-rpc |
| 典型功能 | 个人资料 | 用户列表、数据统计 |

## Admin 权限校验

```go
// middleware/auth_middleware.go 中增加了 AdminRole 检查
if int64(resp.User.Role) < user.RoleAdmin {
    http.Error(w, "unauthorized", http.StatusUnauthorized)
    return
}
```

> 当前 RoleAdmin 值定义在 `services/rpc/auth/model/user/vars.go` 中。
