package v1

import (
    "context"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/database"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/response"
    "log"
    "net/http"
    "time"
)

var ctx context.Context

func IsRegistered(w http.ResponseWriter, _ *http.Request) {
    // 从数据库中获取
    c, err := database.Connect()
    if err != nil {
        log.Printf("数据库连接失败：%v\n", err)
        rd := response.RespData{
            Code:    response.DBConnectFailed,
            Message: "数据库连接失败",
        }

        response.Send(w, rd)
        return
    }
    defer c.Close()

    s := "select count(*) from merchant"
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    res, err := c.ExecContext(ctx, s)
    if err != nil {
        log.Printf("数据库查询失败：%v\n", err)
        rd := response.RespData{
            Code:    response.DBQueryFailed,
            Message: "查找商户数据失败",
        }

        response.Send(w, rd)
        return
    }

    log.Printf("查询结果：%v\n", res)
}

// Register 商户注册 仅在商户表为空时方可注册成功
func Register(w http.ResponseWriter, r *http.Request) {
    // 获取 商户全称、商户简称、手机号、
}

func Login(w http.ResponseWriter, r *http.Request) {
    // 获取username、password
}
