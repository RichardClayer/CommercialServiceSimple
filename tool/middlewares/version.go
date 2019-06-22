package middlewares

import (
	"github.com/BiLuoHui/CommercialServiceSimple/config"
	"github.com/BiLuoHui/CommercialServiceSimple/tool/response"
	"net/http"
	"regexp"
)

// RequestVersionCheck 请求头版本检查
func RequestVersionCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("Accept")
		if len(a) == 0 {
			rd := response.RespData{
				Code:    response.AcceptNone,
				Message: "缺少 Accept 头",
			}

			response.Send(w, rd)
			return
		}

		regexStr := `application\/` + config.APIStandardsTree + `\.` + config.APISubType + `\.v\d{1,1}\+json`
		regex := regexp.MustCompile(regexStr)
		if rs := regex.FindIndex([]byte(a)); rs == nil {
			rd := response.RespData{
				Code:    response.VersionNone,
				Message: "缺少版本号或版本号配置不正确",
			}

			response.Send(w, rd)
			return
		}

		next.ServeHTTP(w, r)
	})
}
