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

// 将数据格式化
func (d *RespData) json() []byte {
	bytes, _ := json.Marshal(d)

	return bytes
}

// 响应数据
func Send(w http.ResponseWriter, data RespData) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(data.json()); err != nil {
		log.Printf("%v\n", err)
	}
}
