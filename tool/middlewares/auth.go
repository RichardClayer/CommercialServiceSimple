package middlewares

import (
	"github.com/BiLuoHui/CommercialServiceSimple/config"
	"github.com/BiLuoHui/CommercialServiceSimple/tool/auth"
	"github.com/BiLuoHui/CommercialServiceSimple/tool/response"
	"net/http"
)

// RequestAuth 是否需要登录验证
func RequestAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		if isNeed(path, config.NoNeedRoutes) {
			token := r.Header.Get("Authorization")
			if ok := auth.IsOnline(token); !ok {
				rd := response.RespData{
					Code:    response.NeedLogin,
					Message: "请先登录",
				}
				response.Send(w, rd)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// 是否需要登录认证
func isNeed(path string, noNeed []string) bool {
	for _, s := range noNeed {
		if path == s {
			return false
		}
	}

	return true
}
