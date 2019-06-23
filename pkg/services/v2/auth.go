package v2

import (
    "github.com/BiLuoHui/CommercialServiceSimple/tool/response"
    "net/http"
)

func IsRegistered(w http.ResponseWriter, _ *http.Request) {
    rd := response.RespData{
        Code:    response.Success,
        Message: "请求成功",
        Data:    map[string]string{
            "version": "v2",
        },
    }

    response.Send(w, rd)
    return
}
