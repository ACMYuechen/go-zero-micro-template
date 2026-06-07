# infra/database — 数据库封装

基于 GORM v2 的 PostgreSQL 连接池封装。

## 配置

```go
type Config struct {
    LogLevel        int           `json:"logLevel"`        // GORM 日志级别
    Driver          string        `json:"driver"`          // 驱动名（"Postgres"）
    DSN             string        `json:"dsn"`             // 连接串
    MaxOpenConns    int           `json:"maxOpenConns"`    // 最大打开连接数
    MaxIdleConns    int           `json:"maxIdleConns"`    // 最大空闲连接数
    ConnMaxLifetime time.Duration `json:"connMaxLifetime"` // 连接最大存活时间
    ConnMaxIdleTime time.Duration `json:"connMaxIdleTime"` // 连接最大空闲时间
    AutoMigrate     bool          `json:"autoMigrate"`     // 自动建表
    PrepareStmt     bool          `json:"prepareStmt"`     // 预编译语句
}
```

## 使用

```go
db, err := database.NewDatabase(c.Database)
if err != nil {
    logx.Must(err)
}
defer db.Close()

// 获取 *gorm.DB 原生对象
gormDB := db.DB()

// 原生 GORM 操作
gormDB.Where("id = ?", id).First(&user)
```

## yaml 配置示例

```yaml
Database:
  LogLevel: 4
  Driver: "Postgres"
  DSN: "host=127.0.0.1 user=root password=123456 dbname=gomicrox port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  MaxOpenConns: 100
  MaxIdleConns: 20
  ConnMaxLifetime: 7200s
  ConnMaxIdleTime: 300s
  AutoMigrate: true
  PrepareStmt: true
```
