package routes

import (
    "github.com/BiLuoHui/CommercialServiceSimple/tool/middlewares"
    "github.com/gorilla/mux"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes = []Route

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    useMiddlewares(router)

    // 加载路由
    for _, route := range AuthRoutes {
        router.Name(route.Name).
            Methods(route.Method).
            Path(route.Pattern).
            HandlerFunc(route.HandlerFunc)
    }

    return router
}

// 中间件
func useMiddlewares(r *mux.Router) {
    r.Use(middlewares.RequestVersionCheck)
    r.Use(middlewares.RequestAuth)
}
