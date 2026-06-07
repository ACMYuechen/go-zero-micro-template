# infra/ — 基础设施层

所有服务共享的基础设施封装，**禁止包含业务逻辑**。只提供纯技术能力的抽象。

## 设计原则

1. **零业务耦合**: 只封装技术组件，不涉及任何业务概念
2. **统一配置**: 通过 `Config` 结构体接收配置，不依赖全局变量
3. **接口抽象**: 对外暴露接口，内部实现可替换
4. **错误透明**: 底层错误包装后返回，不吞掉原始错误

## 目录结构

```
infra/
├── database/           # GORM PostgreSQL 封装
├── redis/              # go-redis 客户端封装
├── auth/               # JWT / bcrypt 工具
├── oss/                # 阿里云 OSS 封装
├── errors/             # 通用错误变量
└── rand/               # 随机工具
```

## 使用方式

```go
// 初始化（在 main.go 或 svc 中）
db, err := database.NewDatabase(c.Database)
rds, err := redis.NewRedisDB(c.Redis)

// 使用
user, err := db.DB().Where("id = ?", id).First(&u).Error
rds.Setex("key", "value", 5*time.Minute)
```

## 新增基础设施

```
infra/<component>/
├── <component>.go      # 封装实现
└── README.md           # 使用说明
```

> 每个 infra 包必须提供：
> - `Config` 结构体（映射 yaml 配置）
> - `NewXxx()` 构造函数（返回错误而非 panic）
> - `Close()` 方法（资源释放）
