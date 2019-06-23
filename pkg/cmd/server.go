package cmd

import (
    "context"
    "flag"
    "fmt"

    "database/sql"

    "github.com/BiLuoHui/CommercialServiceSimple/pkg/protocol/rest"
    "github.com/BiLuoHui/CommercialServiceSimple/pkg/services"

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

    if len(cfg.HTTPPort) == 0 {
        return fmt.Errorf("请指定HTTP端口号\n")
    }

    if err := initProject(cfg); err != nil {
        return err
    }

    // 创建HTTP服务器
    return rest.RunServer(ctx, cfg.HTTPPort)
}

func initProject(cfg Config) error {
    if err := initDb(cfg); err != nil {
        return err
    }

    return nil
}

func initDb(cfg Config) error {
    if len(cfg.DataBaseHost) == 0 {
        return fmt.Errorf("请指定数据库服务器主机地址\n")
    }
    if len(cfg.DataBaseUser) == 0 {
        return fmt.Errorf("请指定数据库服务器登录用户名\n")
    }
    if len(cfg.DataBasePassword) == 0 {
        return fmt.Errorf("请指定数据库服务器登录密码\n")
    }
    if len(cfg.DataBaseSchema) == 0 {
        return fmt.Errorf("请指定数据库名\n")
    }

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

    services := services.NewService(db)

    return nil
}