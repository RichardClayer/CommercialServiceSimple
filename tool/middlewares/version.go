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
		// Accept 头
	    a := r.Header.Get("Accept")
		if len(a) == 0 {
			rd := response.RespData{
				Code:    response.AcceptNone,
				Message: "缺少 Accept 头",
			}

			response.Send(w, rd)
			return
		}

	    // Accept 头格式判断
		regexStr := `application\/` + config.APIStandardsTree + `\.` + config.APISubType + `\.v\d\+json`
		regex := regexp.MustCompile(regexStr)
		if rs := regex.FindIndex([]byte(a)); rs == nil {
			rd := response.RespData{
				Code:    response.VersionNone,
				Message: "缺少版本号或版本号配置不正确",
			}

			response.Send(w, rd)
			return
		}

        // 析取版本号
        regex = regexp.MustCompile(`\.(v\d)\+json`)
        rs := regex.FindStringSubmatch(a)
        if len(rs) < 2 {
            rd := response.RespData{
                Code:    response.VersionNotMatch,
                Message: "未知的版本号",
            }

            response.Send(w, rd)
            return
        }

        // 检查是否为受支持的版本号
        if ok := isSupportVersion(rs[1]); !ok {
            rd := response.RespData{
                Code:    response.VersionNotSupport,
                Message: "不受支持的版本号",
            }

            response.Send(w, rd)
            return
        }

		next.ServeHTTP(w, r)
	})
}

// isSupportVersion 检查是否为受支持的版本号
func isSupportVersion(v string) bool {
    rs, ok := config.SupportVersion[v]
    if !ok {
        return false
    }

    return rs
}
