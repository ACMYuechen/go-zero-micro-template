# infra/redis — Redis 封装

基于 go-redis/v9 的客户端封装。

## 配置

```go
type Config struct {
    Address      string        `json:"address"`      // 地址（如 "127.0.0.1:6379"）
    Password     string        `json:"password"`     // 密码
    DB           int           `json:"db"`           // 数据库索引
    PoolSize     int           `json:"poolSize"`     // 连接池大小
    MinIdleConns int           `json:"minIdleConns"` // 最小空闲连接数
    DialTimeout  time.Duration `json:"dialTimeout"`  // 连接超时
    ReadTimeout  time.Duration `json:"readTimeout"`  // 读超时
    WriteTimeout time.Duration `json:"writeTimeout"` // 写超时
    PoolTimeout  time.Duration `json:"poolTimeout"`  // 连接池获取超时
    MaxRetries   int           `json:"maxRetries"`   // 最大重试次数
}
```

## 使用

```go
rds, err := redis.NewRedisDB(c.Redis)
if err != nil {
    logx.Must(err)
}
defer rds.Close()

// 获取原生客户端
client := rds.Client()

// 封装方法
rds.Setex("key", "value", 5*time.Minute)
val, err := rds.Get("key")
rds.Del("key1", "key2")
```

## yaml 配置示例

```yaml
Redis:
  Address: "127.0.0.1:6379"
  Password: "12345678"
  DB: 2
  PoolSize: 10
  MinIdleConns: 5
  DialTimeout: 5s
  ReadTimeout: 3s
  WriteTimeout: 3s
  PoolTimeout: 4s
  MaxRetries: 3
```
