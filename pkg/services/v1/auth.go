package v1

import (
	"net/http"
)

func IsRegistered(w http.ResponseWriter, _ *http.Request)  {
	// 从数据库中获取
}

// Register 商户注册 仅在商户表为空时方可注册成功
func Register(w http.ResponseWriter, r *http.Request) {
	// 获取 商户全称、商户简称、手机号、
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 获取username、password
}
