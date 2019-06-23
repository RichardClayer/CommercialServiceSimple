package response

// 响应格式
import (
    "encoding/json"
    "log"
    "net/http"
)

type RespData struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// json 将数据格式化
func (d *RespData) json() []byte {
    bytes, _ := json.Marshal(d)

    return bytes
}

// Send 响应数据
func Send(w http.ResponseWriter, data RespData) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    if _, err := w.Write(data.json()); err != nil {
        log.Printf("%v\n", err)
    }
}

// SendSuccess 发送成功的响应
func SendSuccess(w http.ResponseWriter, d interface{}) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    rd := RespData{
        Code:    Success,
        Message: "请求成功",
        Data:    d,
    }
    if _, err := w.Write(rd.json()); err != nil {
        log.Printf("%v\n", err)
    }
    return
}

// SendError 发送失败的响应
func SendError(w http.ResponseWriter, code int, msg string) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    rd := RespData{
        Code:    code,
        Message: msg,
    }
    if _, err := w.Write(rd.json()); err != nil {
        log.Printf("%v\n", err)
    }
    return
}
