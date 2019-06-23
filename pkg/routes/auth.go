package routes

import (
    "github.com/BiLuoHui/CommercialServiceSimple/pkg/services/v1"
    "github.com/BiLuoHui/CommercialServiceSimple/pkg/services/v2"
    "net/http"
    "regexp"
)

var AuthRoutes = Routes{
    Route{
        Name:    "isRegistered",
        Method:  "GET",
        Pattern: "/is-registered",
        HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
            v := getVersion(r.Header.Get("Accept"))
            switch v {
            case "v1":
                v1.IsRegistered(w, r)
            case "v2":
                v2.IsRegistered(w, r)
            }
        },
    },
    Route{
        Name:    "register",
        Method:  "POST",
        Pattern: "/register",
        HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
            v := getVersion(r.Header.Get("Accept"))
            switch v {
            case "v1":
                v1.Register(w, r)
            }
        },
    },
}

// 获取请求版本号
func getVersion(a string) string {
    regex := regexp.MustCompile(`\.(v\d)\+json`)
    rs := regex.FindStringSubmatch(a)

    return rs[1]
}
