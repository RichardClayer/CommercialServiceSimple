package database

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/Unknwon/goconfig"

    _ "github.com/go-sql-driver/mysql"
)

const configFile = "config/app.ini"

func getDbConfig() (dbConfig map[string]string, err error) {
    conf, err := goconfig.LoadConfigFile(configFile)
    if err != nil {
        return
    }

    return conf.GetSection("db")
}

func Connect() (*sql.Conn, error) {
    c, err := getDbConfig()
    if err != nil {
        return nil, err
    }

    // user:password@tcp(host)/schema?parseTime=true
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
        c["user"],
        c["password"],
        c["host"],
        c["schema"],
        "parseTime=true")
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("连接数据库失败：%v", err)
    }
    defer db.Close()

    ctx := context.Background()
    conn, err := db.Conn(ctx)
    if err != nil {
        return nil, fmt.Errorf("连接数据库失败：%v", err)
    }

    return conn, nil
}
