package cmd

import (
    "context"
    "flag"
    "fmt"

    "database/sql"

    "github.com/BiLuoHui/CommercialServiceSimple/pkg/protocol/rest"

    _ "github.com/go-sql-driver/mysql"
)

type Config struct {
    HTTPPort         string
    DataBaseHost     string
    DataBaseUser     string
    DataBasePassword string
    DataBaseSchema   string
}

func RunServer() error {
    // 创建Server
    ctx := context.Background()
    var cfg Config
    flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP 端口号")
    flag.StringVar(&cfg.DataBaseHost, "db-host", "", "数据库服务器主机地址")
    flag.StringVar(&cfg.DataBaseUser, "db-user", "", "数据库服务器登录用户名")
    flag.StringVar(&cfg.DataBasePassword, "db-password", "", "数据库服务器登录密码")
    flag.StringVar(&cfg.DataBaseSchema, "db-schema", "", "数据库名")
    flag.Parse()

    param := "parseTime=true"
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        cfg.DataBaseUser,
        cfg.DataBasePassword,
        cfg.DataBaseHost,
        cfg.DataBaseSchema,
        param)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("数据库连接失败：%v", err)
    }
    defer db.Close()

    // 创建HTTP服务器
    return rest.RunServer(ctx, cfg.HTTPPort)
}
