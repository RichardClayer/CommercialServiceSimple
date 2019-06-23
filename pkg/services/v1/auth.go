package v1

import (
    "context"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/database"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/response"
    "log"
    "net/http"
)

func IsRegistered(w http.ResponseWriter, _ *http.Request) {
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

    row := c.QueryRowContext(context.Background(), "select count(*) from merchant")
    var count int
    row.Scan(&count)

    response.Send(w, response.RespData{
        Code:    response.Success,
        Message: "请求成功",
        Data: map[string]int{
            "is_registered": count,
        },
    })
}

// Register 商户注册 仅在商户表为空时方可注册成功
func Register(w http.ResponseWriter, r *http.Request) {
    // 获取 商户全称、商户简称、手机号、
}

func Login(w http.ResponseWriter, r *http.Request) {
    // 获取username、password
}
