package routes

import (
    "net/http"
    "regexp"

    "github.com/BiLuoHui/CommercialServiceSimple/pkg/services/v1"
    "github.com/BiLuoHui/CommercialServiceSimple/pkg/services/v2"
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
    Route{
        Name:    "login",
        Method:  "POST",
        Pattern: "/login",
        HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
            v := getVersion(r.Header.Get("Accept"))
            switch v {
            case "v1":
                v1.Login(w, r)
            }
        },
    },
    Route{
        Name:    "userInfo",
        Method:  "GET",
        Pattern: "/user-info",
        HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
            v := getVersion(r.Header.Get("Accept"))
            switch v {
            case "v1":
                v1.UserInfo(w, r)
            }
        },
    },
    Route{
        Name:"logout",
        Method:"DELETE",
        Pattern:"/logout",
        HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
            v := getVersion(r.Header.Get("Accept"))
            switch v {
            case "v1":
                v1.LoginOut(w, r)
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
