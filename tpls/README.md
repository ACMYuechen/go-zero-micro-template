# tpls/ — goctl 自定义模板

覆盖 goctl 默认代码生成模板，统一项目代码风格。

## 目录结构

```
tpls/
├── api/                    # REST API 服务模板
│   ├── main.tpl            # 入口文件
│   ├── config.tpl          # 配置结构体
│   ├── context.tpl         # ServiceContext
│   ├── handler.tpl         # HTTP Handler
│   ├── logic.tpl           # Logic 骨架
│   ├── middleware.tpl      # 中间件
│   ├── routes.tpl          # 路由注册
│   ├── types.tpl           # 请求/响应类型
│   └── ...
├── rpc/                    # RPC 服务模板
│   ├── main.tpl            # 入口文件
│   ├── config.tpl          # 配置结构体
│   ├── svc.tpl             # ServiceContext
│   ├── logic.tpl           # Logic 骨架
│   ├── server.tpl          # gRPC Server
│   ├── call.tpl            # RPC 客户端
│   └── ...
└── model/                  # Model 模板（GORM 风格）
    ├── model.tpl           # 自定义 model 包装
    ├── model-gen.tpl       # 生成文件模板
    ├── find-one.tpl        # FindOne 实现
    ├── insert.tpl          # Insert 实现
    ├── update.tpl          # Update 实现
    ├── delete.tpl          # Delete 实现
    └── ...
```

## 使用方式

```bash
# API 生成
goctl api go -home ./tpls -api cmd/app/desc/app.api -dir cmd/app -style go_zero

# RPC 生成
goctl rpc protoc services/rpc/auth/desc/auth.proto \
  --go_out=services/rpc/auth \
  --go-grpc_out=services/rpc/auth \
  --zrpc_out=services/rpc/auth \
  --style=go_zero

# Model 生成
goctl model pg datasource \
  --url "$DB_URL" \
  --home ./tpls \
  --dir model/user \
  --table users
```

## 与默认模板的区别

| 模板 | 默认 | 本模板 |
|------|------|--------|
| `api/main.tpl` | 标准 go-zero | 添加 `logx.AddGlobalFields` |
| `api/handler.tpl` | 标准 | 集成 `validator.Struct` 校验 |
| `api/context.tpl` | 标准 | 内置 `validator.Validate` |
| `model/*.tpl` | sqlx | **GORM v2** |
| `rpc/main.tpl` | 标准 | 添加 reflection 注册 |

## 修改模板

编辑 `tpls/` 下的文件后，重新运行 `make api-all` 即可生效。

> ⚠️ **注意**: 模板修改只会影响**新生成**的代码，已有代码不会被覆盖。
